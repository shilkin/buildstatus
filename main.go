package main

import (
	"flag"
	jenkins "github.com/bndr/gojenkins"
	"github.com/davecheney/gpio"
	"github.com/shilkin/buildstatus/dispatcher"
	"github.com/shilkin/buildstatus/status"
	"github.com/shilkin/buildstatus/view"
	"log"
	"strings"
	"time"
)

var (
	// jenkins options
	url      string
	login    string
	pass     string
	projects string
	timeout  time.Duration
	// pin numbers
	yellow int
	red    int
	green  int
	// common
	rendertype string
)

func init() {
	const (
		defUrl      = "http://localhost"
		defLogin    = "shilkin"
		defPass     = "qwerty12345"
		defProjects = "Acceptance tests, Local projects"
		defYellow   = gpio.GPIO10
		defRed      = gpio.GPIO9
		defGreen    = gpio.GPIO11
		defTimeout  = 500 * time.Millisecond
		defRender   = "stdout"
	)
	flag.StringVar(&url, "url", defUrl, "jenkins url 'http://jenkinsurl<:port>'")
	flag.StringVar(&login, "login", defLogin, "jenkins login")
	flag.StringVar(&pass, "pass", defPass, "jenkins pass")
	flag.StringVar(&projects, "projects", defProjects, "comma-separated list of projects")
	flag.DurationVar(&timeout, "timeout", defTimeout, "jenkins read timeout")
	flag.IntVar(&yellow, "ypin", defYellow, "yellow pin number")
	flag.IntVar(&red, "rpin", defRed, "red pin number")
	flag.IntVar(&green, "gpin", defGreen, "green pin number")
	flag.StringVar(&rendertype, "render", defRender, "render type (trafficlight, stdout)")
}

func main() {
	flag.Parse()

	client, err := jenkins.CreateJenkins(url, login, pass).Init()
	if err != nil {
		log.Fatal(err)
	}

	projectsList := strings.Split(projects, ",")
	if len(projectsList) == 1 && projectsList[0] == "" {
		projectsList = []string{}
	}

	log.Printf("watching projects: %#v", projectsList)

	reader := status.NewReader(client,
		status.ReaderOpts{
			TimeoutRead: timeout,
			Views:       projectsList,
		})

	var render view.Render
	switch rendertype {
	case "stdout":
		render = view.NewStdoutRender()
	case "trafficlight":
		render, err = view.NewRaspberryRender(view.RaspberryOpts{
			GpioYellow: yellow,
			GpioRed:    red,
			GpioGreen:  green,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	disp := dispatcher.NewDispatcher(reader, render)
	disp.Run()
}
