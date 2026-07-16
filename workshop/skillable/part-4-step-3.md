## Step 3 - Load the Persona from MCP

In Step 2, Sparkles learned what actions it can perform. Now it will download
two **prompts** from the same MCP server:

- 'agent_instructions' tells the model how Sparkles should behave and guide an
    order.
- 'welcome_banner' is text displayed in the terminal when the program starts.

An MCP prompt is reusable text supplied by an MCP server. It is not a question
that the customer types. No new credential or MCP connection is needed in this
step; use the 'mcpSession' you created in Step 2.

### 1. Add a Helper That Reads an MCP Prompt

Scroll to the end of 'main.go'. Add this function **after** the closing brace
of 'func main()':


```go-notype
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
```

The helper asks the MCP server for a prompt by name and combines its text into
one Go string.

### 2. Download Sparkles' Instructions and Banner

Inside 'main', find the block that starts with:

```go-notype
tools, err := mcptool.ListTools(ctx, mcpSession)
```

Add the following code immediately **after that block's error check** and
**before** 'sparkles := openaiprovider.NewAgent(':

```go-notype
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
```

The program must download these prompts before creating the agent because the
agent configuration uses the 'instructions' variable.

### 3. Apply the Downloaded Instructions

Inside 'openaiprovider.NewAgent', find this line:

```go-notype
Instructions: "Help customers choose and order cupcakes. Use the available tools.",
```

Replace that line with:

```go-notype
Instructions: instructions,
```

The complete configuration block should now look like this:

```go-notype
openaiprovider.AgentConfig{
    Model:        os.Getenv("FOUNDRY_MODEL_DEPLOYMENT"),
    Instructions: instructions,
    Config: agent.Config{
        Name:  "Sparkles",
        Tools: tools,
    },
},
```

### 4. Add a Chat Helper

Find the error check immediately after 'sparkles.CreateSession(ctx)':

```go-notype
if err != nil {
    fmt.Fprintln(os.Stderr, "failed to create agent session:", err)
    os.Exit(1)
}
```

Immediately after that error check, and before the line that prints
'Type exit to quit', add:

```go-notype
chat := func(input string) (string, error) {
    response, err := sparkles.RunText(ctx, input, agent.WithSession(session)).Collect()
    if err != nil {
        return "", err
    }
    return response.String(), nil
}
```

The 'chat' helper sends every message through the same Agent Framework
'session', preserving the conversation while avoiding repeated code.

### 5. Display the Banner and Start the Conversation

Find these two lines before the scanner is created:

```go-notype
fmt.Println("Type 'exit' to quit.")
scanner := bufio.NewScanner(os.Stdin)
```

Replace them with this entire block:

```go-notype
fmt.Println(banner)
fmt.Println("Type 'exit' to quit.")

reply, err := chat("hello")
if err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
} else {
    fmt.Printf("\033[1;35mAssistant:\033[0m\n%s\n\n", reply)
}

scanner := bufio.NewScanner(os.Stdin)
```

The banner is printed directly. Then 'chat("hello")' sends a first message to
the model so Sparkles greets the customer using its downloaded instructions.

### 6. Use the Chat Helper in the Input Loop

Near the bottom of the input loop, find this entire block:

```go-notype
response, err := sparkles.RunText(ctx, input, agent.WithSession(session)).Collect()
if err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    continue
}
fmt.Printf("\n\033[1;35mAssistant:\033[0m\n%s\n\n", response)
```

Replace the entire block, including the 'fmt.Printf' line, with:

```go-notype
reply, err := chat(input)
if err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    continue
}
fmt.Printf("\n\033[1;35mAssistant:\033[0m\n%s\n\n", reply)
```

Replacing the whole block is important because the response variable is now
named 'reply', not 'response'.

Your 'main' function should now do these things in order:

1. Connect to MCP and discover its tools.
2. Download 'agent_instructions' and 'welcome_banner'.
3. Create Sparkles with the downloaded instructions and MCP tools.
4. Create one Agent Framework session and the 'chat' helper.
5. Print the banner, ask Sparkles to greet the customer, and start the input
   loop.

### 7. Run the Finished Agent

In the terminal, run:

```bash
go mod tidy && go run .
```

You should first see the Cupcake Store banner, followed by a greeting from
Sparkles. Follow Sparkles' prompts to choose a flavor and place an order.

If Go reports 'undefined: instructions', move the prompt download code above
'openaiprovider.NewAgent'. If it reports 'undefined: response', replace the
entire input-loop block from Step 6, including its final 'fmt.Printf' line.

---

✅ **In this step you have:** loaded instructions and UI text from MCP, attached
MCP tools to a Foundry-backed Agent Framework agent, and reused one session for
the complete order conversation.

➡️ Click **Next** for the recap.
