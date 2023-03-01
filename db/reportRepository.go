package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// IReportRepository is interface responsible for function based sql queries
type IReportRepository interface {
	SaveLocalReport(tx *gorm.DB, report *LocalReport) (*LocalReport, error)
	SaveExternalReport(tx *gorm.DB, report *ExternalReport) (*ExternalReport, error)
	SaveReports(tx *gorm.DB, reports *[]*Report) ([]*Report, error)
	GetTransaction() (*gorm.DB, error)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
}

// ReportRepository implementation of IReportRepository interface
type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) IReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) SaveLocalReport(tx *gorm.DB, report *LocalReport) (*LocalReport, error) {
	log.Debug("SaveLocalReport.start")
	result := tx.Model(report).Save(report)
	if err := result.Error; err != nil {
		log.Errorf("SaveLocalReport.err %s", result.Error.Error())
		return nil, err
	}
	log.Debugf("SaveLocalReport.end affected rows %d", result.RowsAffected)
	return report, nil
}

func (r *ReportRepository) SaveExternalReport(tx *gorm.DB, report *ExternalReport) (*ExternalReport, error) {
	log.Debug("SaveExternalReport.start")
	result := tx.Model(report).Save(report)
	if err := result.Error; err != nil {
		log.Errorf("SaveExternalReport.err %s", result.Error.Error())
		return nil, err
	}
	log.Debugf("SaveExternalReport.end affected rows %d", result.RowsAffected)
	return report, nil
}

func (r *ReportRepository) SaveReports(tx *gorm.DB, reports *[]*Report) ([]*Report, error) {
	log.Debug("SaveReports.start")
	result := tx.Model(reports).Save(reports)
	if err := result.Error; err != nil {
		log.Errorf("SaveReports.err %s", result.Error.Error())
		return nil, err
	}
	log.Debugf("SaveReports.end affected rows %d", result.RowsAffected)
	return *reports, nil
}

// GetTransaction opens transaction for further db manipulations
func (r *ReportRepository) GetTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		log.Errorf("GetTransaction.err %s", tx.Error.Error())
		return nil, err
	}
	return tx, nil
}

// Rollback revert changes made in database in case of errors
func (r *ReportRepository) Rollback(tx *gorm.DB) {
	result := tx.Rollback()
	if err := result.Error; err != nil {
		log.Error("Failed to rollback current transaction ", err.Error())
	}
}

// Commit is for approve changes made in database is everything is ok
func (r *ReportRepository) Commit(tx *gorm.DB) {
	result := tx.Commit()
	if err := result.Error; err != nil {
		log.Error("Failed to commit current transaction ", err.Error())
	}
}
