package main

import (
	jenkins "github.com/bndr/gojenkins"
	"log"
	"github.com/shilkin/buildstatus/summary"
	"github.com/shilkin/buildstatus/view"
	"github.com/shilkin/buildstatus/dispatcher"
)

func main() {
	url := "http://jenkins11.mailbuild-1.dev.search.km"
	login := "shilkin"
	password := "qwerty12345"

	client, err := jenkins.CreateJenkins(url, login, password).Init()
	if err != nil {
		log.Fatal(err)
	}

	reader := summary.NewReader(client,
		summary.ReaderOpts{
			TimeoutRead: 5000,
			Views: []string{"Local projects", "Acceptance tests"}})
	render := view.NewStdoutRender()

	disp := dispatcher.NewDispatcher(reader, render)
	disp.Run()

}


