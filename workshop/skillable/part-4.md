## Part 4 - Build the Agent

Now that Foundry is set up, it's time to write some code with the public
preview of **Microsoft Agent Framework for Go**. The framework creates and
runs the agent, keeps conversation context, and automatically executes tools.
The **MCP Go SDK** connects it to the Cupcake Store MCP server.

You'll build the agent up over the next three steps - run it after each
step to see it grow.

### Setup

In Visual Studio Code, open a new terminal (**Terminal > New Terminal**).
It opens in 'c:\agents' - the folder where the workshop code lives.

The Go **dependencies are already downloaded** on the lab VM. For
reference, here's what's in 'go.mod':

- **agent-framework-go** - the agent, OpenAI-compatible provider, sessions, and MCP tool adapter.
- **openai-go** - connects to the Foundry model endpoint with an API key.
- **modelcontextprotocol/go-sdk** - the MCP client that connects to the
  Cupcake Store server.
- **godotenv** - loads your '.env' file into environment variables.

An empty **main.go** file is already waiting for you in 'c:\agents'.

---

✅ **In this step you have:** opened the VS Code terminal in 'c:\agents', learned about the installed Go packages, and located the empty 'main.go'.

➡️ Click **Next** to write your first Hello World agent.
