-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE EXTENSION IF NOT EXISTS pgcrypto;

SET TIME ZONE 'UTC';

-- organizations
CREATE TABLE organizations(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- users
CREATE TABLE users(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- organization_users
CREATE TABLE organization_users(
    organization UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    PRIMARY KEY (organization, user)
);

-- plant categories
CREATE TABLE plant_categories(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plants
CREATE TABLE plants(
    id UUID PRIMARY KEY,
    plant_category UUID NOT NULL REFERENCES plant_categories(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);


-- model_tasks
CREATE TABLE model_tasks(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- generated_tasks
CREATE TABLE generated_tasks(
    id UUID PRIMARY KEY,
    task UUID NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plant life cycles
CREATE TABLE plant_life_cycles(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plant stages
CREATE TABLE plant_stages(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- plant life cycle stages
CREATE TABLE plant_life_cycle_stages(
    plant_life_cycle UUID NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    plant_stage UUID NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    PRIMARY KEY (plant_life_cycle, plant_stage)
);

-- plant stage tasks
CREATE TABLE plant_stage_tasks(
    model_task UUID NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_stage UUID NOT NULL REFERENCES plant_stages(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_stage)
);

-- plant life cycle start tasks
CREATE TABLE plant_life_cycle_start_tasks(
    model_task UUID NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_life_cycle UUID NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_life_cycle)
);

-- plant life cycle end tasks
CREATE TABLE plant_life_cycle_end_tasks(
    model_task UUID NOT NULL REFERENCES model_tasks(id) ON DELETE CASCADE,
    plant_life_cycle UUID NOT NULL REFERENCES plant_life_cycles(id) ON DELETE CASCADE,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    PRIMARY KEY (model_task, plant_life_cycle)
);

-- devices (lighting, water, sensors, etc)
CREATE TABLE devices(
    id UUID PRIMARY KEY,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing groups
CREATE TABLE growing_groups(
    id UUID PRIMARY KEY,
    organization UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing group devices
CREATE TABLE growing_group_devices(
    id UUID PRIMARY KEY,
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_group UUID NOT NULL REFERENCES growing_groups(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing locations
CREATE TABLE growing_locations(
    id UUID PRIMARY KEY,
    growing_group UUID NOT NULL REFERENCES growing_groups(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing location devices
CREATE TABLE growing_location_devices(
    id UUID PRIMARY KEY,
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_location UUID NOT NULL REFERENCES growing_locations(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing levels
CREATE TABLE growing_levels(
    id UUID PRIMARY KEY,
    growing_location UUID NOT NULL REFERENCES growing_locations(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- growing level devices
CREATE TABLE growing_level_devices(
    id UUID PRIMARY KEY,
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    growing_level UUID NOT NULL REFERENCES growing_levels(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- grow spots
CREATE TABLE grow_spots(
    id UUID PRIMARY KEY,
    grow_spot UUID NOT NULL REFERENCES grow_spots(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);

-- grow spot devices
CREATE TABLE grow_spots_devices(
    id UUID PRIMARY KEY,
    device UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    grow_spot UUID NOT NULL REFERENCES grow_spots(id) ON DELETE CASCADE,
    display_name VARCHAR(256) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta JSON DEFAULT '{}',
    archived TIMESTAMP
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS organizations CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS organization_users CASCADE;
DROP TABLE IF EXISTS plant_categories CASCADE;
DROP TABLE IF EXISTS plants CASCADE;
DROP TABLE IF EXISTS model_tasks CASCADE;
DROP TABLE IF EXISTS generated_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycles CASCADE;
DROP TABLE IF EXISTS plant_stages CASCADE;
DROP TABLE IF EXISTS plant_life_cycle_stages CASCADE;
DROP TABLE IF EXISTS plant_stage_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycle_start_tasks CASCADE;
DROP TABLE IF EXISTS plant_life_cycle_end_tasks CASCADE;
DROP TABLE IF EXISTS devices CASCADE;
DROP TABLE IF EXISTS growing_groups CASCADE;
DROP TABLE IF EXISTS growing_group_devices CASCADE;
DROP TABLE IF EXISTS growing_locations CASCADE;
DROP TABLE IF EXISTS growing_location_devices CASCADE;
DROP TABLE IF EXISTS growing_levels CASCADE;
DROP TABLE IF EXISTS growing_level_devices CASCADE;
DROP TABLE IF EXISTS grow_spots CASCADE;
DROP TABLE IF EXISTS grow_spots_devices CASCADE;