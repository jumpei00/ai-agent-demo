from google.adk.tools.mcp_tool.mcp_toolset import (
    AsyncExitStack,
    MCPToolset,
    SseServerParams,
)


async def uuid_tool() -> str:
    common_exit_stack = AsyncExitStack()

    tools, _ = await MCPToolset.from_server(
        connection_params=SseServerParams(url="http://localhost:8080/sse"),
        async_exit_stack=common_exit_stack,
    )

    return await tools[0].run_async(args={}, tool_context=None)
