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
import Command from "./command";

class Timestamp extends Command {
  init(
    name: string,
    description: string,
  ): RESTPostAPIChatInputApplicationCommandsJSONBody {
    return new SlashCommandBuilder()
      .setName(name)
      .setDescription(description)
      .addStringOption((option) =>
        option
          .setName("time")
          .setDescription("Any JavaScript-compatible datetime")
          .setRequired(true),
      )
      .toJSON();
  }

  async execute(interaction: ChatInputCommandInteraction): Promise<void> {
    let date = new Date(interaction.options.getString("time")!);

    return interaction.reply(`<t:${Math.floor(date.getTime() / 1000)}`);
  }
}
