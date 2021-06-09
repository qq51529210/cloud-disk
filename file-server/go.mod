module github.com/qq51529210/cloud-service/file-server

go 1.15

require (
	github.com/qq51529210/cloud-service/util v0.0.0-00010101000000-000000000000
	github.com/qq51529210/http-router v0.0.0-20210604070634-429a8b44612e
	github.com/qq51529210/log v0.0.0-20210529132539-d2d52fbd5103
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/qq51529210/cloud-service/util => ../util
