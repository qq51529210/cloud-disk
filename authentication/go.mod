module github.com/qq51529210/micro-service/authentication

go 1.18

require (
	github.com/qq51529210/log v0.0.0
	github.com/qq51529210/web v0.0.0
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace (
	github.com/qq51529210/log => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/log
	github.com/qq51529210/web => /Users/linwenbin/Develop/project/go/src/github.com/qq51529210/web
)
