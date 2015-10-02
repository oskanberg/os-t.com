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

func gitPull() {
	log.Println("Git Pull")

	cmd := "git"
	args := []string{"pull"}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}

}

func main() {
	port := flag.String("p", "8080", "The port on which to serve requests.")
	gitPullInterval := flag.Int("interval", 0, "Number of minutes between Git Pulls (<1 for off)")
	flag.Parse()

	if *gitPullInterval > 0 {

		gitPull()

		ticker := time.NewTicker(time.Duration(*gitPullInterval) * time.Minute)
		go func() {
			for {
				select {
				case <-ticker.C:
					gitPull()
				}
			}
		}()
	}

	log.Printf("Serving on port %s\n", *port)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("."))))
}
