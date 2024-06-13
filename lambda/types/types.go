package types

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type Device struct {
	DeviceID   string `json:"deviceid"`
	DateJoined string `json:"datejoined"`
	SessionID  string `json:"sessionid"`
}

type Session struct {
	SessionID string `json:"sessionid"`
	Device1ID string `json:"device1id"`
	Device2ID string `json:"device2id"`
}

type Heart struct {
	HeartID          string `json:"heartid"`
	SenderDeviceID   string `json:"senderdeviceid"`
	ReceiverDeviceID string `json:"receiverdeviceid"`
	Timestamp        string `json:"timestamp"`
    Seen             bool   `json:"seen"`
    SeenTimestamp    int64  `json:"seentimestamp"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil

}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "secret"

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
		// fmt.Errorf("signingString error %w", err)
		// return tokenString
	}
	return tokenString
}
