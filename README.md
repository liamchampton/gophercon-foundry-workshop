# Microsoft Agent Framework for Go - Foundry Workshop

Build an AI agent in Go with **Microsoft Foundry** that uses a GPT model and calls tools through an MCP server, using [Microsoft Agent Framework for Go](https://github.com/microsoft/agent-framework-go).

## What you will build

An agent for **Sparkles**, a friendly cupcake shop, that:

- Uses a **GPT model** (e.g. `gpt-5.5`) deployed in Microsoft Foundry
- Loads its persona and welcome banner from MCP **prompts**
- Calls live tools from the **Cupcake Store MCP server**

> **Microsoft Foundry** is Microsoft's hosted platform for deploying AI models (OpenAI, Anthropic, Mistral, and more) and exposing them via an endpoint and key.
>
> **MCP** (Model Context Protocol) is an open standard that lets agents discover and call tools, prompts, and resources from a remote server over HTTP.

## Repository layout

```
.
├── workshop/
│   ├── workshop.md              # Step-by-step lab manual
│   ├── skillable.md             # Skillable lab instructions
│   └── sample-code/             # Final reference implementation
│       ├── main.go
│       ├── go.mod / go.sum
│       └── .env.sample
└── sparkles-agent/              # Your working folder for the workshop
    ├── .env                     # Fill in your model endpoint, API key, and deployment name
    ├── main.go                  # Empty starter - you build it up during the workshop
    └── go.mod / go.sum          # Go dependencies
```

## Prerequisites

- An Azure subscription with access to **Microsoft Foundry**
- A deployed GPT model (e.g. `gpt-5.5`)
- Go 1.25+

## Quick start

1. Follow the [workshop lab manual](workshop/workshop.md) to build the agent step-by-step in `sparkles-agent/`.
2. Or, to run the finished reference agent in [workshop/sample-code/](workshop/sample-code/) directly:

   ```bash
   cd workshop/sample-code
   go mod download
   cp .env.sample .env   # then edit .env with the variables listed below
   go run .
   ```

## Environment variables

Configured in `sparkles-agent/.env` (already present - just edit it):

| Variable | Description |
| --- | --- |
| `FOUNDRY_ENDPOINT` | Model target URI, e.g. `https://<resource>.services.ai.azure.com/openai/v1` |
| `FOUNDRY_API_KEY` | API key for the model deployment |
| `FOUNDRY_MODEL_DEPLOYMENT` | Model deployment name (e.g. `gpt-5.5`) |

## Troubleshooting

See the [Troubleshooting section](workshop/workshop.md#troubleshooting) of the workshop for common issues and fixes.
