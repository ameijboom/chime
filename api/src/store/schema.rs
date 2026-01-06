// @generated automatically by Diesel CLI.

diesel::table! {
    members (id) {
        id -> Text,
        created_at -> Nullable<Timestamp>,
        updated_at -> Nullable<Timestamp>,
        deleted_at -> Nullable<Timestamp>,
    }
}

diesel::table! {
    scheduled_event_members (member, event) {
        member -> Text,
        event -> Uuid,
        party_role -> Nullable<Text>,
        created_at -> Nullable<Timestamp>,
        updated_at -> Nullable<Timestamp>,
        deleted_at -> Nullable<Timestamp>,
    }
}

diesel::table! {
    scheduled_events (id) {
        id -> Uuid,
        title -> Text,
        thumbnail -> Text,
        description -> Text,
        start_time -> Timestamp,
        end_time -> Timestamp,
        frequency -> Nullable<Text>,
        party_type -> Nullable<Text>,
        event_status -> Text,
        created_at -> Nullable<Timestamp>,
        updated_at -> Nullable<Timestamp>,
        deleted_at -> Nullable<Timestamp>,
    }
}

diesel::joinable!(scheduled_event_members -> members (member));
diesel::joinable!(scheduled_event_members -> scheduled_events (event));

diesel::allow_tables_to_appear_in_same_query!(members, scheduled_event_members, scheduled_events,);
