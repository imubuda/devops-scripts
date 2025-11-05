package devops_scripts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetDirSize(dir string, size map[string]float64) (float64, error) {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	for _, dir := range dirs {
		path := filepath.Join(dir.Name(), dir.Name())
		if dir.IsDir() {
			size[path], err = GetDirSize(path, size)
			if err != nil {
				return 0, err
			}
		} else {
			size[path] = float64(dir.Size())
		}
	}
	return GetDirSizeHelper(size, dir)
}

func GetDirSizeHelper(size map[string]float64, dir string) (float64, error) {
	size[dir] += 0
	for k, v := range size {
		size[dir] += v
		delete(size, k)
	}
	return size[dir], nil
}

func IsInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(b, a) {
			return true
		}
	}
	return false
}

func GetFileContents(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func IsDirEmpty(path string) (bool, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			return false, nil
		}
	}
	return true, nil
}

func GetFilePermissions(filename string) (os.FileMode, error) {
	return os.FileMode(os.Stat(filename).Mode()), nil
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetDirPath(filename string) string {
	return filepath.Dir(filename)
}

func GetDirName(filename string) string {
	return filepath.Base(filepath.Dir(filename))
}

func GetFileName(filename string) string {
	return filepath.Base(filename)
}

func humanizeBytes(n int64) string {
	units := []string{"", "K", "M", "G", "T", "P", "E", "Z", "Y"}
	l := float64(n)
	i := int64(0)
	for l >= 1024 && i < len(units) {
		l /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %sB", l, units[i])
}