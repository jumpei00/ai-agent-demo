import datetime
from zoneinfo import ZoneInfo
from google.adk.agents import LlmAgent


def get_weather(city: str) -> dict:
    if city.lower() == "東京":
        return {
            "status": "success",
            "report": ("東京の天気は晴れで、気温は25度です。"),
        }
    return {
        "status": "error",
        "error_message": f"Weather information for '{city}' is not available.",
    }


def get_current_time(city: str) -> dict:
    if city.lower() == "東京":
        tz = ZoneInfo("Asia/Tokyo")
        now = datetime.datetime.now(tz)
        report = f"東京の現在時刻は{now.strftime('%Y-%m-%d %H:%M:%S %Z%z')}"
        return {"status": "success", "report": report}

    return {
        "status": "error",
        "error_message": (f"Sorry, I don't have timezone information for {city}."),
    }


weather_time_agent = LlmAgent(
    name="weather_time_agent",
    description=("都市の時刻と天気についての質問に回答するエージェントです"),
    instruction=("都市の時刻と天気についての質問に回答するエージェントです"),
    tools=[get_weather, get_current_time],
)
