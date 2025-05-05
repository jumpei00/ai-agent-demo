from google.adk.agents import LlmAgent
from multi_tool_agent.weather_agent.agent import weather_time_agent
from multi_tool_agent.uuid_agent.agent import uuid_agent

root_agent = LlmAgent(
    name="Coordinator",
    model="gemini-2.0-flash",
    instruction=(
        "ユーザーの質問に対して適切なエージェントを選択して回答してください"
        "天気と時刻についての質問はweather_time_agentを使用してください"
        "UUIDについての質問はuuid_agentを使用してください"
    ),
    description="ユーザーの質問に回答するためのコーディネーターです",
    sub_agents=[weather_time_agent, uuid_agent],
)
