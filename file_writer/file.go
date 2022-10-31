package filewriter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/pkg/errors"
)

var (
	pathRegex = regexp.MustCompile(`(?m)(.*)(\..*)$`)
)

//WriteFile writes the data in the specific path, if the file already exists will be overwritten
func WriteFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return errors.Wrapf(err, "writing file %s", path)
	}

	log.Printf("file writted to %s\n", path)

	return nil
}

//WriteUniqueFile writes a file adding a (n) in the name to create a unique file
//
//Returns the path of where the file finally was saved
func WriteUniqueFile(path string, data []byte) (string, error) {
	pathBase := path
	fileCount := 1
	exists := ExistsFile(path)

	for exists {
		if pathRegex.MatchString(pathBase) {
			path = pathRegex.ReplaceAllString(pathBase, fmt.Sprintf("$1(%d)$2", fileCount))
		} else {
			path = fmt.Sprintf("%s(%d)", pathBase, fileCount)
		}

		fileCount++

		exists = ExistsFile(path)
	}

	err := WriteFile(path, data)
	if err != nil {
		return path, errors.Wrapf(err, "writing unique file %s", path)
	}

	return path, nil
}

func CreateDir(path string) error {
	log.Printf("creating folder: %s\n", path)

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrapf(err, "creating folder %s", path)
	}

	return nil
}

//ConcatStringToByteArray joins strings and returns all strings in one array of bytes
func ConcatStringToByteArray(strings ...string) []byte {
	buffer := bytes.NewBufferString("")

	for _, script := range strings {
		buffer.WriteString(script)
		buffer.WriteString("\n")
	}

	return buffer.Bytes()
}

func ReadFile(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.Wrapf(err, "reading file %s", path)
	}

	return dat, nil
}

func ExistsFile(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

//WriteJSONFile writes the data formatted as JSON in the specific path, if the file already exists will be overwritten
func WriteJSONFile(data interface{}, path string) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return errors.Wrapf(err, "writte JSON to file %s", path)
	}

	return WriteFile(path, jsonData)
}

//WriteJSONUniqueFile writes data formmatted as JSON in a file adding a (n) in the name to create a unique file
//
//Returns the path of where the file finally was saved
func WriteJSONUniqueFile(data interface{}, path string) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return path, errors.Wrapf(err, "writte JSON to unique file %s", path)
	}

	return WriteUniqueFile(path, jsonData)
}

//GetCallerDir returns the path of the file from where the function is called
func GetCallerDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("get dir for connection file")
	}

	return path.Dir(filename), nil
}

//GetCallerDir returns the relative path of the file from where the function is called
func GetCallerRelPath() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("get relative path of file")
	}

	wd, e := os.Getwd()
	if e != nil {
		wd = "./"
	}

	return filepath.Rel(wd, filename)
}
