package account

import (
	"github.com/99designs/gqlgen/graphql"
	account "github.com/heru-wijaya/go-grpc-skeleton/server"
)

type Server struct {
	accountClient *account.Client
}

func NewGraphQLServer(accountUrl string) (*Server, error) {
	// Connect to account service
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		accountClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
