use diesel::{expression::AsExpression, prelude::*, sql_types::Integer};

#[derive(Debug, Clone, Copy, AsExpression)]
#[diesel(sql_type = Integer)]
pub enum Role {
    Tank,
    Healer,
    Dps,
    Invalid,
}

impl From<i32> for Role {
    fn from(value: i32) -> Self {
        match value {
            0 => Role::Tank,
            1 => Role::Healer,
            2 => Role::Dps,
            _ => Role::Invalid,
        }
    }
}

#[derive(Insertable, Queryable, Selectable)]
#[diesel(table_name = crate::store::schema::events)]
pub struct Event {
    pub id: i64,
    pub message_id: i64,
    pub ongoing: bool,
}

#[derive(Insertable, Queryable, Selectable)]
#[diesel(table_name = crate::store::schema::members)]
pub struct Member {
    pub id: i64,
}

#[derive(Insertable)]
#[diesel(belongs_to(Event))]
#[diesel(belongs_to(Member))]
#[diesel(table_name = crate::store::schema::event_members)]
pub struct NewEventMember {
    pub event_id: i64,
    pub member_id: i64,
    pub role: i32,
}

#[derive(Selectable, Debug, Queryable)]
#[diesel(belongs_to(Event))]
#[diesel(belongs_to(Member))]
#[diesel(table_name = crate::store::schema::event_members)]
pub struct EventMember {
    pub id: i32,
    pub event_id: i64,
    pub member_id: i64,
    pub role: i32,
}
