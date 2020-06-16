package server

import (
	"net/http"

	"github.com/theSuess/keypub/pkg/graph"
	"github.com/theSuess/keypub/pkg/graph/generated"
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

var log = logf.Log.WithName("server")

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
	log.Info("Starting server")
	log.Info("Opening database connection", "databasePath", s.Configuration.DatabasePath)
	db, err := gorm.Open("sqlite3", s.Configuration.DatabasePath)
	if err != nil {
		return errors.Wrap(err, "opening database connection")
	}
	log.Info("Running migrations")
	model.RunMigrations(db)
	db.LogMode(true)
	db.SetLogger(logf.WrapGorm(log))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Info("Starting server", "interface", s.Configuration.Interface, "url", "http://0.0.0.0"+s.Configuration.Interface)
	return http.ListenAndServe(s.Configuration.Interface, nil)
}
