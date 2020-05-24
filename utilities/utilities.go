package utilities

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/HamzaAnis/go-merge-images/models"
)

func GetProcessedDirectories(path string) ([]models.Directory, error) {
	directories := []models.Directory{}
	directoriesPath, err := GetDirectories(path)
	if err != nil {
		return nil, err
	}
	for _, directory := range directoriesPath {
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

	fileNames := []string{}
	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".png") {
			fileNames = append(fileNames, filepath.Join(path, f.Name()))
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

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func AskForConfirmation() bool {
	fmt.Print("Do you want to start the processing (Y/N)?")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return AskForConfirmation()
	}
}

// You might want to put the following two functions in a separate utility package.

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func DeleteFile(path string) error {
	log.Println("Deleting", path)
	err := os.Remove(path)

	if err != nil {
		return err
	}
	return nil
}
