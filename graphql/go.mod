module github.com/Aadithya-J/microservices-ecommerce/graphql

replace (
	github.com/Aadithya-J/microservices-ecommerce/account => ../account
	github.com/Aadithya-J/microservices-ecommerce/catalog => ../catalog
	github.com/Aadithya-J/microservices-ecommerce/order => ../order
)

go 1.24.4

require (
	github.com/99designs/gqlgen v0.17.76
	github.com/Aadithya-J/microservices-ecommerce/account v0.0.0-00010101000000-000000000000
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/vektah/gqlparser/v2 v2.5.30
	google.golang.org/grpc v1.74.2
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
