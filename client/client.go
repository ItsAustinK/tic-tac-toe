package main

import (
	"P2/client/game"
	"P2/client/ticket"
	"P2/client/user"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var curUser *user.User
var curTicket *ticket.Ticket
var curGame *game.Game
var ticketTimer *time.Timer
var gameTimer *time.Timer

func main() {
	fmt.Println("Tic Tac Toe - Shell")
	fmt.Println("----------------------------------")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(fmt.Sprintf("\n\nExiting - signal: %s", sig))
		os.Exit(0)
	}()

	// TODO: Read terminal in non-blocking way
	for {
		readTerminalInput()
	}
}

func readTerminalInput() {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	routeCommand(text)
}

func routeCommand(cmd string) {
	if strings.Contains(cmd, "-users") {
		err := printUserIds()
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(cmd, "-login") {
		split := strings.Split(cmd, "=")
		if len(split) < 2 {
			fmt.Println("invalid login cmd")
			return
		}

		id := split[1]
		id = strings.ReplaceAll(id, "\n", "")
		id = strings.ReplaceAll(id, "\r", "")
		err := userLogin(id)
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(cmd, "-queue") {
		err := queueForMatch()
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(cmd, "-game") {
		err := getGame(curTicket.GameId)
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(cmd, "-join") {
		err := joinGame(curUser.Id, curTicket.GameId)
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(cmd, "-action") {
		split := strings.Split(cmd, "=")
		if len(split) < 2 {
			fmt.Println("invalid login cmd")
			return
		}

		id := split[1]
		id = strings.ReplaceAll(id, "\n", "")
		id = strings.ReplaceAll(id, "\r", "")
		loc, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		err = makeBoardAction(curGame.Id, curGame.Token, loc)
		if err != nil {
			panic(err)
		}
	}
}

func readUserIds() (map[string]bool, error) {
	fileName, err := filepath.Abs("client/storage/users.txt")
	if err != nil {
		return nil, err
	}

	u, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var ids map[string]bool
	err = json.Unmarshal(u, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func writeUserIds(ids map[string]bool) error {
	b, err := json.Marshal(ids)
	if err != nil {
		return err
	}

	fileName, err := filepath.Abs("client/storage/users.txt")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, b, fs.ModePerm)
}

func printUserIds() error {
	ids, err := readUserIds()
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		fmt.Println("no saved users")
		return nil
	}

	for k, v := range ids {
		fmt.Println(fmt.Sprintf("%s - in use? %t", k, v))
	}

	return nil
}

func userLogin(id string) error {
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/users", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var u *user.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return err
	}
	curUser = u

	ids, err := readUserIds()
	if ids == nil {
		ids = map[string]bool{}
	}
	ids[u.Id] = true

	err = writeUserIds(ids)
	if err != nil {
		fmt.Println("failed to write user ids to local storage")
	}

	fmt.Println(fmt.Sprintf("successful login! \n%+v", curUser))
	return nil
}

func queueForMatch() error {
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/matchmake", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", curUser.Id)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var t *ticket.Ticket
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return err
	}
	curTicket = t

	go getTicketStatusLongPoll()

	fmt.Println(fmt.Sprintf("queued up! retrieved ticket \n%+v", curTicket))
	return nil
}

func getTicketStatusLongPoll() {
	if curTicket == nil {
		return
	}

	//fmt.Println("long polling ticket status")
	ticketTimer = time.NewTimer(time.Second) // long poll - query every second
	<-ticketTimer.C

	err := getTicketStatus(curTicket.Id)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ticket status- %s", err.Error()))
		return
	}

	if curTicket.Status == string(ticket.Complete) {
		fmt.Println("stopping ticket long polling - ticket status is complete!")
	} else {
		go getTicketStatusLongPoll()
	}
}

func getTicketStatus(id string) error {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/tickets", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var t *ticket.Ticket
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return err
	}
	curTicket = t

	//fmt.Println(fmt.Sprintf("retrieved ticket! \n%+v", curTicket))
	return nil
}

func joinGame(uid, gid string) error {
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/games", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("uid", uid)
	q.Add("gid", gid)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var g *game.Game
	err = json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		return err
	}

	curGame = g
	curGame.Init()
	curGame.Render()

	go getGameStatusLongPoll()
	fmt.Println(fmt.Sprintf("joined game! retrieved update game \n%+v", curGame))
	return nil
}

func getGameStatusLongPoll() {
	if curGame == nil {
		return
	}

	//fmt.Println("long polling game object")
	gameTimer = time.NewTimer(time.Second) // long poll - query every second
	<-gameTimer.C

	err := getGame(curGame.Id)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get game object- %s", err.Error()))
		return
	}

	if curGame.Status == string(game.Complete) {
		fmt.Println("stopping game long polling -game is complete!")
	} else {
		go getGameStatusLongPoll()
	}
}

func getGame(id string) error {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/games", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var g *game.Game
	err = json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		return err
	}

	// TODO: Once terminal input is non-blocking, implement time.Ticker into for loop & Game obj for individual tick rates
	var render bool
	if curGame != nil {
		if curGame.Token != g.Token { // check for delta in game object
			render = true
		}
	} else {
		render = true // render new obj
	}
	curGame = g
	if render {
		curGame.Render()
	}

	//fmt.Println(fmt.Sprintf("retrieved game! \n%+v", curGame))
	return nil
}

func makeBoardAction(id, token string, location int) error {
	row := location % curGame.Board.Row
	col := location / curGame.Board.Col
	a := game.Action{
		PlayerId: id,
		Position: [2]int{row, col},
	}

	b, err := json.Marshal(a)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/actions", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", id)
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()

	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var g *game.Game
	err = json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		return err
	}
	curGame = g

	if curGame != nil {
		curGame.Render()
	}

	fmt.Println(fmt.Sprintf("made action! retrieved updated game \n%+v", curGame))
	return nil
}
