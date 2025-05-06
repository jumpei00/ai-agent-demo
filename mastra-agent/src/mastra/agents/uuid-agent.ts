import { google } from "@ai-sdk/google";
import { Agent } from "@mastra/core/agent";
import { uuidMcp } from "../tools/mcp";

export const uuidAgent = new Agent({
	name: "UUID Agent",
	instructions: `あなたはユーザーのUUIDについての質問に回答するコーディネーターです。
 
もし、ユーザーがUUIDの生成を行いたい場合は、uuidToolを使用し回答してください。`,
	model: google("gemini-2.0-flash"),
	tools: await uuidMcp.getTools(),
});
