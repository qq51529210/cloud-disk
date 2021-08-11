module github.com/qq51529210/micro-services/auth

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/qq51529210/http-router v0.0.0-20210706182707-2fd7a41472ad
	github.com/qq51529210/jwt v0.0.0-20210531120038-6eb2212a7688 // indirect
	github.com/qq51529210/log v0.0.0-20210707174109-e593bd4a8cd7
	github.com/qq51529210/micro-services/util v0.0.0-00010101000000-000000000000
	github.com/qq51529210/redis v0.0.0-20210610092729-dbe50e9924d8
	github.com/qq51529210/uuid v0.0.0-20210410083004-ce2b0df9936f
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/qq51529210/micro-services/auth/api => ./api

replace github.com/qq51529210/micro-services/auth/db => ./db

replace github.com/qq51529210/micro-services/auth/reg => ./reg

replace github.com/qq51529210/micro-services/util => ../util
