import { MCPClient } from "@mastra/mcp";

export const uuidMcp = new MCPClient({
	servers: {
		uuid: {
			url: new URL("http://localhost:8080/sse"),
		},
	},
});
