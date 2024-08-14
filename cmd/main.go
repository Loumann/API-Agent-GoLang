package main

import (
	"awesomeProject/config"
	"awesomeProject/handler"
	"awesomeProject/repos"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg, err := config.ReadCfg()
	if err != nil {
		log.Fatal(err)
	}
	db, err := repos.DbConnection(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := repos.GetRepository(db)
	h := handler.GetHadler(r)

	router := GetRouters(h)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Fatal::", err)
	}
	return

}

func GetRouters(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/create-user", h.AddAgent)
	r.GET("/agents", h.GetAgents)
	r.PUT("/agents/:id", h.UpdateAgent)
	r.DELETE("/agents/:id", h.DeleteAgent)

	return r
}
