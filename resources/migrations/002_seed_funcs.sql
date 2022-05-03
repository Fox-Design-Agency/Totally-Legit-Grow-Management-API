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

CREATE TRIGGER new_plant_category_id
BEFORE INSERT ON plant_categories
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_plants_id
BEFORE INSERT ON plants
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_model_tasks_id
BEFORE INSERT ON model_tasks
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_generated_tasks_id
BEFORE INSERT ON generated_tasks
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_plant_life_cycles_id
BEFORE INSERT ON plant_life_cycles
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_plant_stages_id
BEFORE INSERT ON plant_stages
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_devices_id
BEFORE INSERT ON devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_groups_id
BEFORE INSERT ON growing_groups
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_group_devices_id
BEFORE INSERT ON growing_group_devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_locations_id
BEFORE INSERT ON growing_locations
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_location_devices_id
BEFORE INSERT ON growing_location_devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_levels_id
BEFORE INSERT ON growing_levels
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_growing_level_devices_id
BEFORE INSERT ON growing_level_devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_grow_spots_id
BEFORE INSERT ON grow_spots
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

CREATE TRIGGER new_grow_spots_devices_id
BEFORE INSERT ON grow_spots_devices
FOR EACH ROW EXECUTE PROCEDURE inject_uuid();

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP FUNCTION inject_uuid CASCADE;

DROP TRIGGER IF EXISTS new_plant_category_id ON plant_categories;
DROP TRIGGER IF EXISTS new_plants_id ON plants;
DROP TRIGGER IF EXISTS new_model_tasks_id ON model_tasks;
DROP TRIGGER IF EXISTS new_generated_tasks_id ON generated_tasks;
DROP TRIGGER IF EXISTS new_plant_life_cycles_id ON plant_life_cycles;
DROP TRIGGER IF EXISTS new_plant_stages_id ON plant_stages;
DROP TRIGGER IF EXISTS new_devices_id ON devices;
DROP TRIGGER IF EXISTS new_growing_groups_id ON growing_groups;
DROP TRIGGER IF EXISTS new_growing_group_devices_id ON growing_group_devices;
DROP TRIGGER IF EXISTS new_growing_locations_id ON growing_locations;
DROP TRIGGER IF EXISTS new_growing_location_devices_id ON growing_location_devices;
DROP TRIGGER IF EXISTS new_growing_levels_id ON growing_levels;
DROP TRIGGER IF EXISTS new_growing_level_devices_id ON growing_level_devices;
DROP TRIGGER IF EXISTS new_grow_spots_id ON grow_spots;
DROP TRIGGER IF EXISTS new_grow_spots_devices_id ON grow_spots_devices;