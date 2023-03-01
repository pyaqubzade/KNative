package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// IUserActionRepository is interface responsible for function based sql queries for User actions
type IUserActionRepository interface {
	GetUserActionCount(signature string) (int64, error)
	RemoveUserAction(ID uint) error
	SaveUserAction(tx *gorm.DB, actionRecord *ActionRecord) (*ActionRecord, error)
	GetTransaction() (*gorm.DB, error)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
}

const signatureFilter = "signature = ?"

// UserActionRepository implementation of IUserActionRepository interface
type UserActionRepository struct {
	db *gorm.DB
}

func NewUserActionRepository(db *gorm.DB) IUserActionRepository {
	return &UserActionRepository{db: db}
}

func (r *UserActionRepository) GetUserActionCount(signature string) (int64, error) {
	log.Debug("GetUserActionCount.start")
	var count int64
	result := r.db.Model(ActionRecord{}).
		Where(signatureFilter, signature).
		Count(&count)
	if err := result.Error; err != nil {
		log.Errorf("GetUserActionCount.err %s", result.Error.Error())
		return 0, err
	}
	log.Debugf("GetUserActionCount.end affected rows %d", result.RowsAffected)
	return count, nil
}

func (r *UserActionRepository) RemoveUserAction(ID uint) error {
	log.Debug("RemoveUserAction.start")
	result := r.db.Delete(ActionRecord{}, ID)
	if err := result.Error; err != nil {
		log.Errorf("RemoveUserAction.err %s", result.Error.Error())
		return err
	}
	log.Debugf("RemoveUserAction.end affected rows %d", result.RowsAffected)
	return nil
}

func (r *UserActionRepository) SaveUserAction(tx *gorm.DB, actionRecord *ActionRecord) (*ActionRecord, error) {
	log.Debug("SaveUserAction.start")
	result := tx.Model(actionRecord).Save(actionRecord)
	if err := result.Error; err != nil {
		log.Errorf("SaveUserAction.err %s", result.Error.Error())
		return nil, err
	}
	log.Debugf("SaveUserAction.end affected rows %d", result.RowsAffected)
	return actionRecord, nil
}

// GetTransaction opens transaction for further db manipulations
func (r *UserActionRepository) GetTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		log.Errorf("GetTransaction.err %s", tx.Error.Error())
		return nil, err
	}
	return tx, nil
}

// Rollback revert changes made in database in case of errors
func (r *UserActionRepository) Rollback(tx *gorm.DB) {
	result := tx.Rollback()
	if err := result.Error; err != nil {
		log.Error("Failed to rollback current transaction ", err.Error())
	}
}

// Commit is for approve changes made in database is everything is ok
func (r *UserActionRepository) Commit(tx *gorm.DB) {
	result := tx.Commit()
	if err := result.Error; err != nil {
		log.Error("Failed to commit current transaction ", err.Error())
	}
}
