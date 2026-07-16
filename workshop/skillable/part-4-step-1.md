## Step 1 - Hello World Agent

Create a Microsoft Foundry-backed agent and reuse one Agent Framework session
for the whole conversation.

Replace 'main.go' with:

```go-notype
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
    "github.com/openai/openai-go/v3"
    "github.com/openai/openai-go/v3/option"
)

func main() {
    _ = godotenv.Load()
    ctx := context.Background()

    client := openai.NewClient(
        option.WithBaseURL(os.Getenv("FOUNDRY_ENDPOINT")),
        option.WithAPIKey(os.Getenv("FOUNDRY_API_KEY")),
    )

    sparkles := openaiprovider.NewAgent(
        client,
        openaiprovider.AgentConfig{
            Model: os.Getenv("FOUNDRY_MODEL_DEPLOYMENT"),
            Config: agent.Config{Name: "Sparkles"},
        },
    )
    session, err := sparkles.CreateSession(ctx)
    if err != nil {
        fmt.Fprintln(os.Stderr, "failed to create agent session:", err)
        os.Exit(1)
    }

    fmt.Println("Type 'exit' to quit.")
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

        response, err := sparkles.RunText(ctx, input, agent.WithSession(session)).Collect()
        if err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
            continue
        }
        fmt.Printf("\n\033[1;35mAssistant:\033[0m\n%s\n\n", response)
    }
}
```

Run it:

```bash
go mod tidy && go run .
```

Send two related messages. 'agent.WithSession(session)' lets Agent Framework
carry the first turn into the second.

---

✅ **In this step you have:** connected to a Foundry model with its endpoint
and API key, created an agent, and held a multi-turn conversation.

➡️ Click **Next** to connect MCP tools.
