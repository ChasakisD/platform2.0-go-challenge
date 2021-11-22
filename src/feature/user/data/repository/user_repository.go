package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	db "gwi/assignment/core/data/database"
	http "gwi/assignment/core/data/http"
	ent "gwi/assignment/feature/user/data/entity"
	res "gwi/assignment/feature/user/domain/response"
)

type UserRepository struct{}

var (
	ErrUserInvalidEmail    = errors.New("email address is invalid")
	ErrUserInvalidPassword = errors.New("password should be at least 8 characters")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserJwtFailed       = errors.New("jwt failed")
	ErrUserAuthFailed      = errors.New("invalid password")
	ErrUserInvalidToken    = errors.New("invalid token")
	ErrHasFailed           = errors.New("hash failed")
)

func (repo *UserRepository) CreateUser(user *ent.User) (*res.AuthResponse, error) {
	if err := repo.validateRegister(user); err != nil {
		return nil, err
	}

	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}

	user.Id = uuid.NewString()
	user.Password = string(hashedPasswordBytes)
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	authResponse, err := repo.generateTokens(user.Id)
	if err != nil {
		return nil, ErrUserJwtFailed
	}

	return authResponse, nil
}

func (repo *UserRepository) Authenticate(email, password string) (*res.AuthResponse, error) {
	user, err := repo.validateAuth(email, password)
	if err != nil {
		return nil, err
	}

	authResponse, err := repo.generateTokens(user.Id)
	if err != nil {
		return nil, ErrUserJwtFailed
	}

	return authResponse, nil
}

func (repo *UserRepository) Refresh(accessToken, refreshToken string) (*res.AuthResponse, error) {
	userId, err := repo.validateRefresh(accessToken, refreshToken)
	if err != nil {
		return nil, err
	}

	authResponse, err := repo.generateTokens(userId)
	if err != nil {
		return nil, ErrUserJwtFailed
	}

	return authResponse, nil
}

func (repo *UserRepository) generateTokens(userId string) (*res.AuthResponse, error) {
	accessTokenExp := time.Now().Add(time.Hour)
	_, accessToken, err := http.TokenAuth.Encode(map[string]interface{}{
		"userId": userId,
		"exp":    accessTokenExp.Unix()})
	if err != nil {
		return nil, err
	}

	refreshTokenExp := time.Now().AddDate(0, 6, 0)
	_, refreshToken, err := http.TokenAuth.Encode(map[string]interface{}{
		"userId": userId,
		"exp":    refreshTokenExp.Unix()})
	if err != nil {
		return nil, err
	}

	return &res.AuthResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessTokenExp,
		RefreshTokenExpiresAt: refreshTokenExp,
	}, nil
}

func (repo *UserRepository) getUserByEmail(email string) (*ent.User, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	result := &ent.User{}

	err = db.Where("email = ?", email).First(result).Error
	if err != nil {
		return nil, err
	}

	return result, err
}

func (repo *UserRepository) getUserIdFromToken(accessToken string) (string, error) {
	jwtToken, err := http.TokenAuth.Decode(accessToken)
	if err != nil {
		return "", err
	}

	return jwtToken.PrivateClaims()["userId"].(string), nil
}

func (repo *UserRepository) validateAuth(email, password string) (*ent.User, error) {
	user, err := repo.getUserByEmail(email)
	if user == nil || err != nil || err == gorm.ErrRecordNotFound {
		return nil, ErrUserAuthFailed
	}

	return user, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (repo *UserRepository) validateRefresh(accessToken, refreshToken string) (string, error) {
	userId, err := repo.getUserIdFromToken(accessToken)
	if err != nil {
		return "", ErrUserInvalidToken
	}

	jwtRefreshToken, err := http.TokenAuth.Decode(refreshToken)
	if err != nil {
		return "", ErrUserInvalidToken
	}

	if err = jwt.Validate(jwtRefreshToken); err != nil {
		return "", ErrUserInvalidToken
	}

	return userId, nil
}

func (repo *UserRepository) validateRegister(user *ent.User) error {
	user, err := repo.getUserByEmail(user.Email)
	if user != nil {
		return ErrUserAlreadyExists
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
