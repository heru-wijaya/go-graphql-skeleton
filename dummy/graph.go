package dummy

import (
	"github.com/99designs/gqlgen/graphql"
)

type Server struct {
}

func NewGraphQLServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
