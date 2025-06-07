package utils

import (
	"fmt"
	"log"
	"os"
	"path"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

func CreateNewFile(filePath, content string) (string, error) {
	dir := path.Dir(filePath)

	if dir != "." {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}

	return fmt.Sprintf("successfully created file: %s", filePath), nil
}
