module github.com/nicholasjackson/building-microservices-youtube/product-api

go 1.13

require (
	github.com/PacktPublishing/Building-Microservices-with-Go-Second-Edition/product-api v0.0.0-20200205074745-5ec21a886558
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/runtime v0.19.20
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-swagger/go-swagger v0.25.0 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/hashicorp/go-hclog v0.12.1
	github.com/nicholasjackson/building-microservices-youtube/currency v0.0.0-20200329100342-3c14bf3f378d
	github.com/nicholasjackson/env v0.6.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.28.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/nicholasjackson/building-microservices-youtube/currency => ../currency
