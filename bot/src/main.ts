import { Client, GatewayIntentBits, Events } from "discord.js";
import { config } from "@utils/config";
import Hooks, { HookMethods } from "@hooks/server";
import { Server } from "http";
import GetEvents from "@hooks/events";

const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent,
    GatewayIntentBits.GuildMembers,
    GatewayIntentBits.GuildScheduledEvents,
  ],
});

let server: Server;

client.once(Events.ClientReady, (readyClient) => {
  console.log(`Ready! Logged in as ${readyClient.user.tag}`);
  const hooks = new Hooks(readyClient);

  hooks.addHook(HookMethods.Post, "/hooks/events", GetEvents);

  server = hooks.serve();
});

client.on(Events.MessageCreate, async (message) => {
  if (message.author.bot) return;
});

client.login(config.DISCORD_TOKEN).catch((error) => {
  console.error("Failed to login:", error);
  process.exit(1);
});

process.on("SIGINT", () => {
  console.log("Shutting down...");
  if (server) {
    server.close(() => {
      console.log("webhook server closed");
      client.destroy();
      process.exit(0);
    });
  } else {
    client.destroy();
    process.exit(0);
  }
});
