package handler

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"
)

type AllClinicResponse struct {
	ID         uint                `json:"admin_id"`
	ClinicName string              `json:"clinic_name"`
	Open       availdaydoctor.Core `json:"open"`
	Service    string              `json:"service"`
	Veterinary string              `json:"veterinary"`
	Location   string              `json:"location"`
}

func ResponseAllClinic(admin admin.Core, doctor doctor.Core) AllClinicResponse {
	result := AllClinicResponse{
		ID:         admin.ID,
		ClinicName: admin.FullName,
		Open:       availdaydoctor.Core{},
		Veterinary: doctor.FullName,
		Location:   admin.Address,
	}
	return result
}
