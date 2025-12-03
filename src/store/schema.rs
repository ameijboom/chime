// @generated automatically by Diesel CLI.

diesel::table! {
    event_members (id) {
        id -> Integer,
        event_id -> BigInt,
        member_id -> BigInt,
        role -> Integer,
    }
}

diesel::table! {
    events (id) {
        id -> BigInt,
        message_id -> BigInt,
        ongoing -> Bool,
    }
}

diesel::table! {
    members (id) {
        id -> BigInt,
    }
}

diesel::joinable!(event_members -> events (event_id));
diesel::joinable!(event_members -> members (member_id));

diesel::allow_tables_to_appear_in_same_query!(event_members, events, members,);
