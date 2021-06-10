module github.com/qq51529210/cloud-service/authentication

go 1.15

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/qq51529210/cloud-service/util v0.0.0
	github.com/qq51529210/http-router v0.0.0-20210609035309-7adf4360ac44
	github.com/qq51529210/jwt v0.0.0-20210531120038-6eb2212a7688
	github.com/qq51529210/log v0.0.0-20210529132539-d2d52fbd5103
	github.com/qq51529210/redis v0.0.0-20210610092729-dbe50e9924d8
	github.com/qq51529210/uuid v0.0.0-20210410083004-ce2b0df9936f
)

replace github.com/qq51529210/cloud-service/util => ../util
