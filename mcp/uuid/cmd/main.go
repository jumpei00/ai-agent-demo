package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	mcpServer := server.NewMCPServer(
		"MCP Demo",
		"0.0.1",
		server.WithLogging(),
	)

	mcpServer.AddTool(mcp.NewTool("no_uuid", mcp.WithDescription("指定されたバージョンのUUIDが生成不可能な時に使用されるツールです。")), noUUIDToolHandler)
	mcpServer.AddTool(mcp.NewTool("uuidV1", mcp.WithDescription("V1のUUIDの生成が指定された時に使用されるツールです。")), uuidV1ToolHandler)
	mcpServer.AddTool(mcp.NewTool("uuidV4", mcp.WithDescription("V4のUUIDの生成が指定された時に使用されるツールです。")), uuidV4ToolHanlder)
	mcpServer.AddTool(mcp.NewTool("uuidV6", mcp.WithDescription("V6のUUIDの生成が指定された時に使用されるツールです。")), uuidV6ToolHandler)
	mcpServer.AddTool(mcp.NewTool("uuidV7", mcp.WithDescription("V7のUUIDの生成が指定された時に使用されるツールです。")), uuidV7ToolHandler)

	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://localhost:8080"))
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func noUUIDToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("CallToolRequest received in noUUIDToolHandler: ToolName=%s, Arguments=%v", request.Params.Name, request.Params.Arguments)
	return mcp.NewToolResultText("指定されたバージョンのUUIDは生成できません。"), nil
}

func uuidV1ToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("CallToolRequest received in uuidV1ToolHandler: ToolName=%s, Arguments=%v", request.Params.Name, request.Params.Arguments)
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(id.String()), nil
}

func uuidV4ToolHanlder(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("CallToolRequest received in uuidV4ToolHanlder: ToolName=%s, Arguments=%v", request.Params.Name, request.Params.Arguments)
	return mcp.NewToolResultText(uuid.New().String()), nil
}

func uuidV6ToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("CallToolRequest received in uuidV6ToolHandler: ToolName=%s, Arguments=%v", request.Params.Name, request.Params.Arguments)
	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(id.String()), nil
}

func uuidV7ToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("CallToolRequest received in uuidV7ToolHandler: ToolName=%s, Arguments=%v", request.Params.Name, request.Params.Arguments)
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(id.String()), nil
}
