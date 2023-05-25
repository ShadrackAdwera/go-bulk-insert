package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-bulk-insert/db/sqlc"
	"github.com/ShadrackAdwera/go-bulk-insert/worker"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	distro worker.TaskDistributor
	store  db.TxStore
}

func NewServer(distro worker.TaskDistributor, store db.TxStore) *Server {
	router := gin.Default()

	// add auth?

	srv := &Server{
		distro: distro,
		store:  store,
	}

	router.GET("/api/data", srv.getData)
	router.POST("/api/data", srv.uploadFile)

	srv.router = router

	return srv
}

func errJSON(err error) gin.H {
	return gin.H{"message": fmt.Sprintf("error occured: %v", err)}
}

func (srv *Server) StartServer(serverAddr string) error {
	return srv.router.Run(serverAddr)
}
