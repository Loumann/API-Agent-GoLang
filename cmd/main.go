package main

import (
	"awesomeProject/config"
	"awesomeProject/handler"
	"awesomeProject/models"
	"awesomeProject/repos"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	env := models.LoadEnv()
	cfg, err := config.ReadCfg()
	if err != nil {
		log.Fatal(err)
	}
	db, err := repos.DbConnection(env, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env.local: %v", err)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	println(dbPassword)

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

	r.POST("/agents", h.AddAgent)
	r.GET("/agents", h.GetAgents)
	r.PUT("/agents/:id", h.UpdateAgent)
	r.DELETE("/agents/:id", h.DeleteAgent)

	return r
}
