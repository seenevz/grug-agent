package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"local-agent/agentTools"
	"local-agent/tui"
	"local-agent/utils"
	"log"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

//go:embed .key
var ANTHROPIC_AGENT_KEY string

//go:embed systemPrompt
var SYSTEM_PROMPT string

func init() {
	if len(ANTHROPIC_AGENT_KEY) == 0 {
		log.Fatal("Anthropic key is missing")
	}
}

type Agent struct {
	*tui.TUI
	client           *anthropic.Client
	model            anthropic.Model
	systemPrompt     []anthropic.TextBlockParam
	getUserMessage   utils.GetUserInput
	conversation     []anthropic.MessageParam
	tools            []anthropic.ToolUnionParam
	toolsDefinitions []agenttools.ToolDefinition
}

func (a *Agent) runInference(ctx context.Context) (*anthropic.Message, error) {
	return a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     a.model,
		MaxTokens: 1024,
		System:    a.systemPrompt,
		Messages:  a.conversation,
		Tools:     a.tools,
	})
}

func (a *Agent) executeTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	var toolDef agenttools.ToolDefinition

	for _, tool := range a.toolsDefinitions {
		if tool.Name == name {
			toolDef = tool
			break
		}
	}

	if toolDef.IsEmpty() {
		return anthropic.NewToolResultBlock(id, "tool not found", true)
	}

	a.PrintTool(name, input)
	toolResponse, err := toolDef.Function(input)

	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	return anthropic.NewToolResultBlock(id, toolResponse, false)
}

func (a *Agent) Run(ctx context.Context) error {
	readUserInput := true

	a.PrintMessage("Chat with Grug (use ctrl-c to quit)")

	for {
		if readUserInput {

			userInput, ok := a.GetUserInput()

			if !ok {
				break
			}

			userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
			a.conversation = append(a.conversation, userMessage)

		}

		responseMessage, err := a.runInference(ctx)

		if err != nil {
			return err
		}

		a.conversation = append(a.conversation, responseMessage.ToParam())

		toolsResults := []anthropic.ContentBlockParamUnion{}
		for _, content := range responseMessage.Content {
			switch content.Type {
			case "text":
				a.PrintAgent(content.Text)
			case "tool_use":
				result := a.executeTool(content.ID, content.Name, content.Input)
				toolsResults = append(toolsResults, result)
			}
		}

		if len(toolsResults) == 0 {
			readUserInput = true
			continue
		}

		readUserInput = false
		a.conversation = append(a.conversation, anthropic.NewUserMessage(toolsResults...))
	}

	return nil
}

func NewAgent(client *anthropic.Client, tui *tui.TUI, tools []agenttools.ToolDefinition) *Agent {
	systemPrompt := []anthropic.TextBlockParam{{Text: SYSTEM_PROMPT}}
	conversation := []anthropic.MessageParam{}
	anthropicTools := []anthropic.ToolUnionParam{}

	for _, tool := range tools {
		anthropicTools = append(anthropicTools, anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
	}

	return &Agent{
		client:           client,
		model:            anthropic.ModelClaudeSonnet4_0,
		systemPrompt:     systemPrompt,
		conversation:     conversation,
		tools:            anthropicTools,
		toolsDefinitions: tools,
		TUI:              tui,
	}
}

func main() {
	client := anthropic.NewClient(option.WithAPIKey(strings.Trim(ANTHROPIC_AGENT_KEY, " \n\r")))
	tui := tui.New()
	agentTools := []agenttools.ToolDefinition{agenttools.ReadFileDefinition, agenttools.ListFilesDefinition, agenttools.EditFileDefinition}
	agent := NewAgent(&client, tui, agentTools)

	err := agent.Run(context.Background())

	utils.CheckErr(err)
}
