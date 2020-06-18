package server

import (
	"time"

	"github.com/theSuess/keypub/pkg/auth"
	"github.com/theSuess/keypub/pkg/graph"
	"github.com/theSuess/keypub/pkg/graph/generated"
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/model"
	"github.com/theSuess/keypub/pkg/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
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

func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()
		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		log.V(2).Info("http request", "path", c.Request.URL.Path, "latency", latency, "status", status)
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

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Logger())
	resolver := &graph.Resolver{
		UserService:  service.User(db),
		KeyService:   service.Key(db),
		GroupService: service.Group(db),
	}
	r.POST("/query", auth.Middleware(), graphqlHandler(resolver))
	r.GET("/playground", playgroundHandler())
	r.POST("/auth", auth.Authenticate(db))
	log.Info("Starting server", "interface", s.Configuration.Interface, "url", "http://0.0.0.0"+s.Configuration.Interface)
	return r.Run(s.Configuration.Interface)
}
