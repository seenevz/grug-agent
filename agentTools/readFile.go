package agenttools

import (
	"encoding/json"
	"log"
	"os"
)

var ReadFileDefinition = ToolDefinition{
	Name:        "read_file",
	Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
	InputSchema: ReadFileInputSchema,
	Function:    ReadFile,
}

func ReadFile(input json.RawMessage) (string, error) {
	readFileInput := ReadFileInput{}
	err := json.Unmarshal(input, &readFileInput)
	if err != nil {
		log.Panic(err)
	}

	content, err := os.ReadFile(readFileInput.Path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

type ReadFileInput struct {
	Path string `json:"path" jsonschema_descriptiton:"The relative path of a file in the working directory."`
}

var ReadFileInputSchema = GenerateSchema[ReadFileInput]()
