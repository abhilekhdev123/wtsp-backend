package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("snapeat")

const accessExp = 60 * time.Minute
const refreshExp = 7 * 24 * time.Hour

func GenerateJWT(userID string, role []string) (string, string, error) {

	data := map[string]interface{}{
		"userID": userID,
		"role":   role,
	}
	// Access Token
	accessClaims := jwt.MapClaims{
		"data":   data,
		"exp":    time.Now().Add(accessExp).Unix(),
		"secret": jwtSecret,
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := jwt.MapClaims{
		"data": data
		"exp":    time.Now().Add(refreshExp).Unix(),
		"secret": jwtSecret
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	/////after generating the tokens, you can store them in the database

	return accessToken, refreshToken, nil
}
