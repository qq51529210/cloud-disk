module github.com/qq51529210/cloud-service/cloud-disk

go 1.15

require (
	github.com/qq51529210/cloud-service/util v0.0.0-00010101000000-000000000000
	github.com/qq51529210/http-router v0.0.0-20210609035309-7adf4360ac44
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/grpc v1.38.0
)

replace github.com/qq51529210/cloud-service/util => ../util
