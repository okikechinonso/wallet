package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handler "github.com/okikechinonso/internals/handlers"
	"github.com/okikechinonso/internals/ports"
	redismemo "github.com/okikechinonso/internals/repository/redis_memo"
	"github.com/okikechinonso/pkg/database"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Hdl   *handler.Handler
	Redis ports.RedisRepository
}

func (s *Server) DefineRoute(router *gin.Engine) {
	apirouter := router.Group("/api/v1/wallets/:wallet_id")

	authorized := apirouter.Group("/")
	//authorized.Use(middleware.Authorize(s.Hdl.Wallet.Repo.GetWallet))
	authorized.POST("/credit", s.Hdl.Credit())
	authorized.POST("/debit", s.Hdl.Debit())
	authorized.GET("/balance", s.Hdl.GetBalance())
}

func (s *Server) setupRoute() *gin.Engine {
	r := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.DefineRoute(r)
	return r
}

func (s *Server) Start() {
	r := s.setupRoute()
	port := os.Getenv("PORT")

	go s.Redis.Sub()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable to start serve", err)
	}
}

func NewServer() *Server {
	log.SetFlags(log.Lshortfile | log.Llongfile)
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}
	hdl := handler.NewHandler()
	redis := database.ConnectRedisDB()
	redisrepo := redismemo.NewRedisRepository(redis)
	return &Server{
		Hdl:   hdl,
		Redis: redisrepo,
	}
}
