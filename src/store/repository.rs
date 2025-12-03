use diesel::{ExpressionMethods, QueryDsl, QueryResult, SqliteConnection};
use diesel_async::{RunQueryDsl, sync_connection_wrapper::SyncConnectionWrapper};

use crate::store::{
    models::{Event, EventMember, Member, NewEventMember, Role},
    schema,
};

impl Event {
    pub async fn get_by_id(
        conn: &mut SyncConnectionWrapper<SqliteConnection>,
        id: u64,
    ) -> QueryResult<Self> {
        schema::events::table
            .filter(schema::events::id.eq(id as i64))
            .first::<Event>(conn)
            .await
    }

    pub async fn get_ongoing(
        conn: &mut SyncConnectionWrapper<SqliteConnection>,
    ) -> QueryResult<Vec<Self>> {
        schema::events::table
            .filter(schema::events::ongoing.eq(true))
            .load(conn)
            .await
    }

    pub async fn get_members(
        &self,
        conn: &mut SyncConnectionWrapper<SqliteConnection>,
    ) -> QueryResult<Vec<EventMember>> {
        schema::event_members::table
            .filter(schema::event_members::event_id.eq(self.id))
            .load::<EventMember>(conn)
            .await
    }

    pub async fn get_member_count(
        &self,
        conn: &mut SyncConnectionWrapper<SqliteConnection>,
    ) -> QueryResult<i64> {
        schema::event_members::table
            .filter(schema::event_members::event_id.eq(self.id))
            .select(schema::event_members::member_id)
            .count()
            .get_result(conn)
            .await
    }

    pub async fn add_member(
        &self,
        conn: &mut SyncConnectionWrapper<SqliteConnection>,
        member: &Member,
        role: Role,
    ) -> QueryResult<()> {
        diesel::insert_into(schema::event_members::table)
            .values(&NewEventMember {
                event_id: self.id,
                member_id: member.id,
                role: role as i32,
            })
            .on_conflict_do_nothing()
            .execute(conn)
            .await?;

        Ok(())
    }
}
