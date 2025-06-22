package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

type GetUserInput func() (string, bool)

func ScanUserInput() *GetUserInput {
	scanner := bufio.NewScanner(os.Stdin)
	var fn GetUserInput
	fn = func() (string, bool) {
		if scanner.Scan() {
			return scanner.Text(), true
		}

		return "", false
	}

	return &fn
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

func LoadApiKey() (string, error) {
	userHome, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	apiKeyFile, err := os.Open(path.Join(userHome, ".grug-key"))

	if err != nil {
		return "", err
	}

	key, err := io.ReadAll(apiKeyFile)

	if err != nil {
		return "", err
	}

	return string(key), nil
}
