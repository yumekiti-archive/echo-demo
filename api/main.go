package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/echo-demo/config"
	"github.com/yumekiti/echo-demo/infrastructure"
	"github.com/yumekiti/echo-demo/interface/handler"
	"github.com/yumekiti/echo-demo/usecase"
)

func main() {
	taskRepository := infrastructure.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.InitRouting(e, taskHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
