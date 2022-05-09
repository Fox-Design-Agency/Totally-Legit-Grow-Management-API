-- Copyright 2022 Fox Design Agency. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

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

CREATE TRIGGER new_generated_tasks_id
BEFORE INSERT ON generated_tasks
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_devices_id
BEFORE INSERT ON devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_groups_id
BEFORE INSERT ON growing_groups
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_locations_id
BEFORE INSERT ON growing_locations
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_levels_id
BEFORE INSERT ON growing_levels
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_grow_spots_id
BEFORE INSERT ON grow_spots
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION inject_uuid CASCADE;

DROP TRIGGER IF EXISTS new_generated_tasks_id ON generated_tasks;
DROP TRIGGER IF EXISTS new_devices_id ON devices;
DROP TRIGGER IF EXISTS new_growing_groups_id ON growing_groups;
DROP TRIGGER IF EXISTS new_growing_locations_id ON growing_locations;
DROP TRIGGER IF EXISTS new_growing_levels_id ON growing_levels;
DROP TRIGGER IF EXISTS new_grow_spots_id ON grow_spots;