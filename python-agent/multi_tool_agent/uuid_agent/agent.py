from google.adk.agents import LlmAgent
from multi_tool_agent.uuid_agent.mcp.uuid import uuid_tool

uuid_agent = LlmAgent(
    name="uuid_agent",
    description="UUIDを生成するエージェントです",
    instruction="UUIDを生成するエージェントです",
    tools=[uuid_tool],
)
