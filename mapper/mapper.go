package mapper

import (
	"github.com/pyaqubzade/knative/db"
	"github.com/pyaqubzade/knative/model"
	"github.com/sirupsen/logrus"
)

func MapToLocalReport(request *model.ReportRequest, reports []model.IncidentReport, register *model.RegisterResponse,
	carInfo *model.CarInfoResponse, protocolsResponse []model.Protocol) *db.LocalReport {
	vehicle := reports[0].VehicleWrappers[0].Vehicle
	var incidents []db.Incident
	for _, report := range reports {
		var incidentTypes []string
		for _, incidentType := range report.IncidentTypes {
			incidentTypes = append(incidentTypes, incidentType.Name)
		}
		wrapper := report.VehicleWrappers[0]
		incident := db.Incident{
			ActionDate:         report.ActionDate,
			TechnicalFailure:   wrapper.TechnicalFailure.Name,
			DamagePropertyInfo: wrapper.DamagePropertyInfo,
		}
		_ = incident.IncidentTypes.Set(incidentTypes)
		incidents = append(incidents, incident)
	}
	var protocols []db.Protocol
	for _, protocol := range protocolsResponse {
		protocols = append(protocols, db.Protocol{Protocol: protocol})
	}

	return &db.LocalReport{
		Report: db.Report{
			VIN:         request.VIN,
			Email:       request.Email,
			PhoneNumber: request.PhoneNumber,
		},
		Incidents: incidents,
		CarID:     register.Data.ID,
		Vehicle: db.Vehicle{
			CertNumber:    vehicle.CertNumber,
			VehicleNumber: vehicle.VehicleNumber,
			Brand:         vehicle.Brand,
			Type:          vehicle.Type.Name,
		},
		HasArrest:     carInfo.Data.HasArrest,
		HasArrestText: carInfo.Data.HasArrestText,
		Protocols:     protocols,
	}
}

func MapToExternalReport(request *model.ReportRequest, html []byte) *db.ExternalReport {
	return &db.ExternalReport{
		Report: db.Report{
			VIN:         request.VIN,
			Email:       request.Email,
			PhoneNumber: request.PhoneNumber,
		},
		HTML: string(html),
	}
}

func MapToConvertPDFRequest(report *db.LocalReport) model.ConvertPDFRequest {
	var protocols []model.LightProtocol
	for _, protocol := range report.Protocols {
		protocols = append(protocols, model.LightProtocol{
			Total:       protocol.Total,
			Description: protocol.LawItem,
		})
	}
	var incidents []model.LightIncident
	for _, incident := range report.Incidents {
		var types []string
		err := incident.IncidentTypes.AssignTo(&types)
		if err != nil {
			logrus.Fatal("Unable to map JSONB to String array")
		}
		incidents = append(incidents, model.LightIncident{
			Types:  types,
			Damage: incident.DamagePropertyInfo,
			Date:   incident.ActionDate.Format("02.01.2006"),
		})
	}

	return model.ConvertPDFRequest{
		Vehicle: model.LightVehicle{
			Name:      report.Vehicle.Brand,
			VIN:       report.VIN,
			HasArrest: report.HasArrest,
		},
		Protocols: protocols,
		Incidents: incidents,
	}
}
