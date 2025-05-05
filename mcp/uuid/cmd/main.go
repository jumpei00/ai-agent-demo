package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	uuidTool := mcp.NewTool("uuid",
		mcp.WithDescription("Generate a UUID"),
	)

	mcpServer := server.NewMCPServer(
		"MCP Demo",
		"0.0.1",
		server.WithLogging(),
	)

	mcpServer.AddTool(uuidTool, uuidToolHanlder)

	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://localhost:8080"))
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func uuidToolHanlder(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(uuid.New().String()), nil
}
