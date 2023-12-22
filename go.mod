module snapser/cloudecode

go 1.13

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.2
	golang.org/x/net v0.16.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.31.0
	snapser v0.0.0-00010101000000-000000000000

)

replace snapser => ./snapser
