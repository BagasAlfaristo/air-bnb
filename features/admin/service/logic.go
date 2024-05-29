package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/utils/encrypts"
	"errors"
)

type AdminService struct {
	AdminModel admin.AdminModel
	hashService encrypts.HashInterface
}

func New(am admin.AdminModel, hash encrypts.HashInterface) admin.AdminService {
	return &AdminService{
		AdminModel: am,
		hashService: hash,
	}
}

func (as *AdminService) Register(admin admin.Core) error {
	if admin.FullName == "" || admin.Email == "" || admin.NumberPhone == "" ||  admin.Address == "" || admin.Password == "" || 
	admin.KetikUlangPassword == "" {
		return errors.New("[validation] Fullname/email/numberphone/address/password tidak boleh kosong")
	}

	if admin.Password != admin.KetikUlangPassword {
		return errors.New("[validation] password dan konfirmasi password tidak cocok")
	}

	var errHash error
	if admin.Password, errHash = as.hashService.HashPassword(admin.Password); errHash != nil {
		return errHash
	}

	if admin.KetikUlangPassword, errHash = as.hashService.HashPassword(admin.KetikUlangPassword); errHash != nil {
		return errHash
	}

	return as.AdminModel.Register(admin)
}

func (as *AdminService) Login(email string, password string) (data *admin.Core, token string, err error) {
	data, err = as.AdminModel.AdminByEmail(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := as.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("email atau password tidak sesuai")
	}

	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}