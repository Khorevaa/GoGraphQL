package authHandler

import (
	"net/http"
	"encoding/json"
	"github.com/NiciiA/GoGraphQL/domain/model/accountModel"
)

func Post(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(accountModel.Account{UserName: "username"})
}