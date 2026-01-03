export const config = {
  DISCORD_TOKEN: process.env.DISCORD_TOKEN || "",
  CLIENT_ID: process.env.CLIENT_ID || "",
  GUILD_ID: process.env.GUILD_ID || "",
  WEBHOOK_PORT: parseInt(process.env.WEBHOOK_PORT || "3000", 10),
} as const;

if (!config.DISCORD_TOKEN) {
  console.error("Missing DISCORD_TOKEN environment variable");
  process.exit(1);
}
