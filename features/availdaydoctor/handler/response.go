package handler

import "PetPalApp/features/availdaydoctor"

type AvailableDayResponse struct {
	DoctorID  uint `json:"doctor_id"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
}

func GormToCore(gorm availdaydoctor.Core) AvailableDayResponse {
	return AvailableDayResponse{
		DoctorID:  gorm.DoctorID,
		Monday:    gorm.Monday,
		Tuesday:   gorm.Tuesday,
		Wednesday: gorm.Wednesday,
		Thursday:  gorm.Thursday,
		Friday:    gorm.Friday,
	}
}
