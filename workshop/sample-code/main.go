// Sparkles - The Cupcake ordering agent
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/microsoft/agent-framework-go/agent"
	"github.com/microsoft/agent-framework-go/provider/openaiprovider"
	"github.com/microsoft/agent-framework-go/tool/mcptool"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

const cupcakeMCPURL = "https://ca-cupcake-mcp.jollyplant-ed217b0d.eastus.azurecontainerapps.io/mcp/"

func main() {
	// 1. Load the Foundry settings from .env and create a shared context.
	_ = godotenv.Load()
	ctx := context.Background()

	// 2. Configure the OpenAI-compatible client for the Foundry model.
	client := openai.NewClient(
		option.WithBaseURL(os.Getenv("FOUNDRY_ENDPOINT")),
		option.WithAPIKey(os.Getenv("FOUNDRY_API_KEY")),
	)

	// 3. Connect to the Cupcake Store MCP server.
	mcpSession, err := mcptool.Connect(ctx, &mcpsdk.StreamableClientTransport{
		Endpoint: cupcakeMCPURL,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to connect to MCP server:", err)
		os.Exit(1)
	}
	defer mcpSession.Close()

	// 4. Discover the server's tools and adapt them for Agent Framework.
	tools, err := mcptool.ListTools(ctx, mcpSession)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to list tools:", err)
		os.Exit(1)
	}

	// 5. Download Sparkles' behavior and terminal banner from MCP prompts.
	instructions, err := promptText(ctx, mcpSession, "agent_instructions")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load agent_instructions prompt:", err)
		os.Exit(1)
	}
	banner, err := promptText(ctx, mcpSession, "welcome_banner")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load welcome_banner prompt:", err)
		os.Exit(1)
	}

	// 6. Create Sparkles with the downloaded instructions and discovered tools.
	sparkles := openaiprovider.NewAgent(
		client,
		openaiprovider.AgentConfig{
			Model:        os.Getenv("FOUNDRY_MODEL_DEPLOYMENT"),
			Instructions: instructions,
			Config: agent.Config{
				Name:  "Sparkles",
				Tools: tools,
			},
		},
	)

	// 7. Reuse one session so Sparkles remembers the complete conversation.
	session, err := sparkles.CreateSession(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create agent session:", err)
		os.Exit(1)
	}
	chat := func(input string) (string, error) {
		response, err := sparkles.RunText(ctx, input, agent.WithSession(session)).Collect()
		if err != nil {
			return "", err
		}
		return response.String(), nil
	}

	// 8. Show the store banner and ask Sparkles to greet the customer.
	fmt.Println(banner)
	fmt.Println("Type 'exit' to quit.")

	reply, err := chat("hello")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	} else {
		fmt.Printf("\033[1;35mAssistant:\033[0m\n%s\n\n", reply)
	}

	// 9. Read customer messages until they choose to exit.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[1;35mYou:\033[0m\n")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" || input == "quit" {
			break
		}

		reply, err := chat(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			continue
		}
		fmt.Printf("\n\033[1;35mAssistant:\033[0m\n%s\n\n", reply)
	}
}

// 10. Fetch an MCP prompt by name and combine all of its text content.
func promptText(ctx context.Context, session *mcpsdk.ClientSession, name string) (string, error) {
	result, err := session.GetPrompt(ctx, &mcpsdk.GetPromptParams{Name: name})
	if err != nil {
		return "", err
	}
	var text strings.Builder
	for _, msg := range result.Messages {
		if content, ok := msg.Content.(*mcpsdk.TextContent); ok {
			text.WriteString(content.Text)
		}
	}
	return text.String(), nil
}
