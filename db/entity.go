package db

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/pyaqubzade/knative/model"
)

type BaseEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Report struct {
	BaseEntity
	VIN         string
	Email       string
	PhoneNumber string
}

type ExternalReport struct {
	Report
	HTML string
}

func (ExternalReport) TableName() string {
	return "external_report"
}

type LocalReport struct {
	Report
	Incidents     []Incident `gorm:"foreignKey:ReportID"`
	Vehicle       Vehicle    `gorm:"foreignKey:ReportID"`
	Protocols     []Protocol `gorm:"foreignKey:ReportID"`
	CarID         int
	HasArrest     int
	HasArrestText string
}

func (LocalReport) TableName() string {
	return "local_report"
}

type Incident struct {
	BaseEntity
	ReportID           uint
	ActionDate         time.Time
	IncidentTypes      pgtype.JSONB `gorm:"type:jsonb"`
	TechnicalFailure   string
	DamagePropertyInfo string
}

func (Incident) TableName() string {
	return "incident"
}

type Vehicle struct {
	BaseEntity
	ReportID      uint
	CertNumber    string
	VehicleNumber string
	Brand         string
	Type          string
}

func (Vehicle) TableName() string {
	return "vehicle"
}

type Protocol struct {
	BaseEntity
	ReportID uint
	model.Protocol
}

func (Protocol) TableName() string {
	return "protocol"
}

type ActionRecord struct {
	BaseEntity
	model.ReportRequest
	Signature string
}

func (ActionRecord) TableName() string {
	return "action_record"
}
