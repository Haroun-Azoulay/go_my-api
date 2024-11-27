package security


import (
	"time"
	"github.com/golang-jwt/jwt/v5"
    "fmt"
)

var secretKey = []byte("secret-key")

func CreateToken(username string, isAdmin bool) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "Firstname": username,
		"IsAdmin": isAdmin,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
       return secretKey, nil
    })
   
    if err != nil {
       return err
    }

    if !token.Valid {
       return fmt.Errorf("invalid token")
    }
   
    return nil
 }

