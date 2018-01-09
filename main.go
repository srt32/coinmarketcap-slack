package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// parse form body
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// grab values from form
	team := r.Form.Get("team_domain")
	user := r.Form.Get("user_name")
	cmd := r.Form.Get("command")
	args := strings.Fields(r.Form.Get("text"))

	if len(args) == 0 {
		fmt.Fprintln(w, "token is required")
		return
	}

	// log stuff
	fmt.Printf("%s %v from user %s (%s)\n", cmd, args, user, team)

	// perform the request
	url := "https://api.coinmarketcap.com/v1/ticker/" + args[0] + "/"
	res, err := http.Get(url)
	if err != nil {
		log.Printf("token required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("can't parse API response")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// TODO: serialize into struct
	fmt.Fprintf(w, "%v", body)
}
