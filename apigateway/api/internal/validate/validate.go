package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/qq51529210/util"
)

var (
	t ut.Translator
)

func init() {
	util.GinValidateZH(nil)
}
