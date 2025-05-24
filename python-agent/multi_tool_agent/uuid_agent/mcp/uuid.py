from google.adk.tools.mcp_tool.mcp_toolset import (
    MCPToolset,
    SseServerParams,
)


async def uuid_tools():
    tools, exit_stack = await MCPToolset.from_server(
        connection_params=SseServerParams(url="http://localhost:8080/sse"),
    )
    return tools, exit_stack
