package service

import (
	"database/sql"
	"invoice-go/model"
	"invoice-go/repository"
)

func LoadProfileByEmail(db *sql.DB, email string) (*model.AppProfile, error) {
	return repository.GetUserEmail(db, email)
}

func UpdateProfile(db *sql.DB, profile model.AppProfile) error {
	return repository.SaveUserProfile(db, profile)
}

func IsUserProfileExist(db *sql.DB, email string) bool {
	_, err := repository.GetUserEmail(db, email)
	return err == nil
}
