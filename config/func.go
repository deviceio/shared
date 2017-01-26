package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var configStruct interface{}
var fileNames = []string{}
var filePaths = []string{}

// AddFileName ...
func AddFileName(name string) {
	fileNames = append(fileNames, name)
}

// AddFilePath ...
func AddFilePath(path string) {
	filePaths = append(filePaths, path)
}

// SetConfigStruct ...
func SetConfigStruct(s interface{}) {
	configStruct = s
}

// Parse ...
func Parse() error {
	for _, path := range filePaths {
		for _, file := range fileNames {
			fullpath := fmt.Sprintf("%v/%v", strings.TrimRight(path, "/"), file)

			_, err := os.Stat(fullpath)

			if err != nil {
				continue
			}

			b, err := ioutil.ReadFile(fullpath)

			if err != nil {
				return fmt.Errorf("Error reading config file '%v': %v", fullpath, err.Error())
			}

			if err := json.Unmarshal(b, &configStruct); err != nil {
				return fmt.Errorf("Error unmarshalling config file '%v': %v", fullpath, err.Error())
			}

			return nil
		}
	}

	return errors.New(fmt.Sprintf(
		"No configurations found for files '%v' in paths: \r\n %v",
		strings.Join(fileNames, ","),
		strings.Join(filePaths, "\n"),
	))
}
