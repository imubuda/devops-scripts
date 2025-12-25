package devops_scripts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetEnvVariable returns the value of the specified environment variable.
func GetEnvVariable(varName string) (string, error) {
	return os.Getenv(varName), nil
}

// ReadFile reads the contents of a file and returns them as a string.
func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err!= nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes the specified data to a file.
func WriteFile(filePath string, data string) error {
	return ioutil.WriteFile(filePath, []byte(data), 0644)
}

// DeleteFile deletes the specified file.
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// CreateDir creates a new directory with the specified path.
func CreateDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// DeleteDir deletes the specified directory and all its contents.
func DeleteDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// GetDirContents returns a list of all files and directories in the specified directory.
func GetDirContents(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err!= nil {
		return nil, err
	}
	var contents []string
	for _, file := range files {
		contents = append(contents, file.Name())
	}
	return contents, nil
}

// GetFileExtension returns the file extension of the specified file.
func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// GetFileNameWithoutExtension returns the file name without the extension.
func GetFileNameWithoutExtension(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// IsFile returns true if the specified path is a file, false otherwise.
func IsFile(filePath string) bool {
	stat, err := os.Stat(filePath)
	return err == nil &&!stat.IsDir()
}

// IsDir returns true if the specified path is a directory, false otherwise.
func IsDir(dirPath string) bool {
	stat, err := os.Stat(dirPath)
	return err == nil && stat.IsDir()
}

// LoadJSON loads a JSON file into a Go struct.
func LoadJSON(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err!= nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// SaveJSON saves a Go struct to a JSON file.
func SaveJSON(filePath string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err!= nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// RunCommand runs a shell command and returns its output.
func RunCommand(command string) (string, error) {
	output, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err!= nil {
		return "", err
	}
	return string(output), nil
}

// RunCommandWithOutput runs a shell command and returns its output and error.
func RunCommandWithOutput(command string) (string, string, error) {
	output, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err!= nil {
		return "", string(output), err
	}
	return string(output), "", nil
}

func main() {
	fmt.Println("This is a test")
}