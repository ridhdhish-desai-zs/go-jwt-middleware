package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

/*
URL: /api/users
Method: POST
route: Unprotected
Description: Create new jwt token for static email and password
*/
func CreateToken(res http.ResponseWriter, req *http.Request) {
	claims := jwt.MapClaims{}
	claims["email"] = "ridhdhish@gmail.com"
	claims["password"] = "123456"

	cl := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := cl.SignedString([]byte("mysecret"))

	if err != nil {
		_, _ = res.Write([]byte(`{"error": "Cannot able to generate JWT"}`))
	}

	_, _ = res.Write([]byte(fmt.Sprintf(`{"token": %v}`, token)))
}

/*
URL: /api/users
Method: GET
route: Unprotected
Description: Fetch all users
*/
func ValidateUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	data := ctx.Value("data")
	fmt.Printf("%T, %v", data, data)

	jsonData, _ := json.Marshal(data)

	_, _ = res.Write([]byte(fmt.Sprintf(`"data": {"user": {%v}}`, string(jsonData))))
}
