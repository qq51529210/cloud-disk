package users

import (
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	r.POST("/users", post)
	r.DELETE("/users", delete)
	r.PATCH("/users", patch)
	r.GET("/users", list)
}
