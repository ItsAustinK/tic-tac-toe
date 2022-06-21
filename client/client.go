package main

import (
	"P2/client/ticket"
	"P2/client/user"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var curUser user.User
var curTicket ticket.Ticket

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tic Tac Toe - Shell")
	fmt.Println("----------------------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		routeCommand(text)
	}
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
	}
}

func readUserIds() (map[string]struct{}, error) {
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

	var ids map[string]struct{}
	err = json.Unmarshal(u, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func writeUserIds(ids map[string]struct{}) error {
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

	for i := range ids {
		fmt.Println(fmt.Sprintf("%d - %s", i, ids[i]))
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

	var u user.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return err
	}
	curUser = u

	ids, err := readUserIds()
	if ids == nil {
		ids = map[string]struct{}{}
	}
	ids[id] = struct{}{}

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

	var t ticket.Ticket
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return err
	}
	curTicket = t

	fmt.Println(fmt.Sprintf("queued up! \n%+v", curTicket))
	return nil
}
