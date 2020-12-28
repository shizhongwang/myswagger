module github.com/shizhongwang/myswagger/chatroom-api

go 1.13

require (
	github.com/go-openapi/runtime v0.19.20
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/hashicorp/go-hclog v0.12.1
	github.com/nicholasjackson/env v0.6.0
	github.com/shizhongwang/myswagger/currency v0.0.0-20201228114024-09b64b667544
	github.com/shizhongwang/myswagger/product-api v0.0.0-20201228114024-09b64b667544
	google.golang.org/grpc v1.28.0
)

replace github.com/shizhongwang/myswagger/currency => ../currency
