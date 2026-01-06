import {
  ChatInputCommandInteraction,
  CommandInteraction,
  CommandInteractionOption,
  CommandInteractionOptionResolver,
  Interaction,
  InteractionResponse,
  RESTPostAPIChatInputApplicationCommandsJSONBody,
  SlashCommandBuilder,
} from "discord.js";

export default abstract class Command {
  abstract init(
    name: string,
    description: string,
  ): RESTPostAPIChatInputApplicationCommandsJSONBody;

  abstract execute(interaction: ChatInputCommandInteraction): Promise<void>;
}
