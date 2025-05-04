package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Claude Desktop MCP Demo",
		"0.0.1",
		server.WithLogging(),
	)

	uuidTool := mcp.NewTool("uuid",
		mcp.WithDescription("Generate a UUID"),
	)

	s.AddTool(uuidTool, uuidToolHanlder)

	if err := server.ServeStdio(s); err != nil {
		log.Fatal(err)
	}
}

func uuidToolHanlder(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(uuid.New().String()), nil
}
