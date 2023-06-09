package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface{
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmail(input CheckEmailAvailable) (bool, error)
	SaveAvatar(ID int, filelocation string) (User, error)
	GetUserByID(ID int) (User, error)
}
// struct memakai devendenci dari repository dengan implement instace lewat fun NewService
type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error){
	
	user := User{}
	user.Name 			= input.Name
	user.Occupation 	= input.Occupation
	user.Email			= input.Email
	passwordHash, err 	:= bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil{
		return user, err
	}

	user.PasswordHash 	= string(passwordHash)
	user.Role			= "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return	user, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return	user, err
	}
	if user.ID == 0 {
		return user, errors.New("Email Tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmail(input CheckEmailAvailable) (bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return	true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(ID int, filelocation string) (User, error){
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation
	
	avatarUpload, err := s.repository.Update(user)
	if err != nil {
		return avatarUpload, err
	}
	return avatarUpload, nil
}

func (s *service) GetUserByID(ID int) (User, error){
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("tidak menemukan user dengan id ini")
	}
	return user, nil
}