package main

import (
	"article/internal/configs"
	articleHandler "article/internal/handler/articles"
	"article/internal/models/articles"
	articleRepo "article/internal/repositories/articles"
	articleSvc "article/internal/services/articles"
	"article/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "article/internal/docs" // <-- wajib
)

// @title Articles API
// @version 1.0
// @description API untuk CRUD artikel
// @host localhost:9876
// @BasePath /api/v1
func main(){
	var (
		cfg *configs.Config
	)
 // Swagger endpoint
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal Inisialisasi Config",err)
	}


	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil{
		log.Fatal("failed to connect to database, err: %+v", err)
	}

	db.AutoMigrate(&articles.Article{})

	r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	articleRepo:=articleRepo.NewRepository(db)
	articleSvc:=articleSvc.NewService(cfg,articleRepo)
	articleHandler:=articleHandler.NewHandler(r, articleSvc)
	articleHandler.RegisterRoute()
	
	r.Run(cfg.Service.Port)
}