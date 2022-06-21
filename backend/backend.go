package main

import (
	game "P2/backend/game/handler"
	"P2/backend/infrastructure/http"
	matchmaking "P2/backend/matchmaking/handler"
	"P2/backend/user/handler"
	"errors"
	"fmt"
	gohttp "net/http"
	"os"
)

func main() {
	registerGameHandlers()

	err := http.ListenAndServe(":8080")
	if errors.Is(err, gohttp.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Println(fmt.Sprintf("error starting server - err: %s", err))
		os.Exit(1)
	}
}

func registerGameHandlers() {
	// game
	gh := game.GamesHandler{}
	http.RegisterHandler(gh)

	ah := game.ActionsHandler{}
	http.RegisterHandler(ah)

	sh := game.StatusHandler{}
	http.RegisterHandler(sh)

	// matchmaking
	mh := matchmaking.MatchmakingHandler{}
	http.RegisterHandler(mh)

	th := matchmaking.TicketHandler{}
	http.RegisterHandler(th)

	// user
	uh := handler.UsersHandler{}
	http.RegisterHandler(uh)

	sth := handler.StatsHandler{}
	http.RegisterHandler(sth)
}
