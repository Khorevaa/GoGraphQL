package authHandler

import (
	"github.com/NiciiA/GoGraphQL/domain/model/accountModel"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

func CreateJWT(account accountModel.Account) string {
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