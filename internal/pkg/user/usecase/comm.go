package usecase

import (
	"codex/internal/pkg/domain"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func InitUsrUc(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (uc userUsecase) GetBasicInfo(id uint64) (domain.User, error) {
	us, err := uc.userRepo.GetById(id)
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	return us.ClearPasswords(), nil
}

func (uc userUsecase) Register(us domain.User) (domain.User, error) {
	trimCredentials(&us.Email, &us.Username, &us.Password, &us.RepeatPassword)

	if us.Email == "" || us.Username == "" || us.Password == "" || us.RepeatPassword == "" {
		return domain.User{}, domain.Err.ErrObj.EmptyField
	}

	if err := validateEmail(us.Email); err != nil {
		return domain.User{}, err
	}

	if err := validateUsername(us.Username); err != nil {
		return domain.User{}, err
	}

	if err := validatePassword(us.Password); err != nil {
		return domain.User{}, err
	}

	if us.Password != us.RepeatPassword {
		return domain.User{}, domain.Err.ErrObj.UnmatchedPasswords
	}

	if _, err := uc.userRepo.GetByEmail(us.Email); err == nil {
		return domain.User{}, err
	}

	_, err := uc.userRepo.AddUser(us)
	if err != nil {
		return domain.User{}, err
	}

	return us.ClearPasswords(), nil
}

func (uc userUsecase) Login(ub domain.UserBasic) (domain.User, error) {
	if ub.Email == "" || ub.Password == "" {
		return domain.User{}, domain.Err.ErrObj.EmptyField
	}

	us, err := uc.userRepo.GetByEmail(ub.Email)
	if err != nil {
		return domain.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(ub.Password)); err != nil {
		return domain.User{}, domain.Err.ErrObj.BadPassword
	}

	return us.ClearPasswords(), nil
}

func (uc userUsecase) CheckAuth(id uint64) (domain.User, error) {
	us := domain.User{Id: id}
	return us, nil
}