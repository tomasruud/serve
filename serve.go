package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	host := flag.String("h", "localhost", "Server host")
	port := flag.String("p", "3000", "Server port")
	root := flag.String("d", ".", "Directory to serve")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage: serve [OPTIONS] [DIRECTORY]\n\nOptions:")
		flag.PrintDefaults()
	}

	flag.Parse()

	addr := *host + ":" + *port

	dir := *root
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	abs, err := filepath.Abs(dir)
	if err != nil {
		abs = dir
	}

	log.Printf("Serving %q at %q", abs, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
