package seeder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

// list files in sqls folder
func GetListFile(basePath string) ([]string, error) {
	// open folder
	f, err := os.Open(basePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read directory
	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	// append files to slice and return it
	var fileArr []string
	for _, v := range files {
		if v.IsDir() {
			fmt.Printf("skip %s because directory", v.Name())
			continue
		}
		fileArr = append(fileArr, v.Name())
	}
	return fileArr, nil
}

// write new formatted string sql file to seed.sql (unique)
func WriteSeedSql(files []string) error {
	// open file
	f, err := os.OpenFile("seed.sql", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// read file
	lineChan := make(chan string)
	go func(ch chan string) {
		scanner := bufio.NewScanner(f)
		re := regexp.MustCompile(`\\i sqls\/data-\w*.sql`)
		for scanner.Scan() {
			line := scanner.Text()
			if !re.MatchString(line) {
				continue
			}
			ch <- line
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		close(ch)
	}(lineChan)

	// create unique map old line sql file
	oldLineMap := make(map[string]struct{})
	for vLine := range lineChan {
		if _, ok := oldLineMap[vLine]; !ok {
			oldLineMap[vLine] = struct{}{}
		}
	}

	// check list files from folder sqls already writed to seed.sql if not yet will append to end of file
	for iFile := range files {
		lineStr := fmt.Sprintf("\\i sqls/%s", files[iFile])
		if _, ok := oldLineMap[lineStr]; !ok {
			if _, err := f.WriteString(fmt.Sprintf("%s\n", lineStr)); err != nil {
				return err
			}
		}
	}

	return nil
}

// import sql file to database use exec from bash psql command
func ImportToDb(cmd string) error {
	exec := exec.Command("bash", "-c", cmd)
	stderr, err := exec.StderrPipe()
	if err != nil {
		return fmt.Errorf("exec stderrpipe: %w", err)
	}
	if err := exec.Start(); err != nil {
		return fmt.Errorf("cmd start err: %w", err)
	}

	errout, _ := io.ReadAll(stderr)

	if err := exec.Wait(); err != nil {
		return fmt.Errorf("cmd wait err: %s", errout)
	}
	if len(errout) != 0 {
		return fmt.Errorf("exec err: %s", errout)
	}
	return nil
}
