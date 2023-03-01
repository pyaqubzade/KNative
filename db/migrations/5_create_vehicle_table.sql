-- +migrate Up
CREATE TABLE vehicle
(
    id             bigserial NOT NULL
        CONSTRAINT pk_vehicle PRIMARY KEY,
    report_id      bigint
        CONSTRAINT fk_local_reports_vehicle
            REFERENCES local_report,
    cert_number    text,
    vehicle_number text,
    brand          text,
    "type"         text,
    created_at     timestamp DEFAULT now(),
    updated_at     timestamp DEFAULT now()
);

COMMENT ON TABLE vehicle IS 'Table for storing data about DYP protocols';

COMMENT ON COLUMN vehicle.id IS 'Primary key for vehicle record';

COMMENT ON COLUMN vehicle.report_id IS 'ID number of report record this record related to';

COMMENT ON COLUMN vehicle.cert_number IS 'ID number of driver licence';

COMMENT ON COLUMN vehicle.vehicle_number IS 'Vehicle registration sign number of car';

COMMENT ON COLUMN vehicle.brand IS 'Brand and model name of car';

COMMENT ON COLUMN vehicle.type IS 'Type of vehicle';

COMMENT ON COLUMN vehicle.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN vehicle.updated_at IS 'Update timestamp of record';