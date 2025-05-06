# AI Agent Demo Project

## Overview

This repository contains a project with multiple AI agent demos. It includes agents implemented with different technology stacks (Node.js, Python, Go).

## Project Structure

- `mastra-agent/`: A Node.js-based agent using the Mastra framework.
- `python-agent/`: A Python-based agent using the Google Cloud Agent Development Kit.
- `mcp/`: A component implemented in Go (currently a UUID generation function).

**Important Note:** Both the `mastra-agent` and `python-agent` rely on the Server-Sent Events (SSE) provided by the `mcp` component. Therefore, you must start the `mcp` service before running either agent.

## How to Run Each Project

### Mastra Agent (`mastra-agent/`)

**Required Tools:**
- Node.js (v22.15.0 recommended - `.volta` configuration present)
- npm

**Environment Setup:**
```bash
cd mastra-agent
npm install
```

**Execution:**
```bash
npm run dev
# or
mastra dev
```

### Python Agent (`python-agent/`)

**Required Tools:**
- Python 3.13
- pip

**Environment Setup:**
1. Create and activate a Python virtual environment:
   ```bash
   cd python-agent
   python -m venv .venv
   source .venv/bin/activate  # Linux/macOS
   # .venv\Scripts\activate  # Windows
   ```
2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

**Execution:**
An execution script that utilizes the `root_agent` in `agent.py` is needed. The specific execution method depends on the project's configuration.

### MCP (UUID Generator) (`mcp/uuid/`)

**Required Tools:**
- Go (appropriate version)

**Environment Setup:**
```bash
cd mcp/uuid
go mod tidy
```

**Execution:**
```bash
go run ./cmd/main.go
```
Or, build and then run:
```bash
go build -o uuid-generator ./cmd/main.go
./uuid-generator
``` 