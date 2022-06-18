package backend

import (
	"P2/backend/game/handler"
	"P2/backend/infrastructure/http"
	"log"
)

func main() {
	registerGameHandlers()
	log.Fatal(http.ListenAndServe(":8080"))
}

func registerGameHandlers() {
	h := handler.GamesHandler{}
	http.RegisterHandler(h)
}
