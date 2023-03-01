-- +migrate Up
CREATE TABLE incident
(
    id                   bigserial NOT NULL
        CONSTRAINT pk_incident PRIMARY KEY,
    report_id            bigint
        CONSTRAINT fk_local_report_incident
            REFERENCES local_report,
    action_date          timestamp,
    incident_types       jsonb,
    technical_failure    text,
    damage_property_info text,
    created_at           timestamp DEFAULT now(),
    updated_at           timestamp DEFAULT now()
);


COMMENT ON TABLE incident IS 'Table for storing data about car incidents';

COMMENT ON COLUMN incident.id IS 'Primary key for incident record';

COMMENT ON COLUMN incident.report_id IS 'ID number of report record this record related to';

COMMENT ON COLUMN incident.action_date IS 'Timestamp of incident';

COMMENT ON COLUMN incident.incident_types IS 'List of incident types';

COMMENT ON COLUMN incident.technical_failure IS 'Technical failure description of car';

COMMENT ON COLUMN incident.damage_property_info IS 'Damage degree of car this incident caused';

COMMENT ON COLUMN incident.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN incident.updated_at IS 'Update timestamp of record';