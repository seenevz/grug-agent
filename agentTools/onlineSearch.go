package agenttools

import (
	"encoding/json"
	"local-agent/utils"
)

var (
	browser *utils.BrowserNavigator
)

func init() {
	browser = utils.NewBrowserNavigator()
}

var OnlineSearchDefinition = ToolDefinition{
	Name:        "online_search",
	Description: "Perform an online search using a search engine. The search input can contain common search operators from Google Search. The return value contains the html string response from the request.",
	InputSchema: OnlineSearchInputSchema,
	Function:    OnlineSearch,
}

func OnlineSearch(input json.RawMessage) (string, error) {
	searchInput := OnlineSearchInput{}
	err := json.Unmarshal(input, &searchInput)
	if err != nil {
		return "", err
	}

	return *browser.PerformSearch(searchInput.SearchInput), nil

}

type OnlineSearchInput struct {
	SearchInput string `json:"searchInput" jsonschema_description:"The search input string to perform an online search using a search engine"`
}

var OnlineSearchInputSchema = GenerateSchema[OnlineSearchInput]()
