package main

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Aadithya-J/microservices-ecommerce/account"

	// "github.com/Aadithya-J/microservices-ecommerce/catalog"
	// "github.com/Aadithya-J/microservices-ecommerce/order"
	"google.golang.org/grpc"
)

type Server struct {
	accountClient *account.Client
	// catalogClient *catalog.Client
	// orderClient   *order.Client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	accountConn, err := grpc.DialContext(ctx, accountUrl, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	accountClient := account.NewClient(accountConn)

	// catalogClient, err := catalog.NewClient(catalogUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	return nil, err
	// }

	// orderClient, err := order.NewClient(orderUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	// catalogClient.Close()
	// 	return nil, err
	// }

	return &Server{
		accountClient: accountClient,
		// catalogClient: catalogClient,
		// orderClient:   orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

func (s *Server) Account() AccountResolver {
	return &accountResolver{s}
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{s}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{s}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
