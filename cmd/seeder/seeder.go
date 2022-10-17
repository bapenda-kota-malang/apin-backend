/*
This flow
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/viper"
)

// config struct
type Config struct {
	DbName   string
	Username string
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
func getListFile() ([]string, error) {
	// open folder
	f, err := os.Open("sqls")
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

// import sql file to database use exec from bash psql command
func importToDb(username, database, basePath, sqlFile string) error {
	seedPath := path.Join(basePath, sqlFile)

	cmd := exec.Command("bash", "-c", fmt.Sprintf("psql -U %s -d %s < %s", username, database, seedPath))
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
	files, err := getListFile()
	if err != nil {
		log.Fatalf("get list file: %s", err)
	}
	for _, v := range files {
		if err := importToDb(c.Username, c.DbName, basePath, v); err != nil {
			log.Fatalf("seeding failed: %s: %s", v, err)
		}
	}
	log.Println("seeding done")
}
