-- +migrate Up
CREATE TABLE action_record
(
    id           bigserial NOT NULL
        CONSTRAINT pk_action_record PRIMARY KEY,
    signature    varchar,
    vin          varchar,
    email        varchar,
    phone_number varchar,
    created_at   timestamp DEFAULT now(),
    updated_at   timestamp DEFAULT now()
);

COMMENT ON TABLE action_record IS 'Table for storing data user requests by signature which stored in browser cookie';

COMMENT ON COLUMN action_record.id IS 'Primary key for report record';

COMMENT ON COLUMN action_record.signature IS 'NanoID created by FMS service';

COMMENT ON COLUMN action_record.vin IS 'Vehicle identification number of car';

COMMENT ON COLUMN action_record.email IS 'Email address of report requester';

COMMENT ON COLUMN action_record.phone_number IS 'Phone number of report requester';

COMMENT ON COLUMN action_record.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN action_record.updated_at IS 'Update timestamp of record';