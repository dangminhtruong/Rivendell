package helpers 

import(
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username, password string) string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})

	tokenString, err := token.SignedString([]byte("is_admin"))
	if err != nil {
		panic(err.Error())
	}
	return tokenString
} 