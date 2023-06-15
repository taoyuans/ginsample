package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS512
	JwtSecret        = "j90uj3otjfglirdslfjg3o49wjdfsligs23rfsd"
)

const (
	JWT_KEY = "JWT_SECRET"
)

//	func init() {
//		if s := os.Getenv(JWT_KEY); s != "" {
//			JwtSecret = s
//		}
//	}
func NewToken(m map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"iss": "ginsample",
		"aud": "ginsample",
		"nbf": time.Now().Add(-time.Minute * 5).Unix(),
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwtSigningMethod, claims).SignedString([]byte(JwtSecret))
}

func Renew(token string) (string, error) {
	claim, err := Extract(token)
	if err != nil {
		return "", err
	}
	claim["nbf"] = time.Now().Unix()
	claim["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	return jwt.NewWithClaims(jwtSigningMethod, claim).SignedString([]byte(JwtSecret))
}

func EditPayload(token string, m map[string]interface{}) (string, error) {
	claimInfo, err := Extract(token)
	if err != nil {
		return "", err
	}

	for k, v := range m {
		claimInfo[k] = v
	}

	return jwt.NewWithClaims(jwtSigningMethod, claimInfo).SignedString([]byte(JwtSecret))
}

type ErrContainer struct {
	ErrCode    int
	ErrMessage string
}

func (e *ErrContainer) Error() string {
	return fmt.Sprintf("code:%d - msg:%s", e.ErrCode, e.ErrMessage)
}

func Extract(token string) (jwt.MapClaims, error) {
	return ExtractWithSecret(token, JwtSecret)
}
func ExtractWithSecret(token, jwtSecret string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: "Required authorization token not found."}
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(JwtSecret), nil })
	if err != nil {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: fmt.Sprintf("Error parsing token: %v", err)}
	}

	if jwtSigningMethod != nil && jwtSigningMethod.Alg() != parsedToken.Header["alg"] {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: fmt.Sprintf("Expected %s signing method but token specified %s",
			jwtSigningMethod.Alg(),
			parsedToken.Header["alg"])}
	}

	if !parsedToken.Valid {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: "Token is invalid"}
	}

	claimInfo, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: "invalid token"}
	}
	return claimInfo, nil
}

func GetTokenInfo(ctx *gin.Context) (jwt.MapClaims, error) {
	var token string
	if authHeader := ctx.Request.Header.Get("Authorization"); strings.HasPrefix(authHeader, "Bearer ") {
		token = authHeader[7:]
	} else {
		return nil, nil
	}
	userMapClaims, err := Extract(token)
	if err != nil {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: "invalid token(Extract):There is no valid information.(" + token + "):" + err.Error()}
	} else if userMapClaims == nil {
		return nil, &ErrContainer{ErrCode: http.StatusUnauthorized, ErrMessage: "invalid token(Extract.userMapClaims):There is no valid information.(" + token + ")"}
	}
	return userMapClaims, nil
}
