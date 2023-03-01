-- +migrate Up
CREATE TABLE external_report
(
    id           bigserial NOT NULL
        CONSTRAINT pk_external_report PRIMARY KEY,
    vin          varchar,
    email        varchar,
    phone_number varchar,
    html         text,
    created_at   timestamp DEFAULT now(),
    updated_at   timestamp DEFAULT now()
);

COMMENT ON TABLE external_report IS 'Table for storing data about AvtoYoxla external reports from CarFax';

COMMENT ON COLUMN external_report.id IS 'Primary key for report record';

COMMENT ON COLUMN external_report.vin IS 'Vehicle identification number of car';

COMMENT ON COLUMN external_report.email IS 'Email address of report requester';

COMMENT ON COLUMN external_report.phone_number IS 'Phone number of report requester';

COMMENT ON COLUMN external_report.html IS 'HTML data retrieved from CarFax';

COMMENT ON COLUMN external_report.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN external_report.updated_at IS 'Update timestamp of record';