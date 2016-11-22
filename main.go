package main

import (
	jenkins "github.com/bndr/gojenkins"
	"github.com/shilkin/buildstatus/dispatcher"
	"github.com/shilkin/buildstatus/status"
	"github.com/shilkin/buildstatus/view"
	"log"
)

func main() {
	url := "http://localhost"
	login := "shilkin"
	password := "qwerty12345"

	client, err := jenkins.CreateJenkins(url, login, password).Init()
	if err != nil {
		log.Fatal(err)
	}

	reader := status.NewReader(client,
		status.ReaderOpts{
			TimeoutRead: 5000,
			Views:       []string{"Acceptance tests", "Local projects"},
		})
	render := view.NewStdoutRender()

	disp := dispatcher.NewDispatcher(reader, render)
	disp.Run()

}
