package api

import (
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/jwt"
	"github.com/qq51529210/log"
)

// Parse form before handle, return false there is a error.
func ParseForm(c *router.Context) bool {
	err := c.Req.ParseForm()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func genHSJWT(alg, key string) (string, error) {
	header := make(map[string]interface{})
	payload := make(map[string]interface{})
	switch alg {
	case string(jwt.HS384):
		return jwt.GenerateHS384(header, payload, []byte(key))
	case string(jwt.HS512):
		return jwt.GenerateHS512(header, payload, []byte(key))
	default:
		return jwt.GenerateHS256(header, payload, []byte(key))
	}
}
