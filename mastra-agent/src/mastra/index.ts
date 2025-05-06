import { Mastra } from "@mastra/core";
import { uuidAgent } from "./agents/uuid-agent";
import { weatherAgent } from "./agents/weather-agent";

export const mastra = new Mastra({
	agents: { weatherAgent: weatherAgent, uuidAgent: uuidAgent },
});
