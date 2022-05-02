-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION inject_uuid() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.id IS NULL THEN
        NEW.id := gen_random_uuid();
    END IF;
	RETURN NEW;
END;$$ LANGUAGE plpgsql;
-- +migrate StatementEnd



-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION inject_uuid CASCADE;
