package authHandler

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var CurrentAccount model.Account

func CreateJWT(account model.Account) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountId": account.ID,
		"groups": account.Groups,
		"roles": account.Roles,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().AddDate(4, 0, 0).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secretKeyM8"))
	return tokenString
}