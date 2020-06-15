package server

import (
	"net/http"

	"github.com/theSuess/keypub/pkg/graph"
	"github.com/theSuess/keypub/pkg/graph/generated"
	"github.com/theSuess/keypub/pkg/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

type Configuration struct {
	Interface    string
	DatabasePath string
}

type Server struct {
	Configuration Configuration
}

func New(c Configuration) *Server {
	return &Server{
		Configuration: c,
	}
}

func (s *Server) Run() error {
	db, err := gorm.Open("sqlite3", s.Configuration.DatabasePath)
	if err != nil {
		return errors.Wrap(err, "opening database connection")
	}
	model.RunMigrations(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	return http.ListenAndServe(s.Configuration.Interface, nil)
}
