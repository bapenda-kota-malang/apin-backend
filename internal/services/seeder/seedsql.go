package seeder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/seeder"
	"gorm.io/gorm"
)

type seedSqlService struct {
	basePath string
	Db       *gorm.DB
}

func NewSeed(basePath string, db *gorm.DB) seeder.SeedSqlService {
	return &seedSqlService{basePath, db}
}

// import sql file to database use exec from bash psql command
func (s *seedSqlService) RunImportSql(cmd string) error {
	if cmd == "" {
		return fmt.Errorf("cmd argument empty")
	}
	var ex *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		ex = exec.Command("bash", "-c", cmd)
	case "windows":
		ex = exec.Command("cmd", "/C", cmd)
	default:
		return fmt.Errorf("OS not supported")
	}

	stderr, err := ex.StderrPipe()
	if err != nil {
		return fmt.Errorf("exec stderrpipe: %w", err)
	}
	if err := ex.Start(); err != nil {
		return fmt.Errorf("cmd start err: %w", err)
	}

	errout, _ := io.ReadAll(stderr)

	if err := ex.Wait(); err != nil {
		return fmt.Errorf("cmd wait err: %s", errout)
	}
	if len(errout) != 0 {
		return fmt.Errorf("exec err: %s", errout)
	}

	return nil
}

// list files in sqls folder
func (s *seedSqlService) GetListFile() ([]string, error) {
	// open folder
	f, err := os.Open(path.Join(s.basePath, "sqls"))
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
func (s *seedSqlService) WriteSeedSql(files []string) error {
	// open file
	f, err := os.OpenFile(path.Join(s.basePath, "seed.sql"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
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

// get list function from db
func (s *seedSqlService) getListFunction() ([]seeder.StoreProcedure, error) {
	var listFunction []seeder.StoreProcedure
	queryListSp := `select 
		pp.proname,
		pl.lanname,
		pn.nspname,
		pg_get_functiondef(pp.oid) as queryFunction
	from pg_proc pp 
		inner join pg_namespace pn on (pp.pronamespace = pn.oid) 
		inner join pg_language pl on (pp.prolang = pl.oid) 
	where pl.lanname NOT IN ('c','internal') 
		and pn.nspname NOT LIKE 'pg_%' 
		and pn.nspname <> 'information_schema';`
	resRows, err := s.Db.Raw(queryListSp).Rows()
	if err != nil {
		return nil, err
	}
	defer resRows.Close()

	for resRows.Next() {
		var spRow seeder.StoreProcedure
		s.Db.ScanRows(resRows, &spRow)
		listFunction = append(listFunction, spRow)
	}

	if len(listFunction) == 0 {
		return nil, fmt.Errorf("list function empty")
	}

	return listFunction, nil
}
