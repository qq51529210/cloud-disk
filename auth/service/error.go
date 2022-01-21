package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/web/router"
)

const (
	_ErrCodeParseJSON = iota
	_ErrCodeFormValue
	_ErrCodeQueryData
	_ErrCodeUnauthorized
)

var (
	errQueryData    []byte
	errParseJSON    []byte
	errUnauthorized []byte
)

func init() {
	errQueryData, _ = json.Marshal(map[string]interface{}{
		"error": "query data error",
		"code":  _ErrCodeQueryData,
	})
	errParseJSON, _ = json.Marshal(map[string]interface{}{
		"error": "parse JSON body error",
		"code":  _ErrCodeParseJSON,
	})
	errUnauthorized, _ = json.Marshal(map[string]interface{}{
		"error": "unauthorized",
		"code":  _ErrCodeUnauthorized,
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

func UnauthorizedError(ctx *router.Context) {
	ctx.JSON(http.StatusUnauthorized, errUnauthorized)
}

func FormValueError(ctx *router.Context, query, value string) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"error": fmt.Sprintf(`invalid value "%s" of query "%s"`, value, query),
		"code":  _ErrCodeFormValue,
		"query": query,
		"value": value,
	})
}
