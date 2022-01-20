package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/web/router"
)

var (
	errQueryData    []byte
	errParseJSON    []byte
	errUnauthorized []byte
)

func init() {
	errQueryData, _ = json.Marshal(map[string]string{
		"error": "query data error",
	})
	errParseJSON, _ = json.Marshal(map[string]string{
		"error": "parse JSON body error",
	})
	errUnauthorized, _ = json.Marshal(map[string]string{
		"error": "unauthorized",
	})
}

func QueryDataError(ctx *router.Context, err error) {
	log.DepthError(1, err)
	ctx.JSON(http.StatusInternalServerError, errQueryData)
}

func ParseJSONError(ctx *router.Context, err error) {
	log.DepthError(1, err)
	ctx.JSON(http.StatusBadRequest, errParseJSON)
}

func FormValueError(ctx *router.Context, query, value string) {
	ctx.JSON(http.StatusBadRequest, map[string]string{
		"error": fmt.Sprintf(`invalid value "%s" of query "%s"`, value, query),
	})
}

func UnauthorizedError(ctx *router.Context) {
	ctx.JSON(http.StatusUnauthorized, errUnauthorized)
}
