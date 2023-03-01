package model

import "time"

// LoginResponse response model of GetToken request of client.IIncidentClient
type LoginResponse struct {
	Token string `json:"token"`
}

// IncidentReport is response model of GetToken request of client.IIncidentClient
type IncidentReport struct {
	ActionDate      time.Time          `json:"actionDate"`
	IncidentTypes   []ExternalConstant `json:"incidentTypes"`
	VehicleWrappers []VehicleWrapper   `json:"vehicleWrappers"`
}

type ExternalConstant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VehicleWrapper struct {
	Vehicle            Vehicle          `json:"vehicle"`
	TechnicalFailure   ExternalConstant `json:"technicalFailure"`
	DamagePropertyInfo string           `json:"damagePropertyInfo"`
}

type Vehicle struct {
	CertNumber    string           `json:"certNumber"`
	VehicleNumber string           `json:"vehicleNumber"`
	Brand         string           `json:"brand"`
	Type          ExternalConstant `json:"type"`
}

// RegisterResponse is response model of RegisterCar request of client.ISMSRadarClient
type RegisterResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		ID int `json:"id"`
	} `json:"data"`
}

// CarInfoResponse is response model of GetCarInfo request of client.ISMSRadarClient
type CarInfoResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		HasArrest     int    `json:"has_arrest"`
		HasArrestText string `json:"has_arrest_text"`
	} `json:"data"`
}

// ProtocolsInfoResponse is response model of GetProtocols request of client.ISMSRadarClient
type ProtocolsInfoResponse struct {
	Status string     `json:"status"`
	Count  int        `json:"count"`
	Data   []Protocol `json:"data"`
}

type Protocol struct {
	Penalty    float32 `json:"penalty"`
	Total      float32 `json:"total"`
	LawItem    string  `json:"law_item"`
	HasFiles   bool    `json:"has_files"`
	CanPay     bool    `json:"can_pay"`
	DlRequired bool    `json:"dl_required"`
}

// ConvertPDFRequest is response model of ConvertByTemplateNameAndSendToMail request of client.IPDFConverterClient
type ConvertPDFRequest struct {
	Vehicle   LightVehicle    `json:"vehicle"`
	Protocols []LightProtocol `json:"protocols"`
	Incidents []LightIncident `json:"incidents"`
}

type LightVehicle struct {
	Name      string `json:"name"`
	VIN       string `json:"vin"`
	HasArrest int    `json:"has_arrest"`
}

type LightProtocol struct {
	Total       float32 `json:"total"`
	Description string  `json:"description"`
}

type LightIncident struct {
	Date   string   `json:"date"`
	Types  []string `json:"types"`
	Damage string   `json:"damage"`
}

type EMailDTO struct {
	Subject        string `json:"subject" url:"subject"`
	Sender         string `json:"sender" url:"sender"`
	Receivers      string `json:"receivers" url:"receivers"`
	Template       string `json:"template" url:"template"`
	AttachmentName string `json:"attachmentName" url:"attachmentName"`
}

// RecordCheckResponse is response model of CheckRecordsByVIN request of client.ICarfaxClient
type RecordCheckResponse struct {
	Carfax struct {
		Records int `json:"records"`
	} `json:"carfax"`
}

// GetReportResponse is response model of GetReport request of client.ICarfaxClient
type GetReportResponse struct {
	Report struct {
		HTML string `json:"report"`
	} `json:"report"`
}
