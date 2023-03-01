-- +migrate Up
CREATE TABLE protocol
(
    id          bigserial NOT NULL
        CONSTRAINT pk_protocol PRIMARY KEY,
    report_id   bigint
        CONSTRAINT fk_local_report_protocol
            REFERENCES local_report,
    penalty     int,
    total       int,
    law_item    text,
    has_files   boolean,
    can_pay     boolean,
    dl_required boolean,
    created_at  timestamp DEFAULT now(),
    updated_at  timestamp DEFAULT now()
);


COMMENT ON TABLE protocol IS 'Table for storing data about DYP protocols';

COMMENT ON COLUMN protocol.id IS 'Primary key for protocol record';

COMMENT ON COLUMN protocol.report_id IS 'ID number of report record this record related to';

COMMENT ON COLUMN protocol.penalty IS 'Penalty point of protocol';

COMMENT ON COLUMN protocol.total IS 'Total amount of fee';

COMMENT ON COLUMN protocol.law_item IS 'The Law that driver broke';

COMMENT ON COLUMN protocol.has_files IS 'Does protocol has files like pictures or videos';

COMMENT ON COLUMN protocol.can_pay IS 'Does fee ready to be payed';

COMMENT ON COLUMN protocol.created_at IS 'Creation timestamp of record';

COMMENT ON COLUMN protocol.updated_at IS 'Update timestamp of record';