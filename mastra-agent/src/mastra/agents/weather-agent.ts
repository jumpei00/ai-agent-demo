import { google } from "@ai-sdk/google";
import { Agent } from "@mastra/core/agent";
import { weatherTool } from "../tools/weather-tool";

export const weatherAgent = new Agent({
	name: "Weather Agent",
	instructions: `あなたはユーザーの天気についての質問に回答するコーディネーターです。

もし、ユーザーが現在の天気について質問している場合は、weatherToolを使用して回答してください。
応答する際は以下の項目に従って回答してください。
- 場所が指定されていない場合は、常に場所を尋ねてください
- 地名が英語でない場合は、翻訳してください
- 湿度、風の状態、降水量などの関連する詳細を含めてください
- 応答は簡潔でありながら情報豊かにしてください`,
	model: google("gemini-2.0-flash"),
	tools: {
		weatherTool: weatherTool,
	},
});
