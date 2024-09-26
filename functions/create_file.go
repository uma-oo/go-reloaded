package reloaded

import (
	"errors"
	"os"
)

func CreateFile(name_file string, directory_path string) (*os.File, error) {
	// Create the directory if it does not exist
	// mkdirall here is used to handle if we give dirs/dir and create them 
	if err := os.MkdirAll(directory_path, os.ModePerm); err != nil {
		return nil, errors.New("directory not found")
	}

	filepath := directory_path + "/" + name_file
	newfile, err := os.Create(filepath)
	if err != nil {
		return nil, errors.New("failed to create the file")
	}
	

	return newfile, nil
}
