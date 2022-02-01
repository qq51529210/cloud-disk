module github.com/qq51529210/micro-services/auth

go 1.15

require (
	github.com/qq51529210/log v0.0.0
	github.com/qq51529210/redis v0.0.0
	github.com/qq51529210/uuid v0.0.0
	github.com/qq51529210/web v0.0.0
	go.mongodb.org/mongo-driver v1.8.2
)

replace (
	github.com/qq51529210/log => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/log
	github.com/qq51529210/redis => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/redis
	github.com/qq51529210/uuid => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/uuid
	github.com/qq51529210/web => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/web
	github.com/qq51529210/web/router => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/web/router
)
