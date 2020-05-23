package utilities

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/HamzaAnis/go-merge-images/models"
)

func GetProcessedDirectories(path string) ([]models.Directory, error) {
	directories := []models.Directory{}
	directoriesPath, err := GetDirectories(path)
	if err != nil {
		return nil, err
	}
	for _, directory := range directoriesPath {
		log.Println("Processing ", directory)

		files, err := GetFiles(directory)
		if err != nil {
			return nil, err
		}
		directory := models.Directory{
			DirectoryPath: directory,
			Files:         files,
		}
		directories = append(directories, directory)
	}
	return directories, nil
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
