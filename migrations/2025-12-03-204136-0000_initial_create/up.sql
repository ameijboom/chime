CREATE TABLE events (
    id          BIGINT  PRIMARY KEY NOT NULL,
    message_id  BIGINT  NOT NULL,
    ongoing     BOOLEAN NOT NULL
);

CREATE TABLE members (
    id  bigint  PRIMARY KEY NOT NULL
);

CREATE TABLE event_members (
    id          INTEGER PRIMARY KEY AUTOINCREMENT   NOT NULL,
    event_id    BIGINT  REFERENCES  events(id)      NOT NULL,
    member_id   BIGINT  REFERENCES  members(id)     NOT NULL,
    role        INTEGER NOT NULL
);
