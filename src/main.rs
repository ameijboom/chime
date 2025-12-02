use std::sync::Arc;

use clap::Parser;
use poise::serenity_prelude as serenity;
use tracing_subscriber::EnvFilter;

use crate::{cmd::Data, store::Database};

mod cmd;
mod store;

#[derive(Parser)]
struct Opts {
    #[arg(long, env = "DISCORD_TOKEN")]
    pub discord_token: String,

    #[arg(long, env = "DISCORD_CHANNEL")]
    pub discord_channel: String,

    #[arg(long, env = "DATABASE_URL")]
    pub database_url: String,
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let opts = Opts::parse();

    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::from_default_env())
        .init();

    let db = Database::new(&opts.database_url).await?;

    let intents = serenity::GatewayIntents::non_privileged()
        | serenity::GatewayIntents::MESSAGE_CONTENT
        | serenity::GatewayIntents::GUILD_MEMBERS;

    let options = poise::FrameworkOptions {
        commands: vec![cmd::events::events()],
        // event_handler: |ctx, event, framework, data| {},
        ..Default::default()
    };

    let framework = poise::Framework::builder()
        .setup(move |ctx, _ready, framework| {
            Box::pin(async move {
                poise::builtins::register_globally(ctx, &framework.options().commands).await?;
                Ok(Data {
                    db: Arc::new(db),
                    channel: serenity::ChannelId::new(opts.discord_channel.parse()?),
                })
            })
        })
        .options(options)
        .build();

    let mut client = serenity::ClientBuilder::new(opts.discord_token, intents)
        .framework(framework)
        .await?;

    client.start().await.map_err(|e| {
        tracing::error!({error = ?e}, "error starting client");
        anyhow::anyhow!(e)
    })
}
