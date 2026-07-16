## Recap

Nice work - you just shipped a working AI agent. In a small Go program, you
built something that:

- ✅ Talks to a **GPT model** deployment hosted on **Microsoft
  Foundry** through its OpenAI-compatible endpoint and API key.
- ✅ Calls **live tools** through an **MCP server** - listing flavors,
  checking stock, and placing real orders against the Cupcake Store
  backend.
- ✅ Pulls its **persona** and **welcome banner** straight from the MCP
  server's **prompts**, so the same code becomes a different agent the
  moment the server changes its mind.

The pattern you just used - **Foundry for the model, Agent Framework for the
agent loop and session, and MCP for tools and prompts** - is exactly the
one you'd use to build a customer-support bot, an internal company
helper, a coding assistant, or pretty much any agent that needs both a
brain and hands. Swap the MCP server for a different one, point at a
different model, tweak the persona, and you have an entirely new agent
without rewriting the plumbing.

### Where to go next

Now that the agent works, here are a few directions to push it:

- **Add another MCP server.** Use 'mcptool.Connect' for a second server and
  merge its tools into the same slice passed to 'agent.Config'. Try
  giving the agent a weather server, a calendar, or a search
  tool alongside the cupcake store and watch it pick the right tool
  for each question.
- **Swap the model.** Deploy a different model on Foundry,
  change 'FOUNDRY_MODEL_DEPLOYMENT' in your '.env', and compare personality, latency,
  and tool-calling style.
- **Stream the responses.** Iterate over 'RunText' with 'agent.Stream(true)'
  to print updates as they arrive.
- **Persist the session.** Marshal the 'agent.Session' to JSON and restore it
  later so the agent can continue a customer's earlier conversation.
- **Add observability.** Register Agent Framework middleware and OpenTelemetry
  tracing to inspect model and tool activity.
- **Build your own MCP server.** Now that you've consumed one, writing
  one is the natural next step - and the moment you do, every agent
  that speaks MCP (yours or someone else's) can use your tools.

**Thanks for building with us - now go eat your cupcake. 🧁**