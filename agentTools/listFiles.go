package agenttools

import (
	"encoding/json"
	"local-agent/utils"
	"os"
)

var ListFilesDefinition = ToolDefinition{
	Name:        "list_files",
	Description: "List files and directories at a given path. If no path is provided, list files in the current directory.",
	InputSchema: ListFilesInputSchema,
	Function:    ListFiles,
}

type ListFilesInput struct {
	Path string `json:"path,omitempty" jsonschema_description:"Optional relative path to list files from. Defaults to current directory if not provided."`
}

var ListFilesInputSchema = GenerateSchema[ListFilesInput]()

func ListFiles(input json.RawMessage) (string, error) {
	listFilesInput := ListFilesInput{}
	err := json.Unmarshal(input, &listFilesInput)
	utils.CheckErr(err)

	root := "."

	if listFilesInput.Path != "" {
		root = listFilesInput.Path
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return "", err
	}

	var files []string
	for _, v := range entries {
		name := v.Name()
		if v.IsDir() {
			name += "/"
		}
		files = append(files, name)
	}

	parsedFiles, err := json.Marshal(files)
	if err != nil {
		return "", err
	}

	return string(parsedFiles), nil
}
