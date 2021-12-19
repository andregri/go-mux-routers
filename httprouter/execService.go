package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

// Execute a system command and return the output
func getCommandOutput(command string, args ...string) string {
	// args.. unpack arguments array into elements
	cmd := exec.Command(command, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Exec the command
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	// Wait it to finish
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	return stdout.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, getCommandOutput("/usr/bin/go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintln(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	log.Fatal(http.ListenAndServe(":8000", router))
}
