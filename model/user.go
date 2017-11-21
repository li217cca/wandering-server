package model

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
)

/*
User user struct
*/
type User struct {
	ID       int    `json:"id"`
	Username string `json:"-"   gorm:"not null;unique"` // 用户名
	Password string `json:"-"   gorm:"not null"`        // 密码
	Salt     string `json:"-"   gorm:"not null"`
}

/*
GetUserByAuthenticate ...
*/
func GetUserByAuthenticate(username string, password string) (User, error) {
	user := User{} // 获取user
	err := DB.Model(&User{}).Where(User{Username: username}).First(&user).Error
	if err != nil {
		return User{}, fmt.Errorf("\nGetUserByAuthenticate 01 \n%v", err)
	}
	if user.Password != hashedPassword(password, globalSalt, user.Salt) {
		return User{}, fmt.Errorf("\nGetUserByAuthenticate 02 \nPassword error")
	}
	return user, nil
}

/*
NewUser ...
Type: pure
UnitTest: false
*/
func NewUser(username string, password string) User {
	salt := strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	user := User{
		Username: username,
		Password: hashedPassword(password, globalSalt, salt),
		Salt:     salt,
	}
	return user
}

func hashedPassword(rawPassword, globalSalt, privateSalt string) string {
	h := sha1.New()
	io.WriteString(h, rawPassword+globalSalt)
	h2 := sha1.New()
	io.WriteString(h2, fmt.Sprintf("%x", h.Sum(nil))+privateSalt)
	return fmt.Sprintf("%x", h2.Sum(nil))
}

/*
NewUserToken ...
Type: pure
UnitTest: false
*/
func NewUserToken(user User, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(user.ID),
		"exp":     exp,
	})
	tokenString, err := token.SignedString([]byte(jwtSigninKey))
	if err != nil {
		return "", fmt.Errorf("\nNewUserToken 01 \n%v", err)
	}
	return tokenString, nil
}

/*
ParseUserToken ...
Type: pure
UnitTest: true
*/
func ParseUserToken(token string) (userID int, err error) {
	var parsedToken *jwt.Token
	if parsedToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSigninKey), nil
	}); err != nil {
		return 0, fmt.Errorf("\nParseUserToken 01 \n%v", err)
	}
	if !parsedToken.Valid {
		return 0, fmt.Errorf("\nParseUserToken 02 \nToken is not valid")
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)
	userID, _ = strconv.Atoi(claims["user_id"].(string))
	return userID, nil
}
