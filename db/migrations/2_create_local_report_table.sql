-- +migrate Up
CREATE TABLE local_report
(
    id              bigserial NOT NULL
        CONSTRAINT pk_local_report PRIMARY KEY,
    vin             varchar,
    email           varchar,
    phone_number    varchar,
    car_id          bigint,
    has_arrest      int,
    has_arrest_text text,
    created_at      timestamp DEFAULT now(),
    updated_at      timestamp DEFAULT now()
);

COMMENT ON TABLE local_report IS 'Table for storing data about AvtoYoxla local reports from SMSRadar Insurance API';

COMMENT ON COLUMN local_report.id IS 'Primary key for report record';

COMMENT ON COLUMN local_report.vin IS 'Vehicle identification number of car';

COMMENT ON COLUMN local_report.email IS 'Email address of report requester';

COMMENT ON COLUMN local_report.phone_number IS 'Phone number of report requester';

COMMENT ON COLUMN local_report.car_id IS 'ID number of car in SMS Radar application';

COMMENT ON COLUMN local_report.has_arrest IS 'Code number of arrest text';

COMMENT ON COLUMN local_report.has_arrest_text IS 'Description of arrest information about car from SMS Radar';

COMMENT ON COLUMN local_report.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN local_report.updated_at IS 'Update timestamp of record';