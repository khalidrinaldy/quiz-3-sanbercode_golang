-- +migrate Up
-- +migrate StatementBegin

SET timezone = 'Asia/Jakarta';

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON category
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +migrate StatementEnd