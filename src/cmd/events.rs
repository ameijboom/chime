use chrono::NaiveDateTime;

use crate::cmd::Context;

#[derive(Debug, poise::Modal)]
#[name = "Create FC Event"]
pub struct EventsCreationModal {
    title: String,
    #[paragraph]
    description: Option<String>,
}

#[poise::command(slash_command, subcommands("create", "list"), subcommand_required)]
pub async fn events(_: Context<'_>) -> anyhow::Result<()> {
    tracing::trace!("invoked events command");
    Ok(())
}

#[poise::command(slash_command)]
pub async fn create(
    ctx: Context<'_>,
    #[description = "Event start time"] start_time: NaiveDateTime,
    #[description = "Event end time"] end_time: NaiveDateTime,
) -> anyhow::Result<()> {
    use poise::Modal as _;

    let data = EventsCreationModal::execute(ctx).await?;
    ctx.say(format!("{:?}, {}, {}", data, start_time, end_time))
        .await?;

    Ok(())
}

#[poise::command(slash_command)]
pub async fn list(ctx: Context<'_>) -> anyhow::Result<()> {
    ctx.say("Listing Event").await?;
    Ok(())
}
