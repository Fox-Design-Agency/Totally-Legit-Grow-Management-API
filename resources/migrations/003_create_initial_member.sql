-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO organizations
    (display_name)
    VALUES ('SYSYEM ORG');

INSERT INTO members
    (display_name)
    VALUES ('SYSTEM USER');

INSERT INTO organization_members
    (organization, member)
    VALUES (1, 1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
