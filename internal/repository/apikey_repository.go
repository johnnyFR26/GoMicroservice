package repository

import (
	"github.com/johnnyFR26/GoMicroservice/pkg/model"
	"gorm.io/gorm"
)

type APIKeyRepository struct {
	DB *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) *APIKeyRepository {
	return &APIKeyRepository{DB: db}
}

func (r *APIKeyRepository) GetAPIKey(key string) (*model.APIKey, error) {
	var apiKey model.APIKey
	err := r.DB.First(&apiKey, "key = ?", key).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &apiKey, nil
}
