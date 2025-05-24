from google.adk.agents import LlmAgent
from multi_tool_agent.uuid_agent.mcp.uuid import uuid_tools


async def uuid_agent():
    tools, exit_stack = await uuid_tools()

    uuid_agent = LlmAgent(
        name="uuid_agent",
        description="UUIDを生成するエージェントです",
        instruction="""
    あなたはUUIDを生成するエージェントです。
	利用可能なツールを利用して、ユーザーが指定したバージョンのUUIDを生成してください。
	利用可能なツールの中に、ユーザーが指定したバージョンのUUIDを生成するツールがない場合は、
	ユーザーに対して、指定されたバージョンのUUIDは生成できませんと返答してください。
	""",
        tools=tools,
    )

    return uuid_agent, exit_stack
