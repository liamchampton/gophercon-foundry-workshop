## Step 2 - Give the Agent MCP Tools

Agent Framework can adapt tools discovered by an MCP client and execute them
automatically. You do not need to parse function calls or build a tool loop.

### Before You Start

A **credential** is information an application uses to prove that it is
allowed to call a service. In Step 1, your '.env' file provided the Foundry
credential in 'FOUNDRY_API_KEY'. The call to 'godotenv.Load()' loads that value
for the Go program.

You do not need to create or copy another credential in this step. The Cupcake
Store MCP endpoint used by this workshop does not require sign-in.

### 1. Add the MCP Imports

In 'main.go', find the 'import' block. Add these two imports with the other
third-party imports:

```go-notype
"github.com/microsoft/agent-framework-go/tool/mcptool"
mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
```

### 2. Add the Cupcake Store Address

Add this constant after the 'import' block and before 'func main()':

```go-notype
const cupcakeMCPURL = "https://ca-cupcake-mcp.jollyplant-ed217b0d.eastus.azurecontainerapps.io/mcp/"
```

This URL tells the MCP client where to find the Cupcake Store server.

### 3. Connect to MCP and Discover Its Tools

In 'main', find the code that creates 'client'. Immediately after the closing
')' of 'openai.NewClient(...)', and **before** the line that starts
'sparkles := openaiprovider.NewAgent(', add:

```go-notype
mcpSession, err := mcptool.Connect(ctx, &mcpsdk.StreamableClientTransport{
    Endpoint: cupcakeMCPURL,
})
if err != nil {
    fmt.Fprintln(os.Stderr, "failed to connect to MCP server:", err)
    os.Exit(1)
}
defer mcpSession.Close()

tools, err := mcptool.ListTools(ctx, mcpSession)
if err != nil {
    fmt.Fprintln(os.Stderr, "failed to list tools:", err)
    os.Exit(1)
}
```

'mcpSession' is the connection to the Cupcake Store. 'ListTools' asks the
server which actions it offers and converts them into tools Agent Framework
can use.

The order matters: the program must create 'tools' before it creates the agent.

### 4. Give the Tools to Sparkles

Find the 'openaiprovider.AgentConfig{...}' block inside
'openaiprovider.NewAgent'. Replace only that configuration block with:

```go-notype
openaiprovider.AgentConfig{
    Model:        os.Getenv("FOUNDRY_MODEL_DEPLOYMENT"),
    Instructions: "Help customers choose and order cupcakes. Use the available tools.",
    Config: agent.Config{
        Name:  "Sparkles",
        Tools: tools,
    },
},
```

Your 'main' function should now do these things in order:

1. Load settings from '.env'.
2. Create the Foundry model client.
3. Connect to the Cupcake Store MCP server and discover 'tools'.
4. Create Sparkles with those tools.
5. Create the chat session and start the input loop.

### 5. Run and Test the Agent

In the terminal, run:

```bash
go mod tidy && go run .
```

Try these prompts:

- 'What cupcake flavors are available?'
- 'Which are in stock?'
- 'Order one for me.'

The model decides when a tool is needed. Agent Framework calls the MCP tool,
sends its result back to the model, and returns the final response.

If Go reports 'undefined: tools', move the MCP connection and 'ListTools' code
above 'openaiprovider.NewAgent'.

---

✅ **In this step you have:** connected an MCP server, discovered its tools,
and given them to Agent Framework for automatic execution.

➡️ Click **Next** to load Sparkles' persona from MCP.
