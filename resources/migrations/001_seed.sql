-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE TIMEMEASURE AS ENUM('minutes', 'hours', 'days', 'weeks', 'months', 'years');

SET TIME ZONE 'UTC';

-- members
CREATE TABLE members(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- organizations
CREATE TABLE organizations(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);


-- organization_members
CREATE TABLE organization_members(
    organization INT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    member INT NOT NULL REFERENCES members(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (organization, member)
);

-- plant categories
CREATE TABLE plant_categories(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plants
CREATE TABLE plants(
    id SERIAL PRIMARY KEY,
    plant_category INT NOT NULL REFERENCES plant_categories(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);


-- model_tasks
CREATE TABLE model_tasks(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    info TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- generated_tasks
CREATE TABLE generated_tasks(
    id UUID PRIMARY KEY,
    task INT NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    completed TIMESTAMPTZ,
    completed_by INT REFERENCES members(id) ON DELETE CASCADE,
    has_issue BOOLEAN DEFAULT FALSE,
    reason TEXT,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plant life cycles
CREATE TABLE plant_life_cycles(
    id SERIAL PRIMARY KEY,
    actual_time_measure NUMERIC,
    display_name VARCHAR(256) NOT NULL,
    total_time_measure_units TIMEMEASURE NOT NULL,
    est_total_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plant stages
CREATE TABLE plant_stages(
    id SERIAL PRIMARY KEY,
    actual_time_measure NUMERIC,
    display_name VARCHAR(256) NOT NULL,
    time_measure_units TIMEMEASURE NOT NULL,
    est_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    completed TIMESTAMPTZ,
    completed_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

CREATE TABLE nutrients(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

CREATE TABLE plant_stage_care(
    id SERIAL PRIMARY KEY,
    plant_stage INT NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    -- water (will generate task)
    actual_water_time_measure NUMERIC,
    water_time_measure_units TIMEMEASURE NOT NULL,
    est_water_time_measure INT NOT NULL,
    -- temperature (can generate task)
    temperature_min INT,
    temperature_max INT,
    -- light (can generate task)
    light_time_measure_hours INT,
    -- observation (will generate task)
    actual_observation_time_measure NUMERIC,
    observation_time_measure_units TIMEMEASURE NOT NULL,
    est_observation_time_measure INT NOT NULL,
    -- time measure and 
    actual_time_measure NUMERIC,
    time_measure_units TIMEMEASURE NOT NULL,
    est_time_measure INT NOT NULL,
    -- over observation time (will generate task)
    over_observation_time_measure_units TIMEMEASURE NOT NULL,
    over_est_observation_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

CREATE TABLE plant_stage_care_nutrients(
    nutrient INT NOT NULL REFERENCES nutrients(id) ON DELETE CASCADE,
    plant_stage INT NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    amount_units VARCHAR(16),
    amount INT,
    -- might make enum for this for trigger events
    when_to_add VARCHAR(64), 
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (nutrient, plant_stage)
);

-- plant life cycle stages
CREATE TABLE plant_life_cycle_stages(
    plant_life_cycle INT NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    plant_stage INT NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    stage_num INT NOT NULL,
    prev_stage INT REFERENCES plant_stages(id) ON DELETE CASCADE,
    next_stage INT REFERENCES plant_stages(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (plant_life_cycle, plant_stage)
);

-- plant stage tasks
CREATE TABLE plant_stage_tasks(
    model_task INT NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_stage INT NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    actual_time_measure NUMERIC,
    time_measure_units TIMEMEASURE NOT NULL,
    est_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_stage)
);

-- plant life cycle start tasks
CREATE TABLE plant_life_cycle_start_tasks(
    model_task INT NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_life_cycle INT NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    actual_time_measure NUMERIC,
    time_measure_units TIMEMEASURE NOT NULL,
    est_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_life_cycle)
);

-- plant life cycle end tasks
CREATE TABLE plant_life_cycle_end_tasks(
    model_task INT NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_life_cycle INT NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    actual_time_measure NUMERIC,
    time_measure_units TIMEMEASURE NOT NULL,
    est_time_measure INT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_life_cycle)
);

-- devices (lighting, water, sensors, etc)
CREATE TABLE devices(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

CREATE TABLE device_actions(
    id SERIAL PRIMARY KEY,
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    invoke_route VARCHAR(512) NOT NULL,
    info TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing groups
CREATE TABLE growing_groups(
    id UUID PRIMARY KEY,
    organization INT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing group devices
CREATE TABLE growing_group_devices(
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_group UUID NOT NULL REFERENCES growing_groups(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP,
    PRIMARY KEY (device, growing_group)
);

-- growing locations
CREATE TABLE growing_locations(
    id UUID PRIMARY KEY,
    growing_group UUID NOT NULL REFERENCES growing_groups(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing location devices
CREATE TABLE growing_location_devices(
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_location UUID NOT NULL REFERENCES growing_locations(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP,
    PRIMARY KEY (device, growing_location)
);

-- growing levels
CREATE TABLE growing_levels(
    id UUID PRIMARY KEY,
    growing_location UUID NOT NULL REFERENCES growing_locations(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    level_num INT DEFAULT 1,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing level devices
CREATE TABLE growing_level_devices(
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_level UUID NOT NULL REFERENCES growing_levels(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP,
    PRIMARY KEY (device, growing_level)
);

-- grow spots
CREATE TABLE grow_spots(
    id UUID PRIMARY KEY,
    growing_level UUID NOT NULL REFERENCES growing_levels(id) ON DELETE CASCADE,
    is_individual BOOLEAN DEFAULT TRUE, 
    is_companion BOOLEAN DEFAULT FALSE,
    num_spots INT DEFAULT 1,
    spot_type VARCHAR(64),
    spot_size VARCHAR(16),
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- grow spot devices
CREATE TABLE grow_spots_devices(
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    grow_spot UUID NOT NULL REFERENCES grow_spots(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INT REFERENCES members(id) ON DELETE CASCADE,
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by INT REFERENCES members(id) ON DELETE CASCADE,
    meta JSON DEFAULT '{}',
    archived TIMESTAMP,
    PRIMARY KEY (device, grow_spot)
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TYPE IF EXISTS TIMEMEASURE CASCADE;

DROP TABLE IF EXISTS organizations CASCADE;
DROP TABLE IF EXISTS members CASCADE;
DROP TABLE IF EXISTS organization_members CASCADE;
DROP TABLE IF EXISTS plant_categories CASCADE;
DROP TABLE IF EXISTS plants CASCADE;
DROP TABLE IF EXISTS model_tasks CASCADE;
DROP TABLE IF EXISTS generated_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycles CASCADE;
DROP TABLE IF EXISTS plant_stages CASCADE;
DROP TABLE IF EXISTS nutrients CASCADE;
DROP TABLE IF EXISTS plant_stage_care CASCADE
DROP TABLE IF EXISTS plant_stage_care_nutrients CASCADE
DROP TABLE IF EXISTS plant_life_cycle_stages CASCADE;
DROP TABLE IF EXISTS plant_stage_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycle_start_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycle_end_tasks CASCADE;
DROP TABLE IF EXISTS devices CASCADE;
DROP TABLE IF EXISTS device_actions CASCADE;
DROP TABLE IF EXISTS growing_groups CASCADE;
DROP TABLE IF EXISTS growing_group_devices CASCADE;
DROP TABLE IF EXISTS growing_locations CASCADE;
DROP TABLE IF EXISTS growing_location_devices CASCADE;
DROP TABLE IF EXISTS growing_levels CASCADE;
DROP TABLE IF EXISTS growing_level_devices CASCADE;
DROP TABLE IF EXISTS grow_spots CASCADE;
DROP TABLE IF EXISTS grow_spots_devices CASCADE;