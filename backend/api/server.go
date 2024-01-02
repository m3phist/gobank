package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	db "github.com/m3phist/gobank/backend/db/sqlc"
	"github.com/m3phist/gobank/backend/utils"
)

type Server struct {
	queries *db.Store
	router  *gin.Engine
	config  *utils.Config
}

var tokenController *utils.JWTToken

func NewServer(envPath string) *Server {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load env config: %v", err))
	}

	conn, err := sql.Open(config.DB_driver, config.DB_source_live)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	tokenController = utils.NewJWTToken(config)

	q := db.NewStore(conn)
	g := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", currencyValidator)
	}

	g.Use(cors.Default())

	return &Server{
		queries: q,
		router:  g,
		config:  config,
	}
}

func (s *Server) Start(port int) {
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gobank API",
		})
	})

	User{}.router(s)
	Auth{}.router(s)
	Account{}.router(s)

	s.router.Run(fmt.Sprintf(":%v", port))
}
