# Code with Claude - Microsoft Foundry Workshop

Build an AI agent with **Microsoft Foundry** and the **Microsoft Agent Framework** that uses a Claude model and calls tools through an MCP server.

## What you will build

A simple cupcake-ordering agent that:

- Uses **Claude Sonnet 4.6** deployed in Microsoft Foundry
- Follows a custom persona via a system prompt (`instructions.md`)
- Calls live tools from the **Cupcake Store MCP server**

## Repository layout

```
.
├── workshop/
│   └── workshop.md              # Step-by-step lab manual
├── agent-framework/             # Your working folder for the workshop
│   ├── .env                     # Fill in your Foundry endpoint + key here
│   └── requirements.txt         # Python dependencies
└── sample/
    ├── agent.py                 # Final reference implementation
    ├── instructions.md          # System prompt for the agent
    └── requirements.txt         # Python dependencies
```

## Prerequisites

- An Azure subscription with access to **Microsoft Foundry**
- A deployed Claude model (e.g. `claude-sonnet-4-6`)
- Python 3.10+

## Quick start

1. Follow the [workshop lab manual](workshop/workshop.md) to build the agent step-by-step in `agent-framework/`.
2. Or, to run the finished reference agent directly:

   ```bash
   cd sample
   pip install -r requirements.txt
   # Create a .env file with the variables listed below
   python agent.py
   ```

## Environment variables

Configured in `agent-framework/.env` (already present - just edit it):

| Variable | Description |
|---|---|
| `FOUNDRY_ENDPOINT` | Target URI of your Foundry deployment, e.g. `https://<resource>.services.ai.azure.com/anthropic` |
| `FOUNDRY_API_KEY` | API key for the Foundry deployment |
| `FOUNDRY_MODEL_DEPLOYMENT` | Deployment name (e.g. `claude-sonnet-4-6`) |
