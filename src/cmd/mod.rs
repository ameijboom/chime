use poise::serenity_prelude as serenity;
use std::sync::Arc;

use crate::store::Database;

pub mod events;

pub type Context<'a> = poise::ApplicationContext<'a, Data, anyhow::Error>;

pub struct Data {
    pub db: Arc<Database>,
    pub channel: serenity::ChannelId,
}
