package usecase

import (
	"xenforo/app/internal/config"
	"xenforo/app/internal/domain/user/model"
	"xenforo/app/pkg/api/jwt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateTokens(thisUser *model.UserAndTokens, cfg *config.Config) error {
	accessToken, err := jwt.CreateToken(
		cfg.App.Jwt.AccessTokenExpiredIn, thisUser.ID,
		cfg.App.Jwt.AccessTokenPrivateKey,
	)
	if err != nil {
		return err
	}

	refreshToken, err := jwt.CreateToken(
		cfg.App.Jwt.RefreshTokenExpiredIn, thisUser.ID,
		cfg.App.Jwt.RefreshTokenPrivateKey,
	)
	if err != nil {
		return err
	}

	thisUser.Tokens.Access = accessToken
	thisUser.Tokens.Refresh = refreshToken

	return nil
}
