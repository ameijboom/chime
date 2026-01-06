use chrono::NaiveDateTime;
use diesel::prelude::*;
use uuid::Uuid;

use crate::store::enums::{EventFrequency, EventStatus, PartyRole, PartyType};

#[derive(Queryable, Identifiable)]
#[diesel(table_name = crate::store::schema::scheduled_events)]
pub struct ScheduledEvent {
    pub id: Uuid,
    pub title: String,
    pub thumbnail: String,
    pub description: String,
    pub start_time: NaiveDateTime,
    pub end_time: NaiveDateTime,
    pub frequency: Option<EventFrequency>,
    pub party_type: Option<PartyType>,
    pub event_status: EventStatus,
    pub created_at: Option<NaiveDateTime>,
    pub updated_at: Option<NaiveDateTime>,
    pub deleted_at: Option<NaiveDateTime>,
}

#[derive(Insertable)]
#[diesel(table_name = crate::store::schema::scheduled_events)]
pub struct NewScheduledEvent {
    pub id: Uuid,
    pub title: String,
    pub thumbnail: String,
    pub description: String,
    pub start_time: NaiveDateTime,
    pub end_time: NaiveDateTime,
    pub frequency: Option<EventFrequency>,
    pub party_type: Option<PartyType>,
    pub event_status: EventStatus,
}

#[derive(Queryable, Identifiable)]
#[diesel(table_name = crate::store::schema::members)]
pub struct Member {
    pub id: String,
    pub created_at: Option<NaiveDateTime>,
    pub updated_at: Option<NaiveDateTime>,
    pub deleted_at: Option<NaiveDateTime>,
}

#[derive(Insertable)]
#[diesel(table_name = crate::store::schema::members)]
pub struct NewMember {
    pub id: String,
}

#[derive(Queryable, Identifiable, Associations)]
#[diesel(table_name = crate::store::schema::scheduled_event_members)]
#[diesel(belongs_to(Member, foreign_key = member))]
#[diesel(belongs_to(ScheduledEvent, foreign_key = event))]
#[diesel(primary_key(member, event))]
pub struct ScheduledEventMember {
    pub member: String,
    pub event: Uuid,
    pub party_role: Option<PartyRole>,
    pub created_at: Option<NaiveDateTime>,
    pub updated_at: Option<NaiveDateTime>,
    pub deleted_at: Option<NaiveDateTime>,
}

#[derive(Insertable)]
#[diesel(table_name = crate::store::schema::scheduled_event_members)]
pub struct NewScheduledEventMember {
    pub member: String,
    pub event: Uuid,
    pub party_role: Option<PartyRole>,
}
