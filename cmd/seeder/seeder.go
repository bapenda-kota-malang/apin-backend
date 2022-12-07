/*
This flow
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"

	"github.com/spf13/viper"
)

// config struct
type Config struct {
	Host     string
	Port     string
	DbUser   string
	Password string
	DbName   string
}

// load from yml files to struct config
func (c *Config) loadConfig() (err error) {
	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// read the file
	if err = viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("error reading config file: %s", err)
		return
	}

	// map to app
	if err = viper.Unmarshal(&c); err != nil {
		err = fmt.Errorf("unable to decode into struct: %s", err)
		return
	}

	// done
	log.Printf("config loaded successfully")
	return
}

// list files in sqls folder
func getListFile(basePath string) ([]string, error) {
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
func writeSeedSql(files []string) error {
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
func importToDb(c *Config) error {
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DbUser, c.Password, c.Host, c.Port, c.DbName)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("psql %s -1 -f seed.sql", uri))
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("exec stderrpipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("cmd start err: %s", err)
	}

	errout, _ := io.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("cmd wait err: %s", errout)
	}
	if len(errout) != 0 {
		return errors.New(string(errout))
	}
	return nil
}

func main() {
	// load config
	var c = &Config{}
	if err := c.loadConfig(); err != nil {
		log.Fatalf("load config: %s", err)
	}

	// get path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("get wd path: %s", err)
	}
	basePath := path.Join(wd, "sqls")

	// get list files
	files, err := getListFile(basePath)
	if err != nil {
		log.Fatalf("get list file: %s", err)
	}

	// create seed.sql for make one file sql
	if err := writeSeedSql(files); err != nil {
		log.Fatalf("write sql file: %s", err)
	}

	if err := importToDb(c); err != nil {
		log.Fatalf("seeding failed: %s", err)
	}
	log.Println("seeding done")
}
