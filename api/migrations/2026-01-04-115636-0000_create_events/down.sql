DROP TRIGGER IF EXISTS update_scheduled_event_members_updated_at ON scheduled_event_members;
DROP TRIGGER IF EXISTS update_members_updated_at ON members;
DROP TRIGGER IF EXISTS update_scheduled_events_updated_at ON scheduled_events;

DROP TABLE IF EXISTS scheduled_event_members CASCADE;
DROP TABLE IF EXISTS members CASCADE;
DROP TABLE IF EXISTS scheduled_events CASCADE;

DROP FUNCTION IF EXISTS update_updated_at_column();
