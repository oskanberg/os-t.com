package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const MIN_PULL_INTERVAL_MINUTES = 5

var lastGitPull time.Time

func gitPull() {
	log.Println("Git Pull")

	cmd := "git"
	args := []string{"pull"}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
	// don't process if we pulled recently
	if time.Since(lastGitPull).Minutes() < MIN_PULL_INTERVAL_MINUTES {
		return
	}

	lastGitPull = time.Now()
	gitPull()
}

func main() {
	port := flag.String("p", "8080", "The port on which to serve requests.")
	bGitPull := flag.Bool("update", false, "Whether or not to accept requests to git pull")
	flag.Parse()

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	if *bGitPull {
		var err error
		lastGitPull, err = time.Parse(time.RFC3339, "1970-01-01T00:00:00+00:00")
		if err != nil {
			return
		}

		pullHandler(nil, nil)
		http.HandleFunc("/api/v1/pull", pullHandler)
	}

	log.Printf("Serving on port %s\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
