package utilities

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/HamzaAnis/go-merge-images/models"
)

func GetProcessedDirectories(path string) []models.Directory {
	directories := []models.Directory{}

	return directories
}

func GetFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileNames := []string{}
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, filepath.Join(currentPath, f.Name()))
			// fileNames = append(fileNames, f.Name())
		}
	}
	sort.Sort(sort.StringSlice(fileNames))
	return fileNames, nil
}

func GetDirectories(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	directoryNames := []string{}

	for _, f := range files {
		if f.IsDir() {
			directoryNames = append(directoryNames, filepath.Join(currentPath, f.Name()))
		}
	}

	return directoryNames, nil
}
