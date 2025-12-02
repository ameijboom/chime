use std::sync::Arc;

use diesel::SqliteConnection;
use diesel_async::{AsyncConnection, sync_connection_wrapper::SyncConnectionWrapper};
use tokio::sync::Mutex;

pub struct Database {
    connection: Arc<Mutex<SyncConnectionWrapper<SqliteConnection>>>,
}

impl Database {
    pub async fn new(url: &str) -> anyhow::Result<Self> {
        Ok(Self {
            connection: Arc::new(Mutex::new(
                SyncConnectionWrapper::<SqliteConnection>::establish(url).await?,
            )),
        })
    }

    pub fn get_connection(&self) -> Arc<Mutex<SyncConnectionWrapper<SqliteConnection>>> {
        Arc::clone(&self.connection)
    }
}
