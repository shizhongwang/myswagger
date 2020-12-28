module github.com/shizhongwang/myswagger/product-api

go 1.13

require (
	github.com/PacktPublishing/Building-Microservices-with-Go-Second-Edition/product-api v0.0.0-20200205074745-5ec21a886558
	github.com/go-openapi/jsonreference v0.19.4 // indirect
	github.com/go-openapi/runtime v0.19.20
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/hashicorp/go-hclog v0.12.1
	github.com/nicholasjackson/env v0.6.0
	github.com/shizhongwang/myswagger/currency v0.0.0-20200329100342-3c14bf3f378d
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.3.5 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/grpc v1.28.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/shizhongwang/myswagger/currency => ../currency
