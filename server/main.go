package main

import (
	"fmt"
	"github.com/AmirRezaM75/skull-king/app"
	"github.com/AmirRezaM75/skull-king/app/middlewares"
	"github.com/AmirRezaM75/skull-king/pkg/router"
	"github.com/AmirRezaM75/skull-king/pkg/validator"
	_userHandler "github.com/AmirRezaM75/skull-king/user/delivery/http"
	_userRepository "github.com/AmirRezaM75/skull-king/user/repository/mongo"
	tokenRepository "github.com/AmirRezaM75/skull-king/user/repository/redis"
	_userService "github.com/AmirRezaM75/skull-king/user/service"
	"github.com/AmirRezaM75/skull-king/ws"
	"log"
	"net/http"
	"os"
)

func main() {
	application := app.App{}
	application.LoadEnvironments()

	client, cancel, disconnect := application.InitDatabase()

	redis := application.InitRedis()

	defer cancel()
	defer disconnect()

	var userRepository = _userRepository.NewMongoUserRepository(
		client.Database(os.Getenv("MONGODB_DATABASE")),
	)

	var tokenRepository = tokenRepository.NewRedisTokenRepository(redis)

	var userService = _userService.NewUserService(userRepository, tokenRepository)

	v := validator.NewValidator()

	r := router.NewRouter()
	r.Middleware(middlewares.CorsPolicy{})

	_userHandler.NewUserHandler(userService, v, r)

	hub := ws.NewHub()

	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	http.HandleFunc("/ws/join", wsHandler.Join)

	fmt.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
