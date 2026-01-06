CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE scheduled_events(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    thumbnail TEXT NOT NULL,
    description TEXT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    frequency TEXT CHECK (frequency IN ('weekly', 'bi-weekly', 'monthly', 'quarterly', 'yearly')),
    party_type TEXT CHECK (party_type IN ('light', 'full')),
    event_status TEXT NOT NULL CHECK (event_status IN ('ongoing', 'onhold', 'ended')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE members(
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE scheduled_event_members(
    member TEXT NOT NULL REFERENCES members(id) ON DELETE CASCADE,
    event UUID NOT NULL REFERENCES scheduled_events(id) ON DELETE CASCADE,
    party_role TEXT CHECK (party_role IN ('tank', 'healer', 'dps')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (member, event)
);

CREATE TRIGGER update_scheduled_events_updated_at
BEFORE UPDATE ON scheduled_events
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_members_updated_at
BEFORE UPDATE ON members
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_scheduled_event_members_updated_at
BEFORE UPDATE ON scheduled_event_members
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
