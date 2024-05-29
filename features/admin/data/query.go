package data

import (
	"PetPalApp/features/admin"

	"gorm.io/gorm"
)

type AdminModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) admin.AdminModel {
	return &AdminModel{
		db: db,
	}
}

func (am *AdminModel) Register(admin admin.Core) error {
	adminGorm := Admin{
		FullName: admin.FullName,
		Email: admin.Email,
		NumberPhone: admin.NumberPhone,
		Address: admin.Address,
		Password: admin.Password,
		ProfilePicture: admin.ProfilePicture,
	}
	tx := am.db.Create(&adminGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}