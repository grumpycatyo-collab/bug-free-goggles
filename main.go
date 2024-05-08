package main

import (
	"fmt"
	"github.com/grumpycatyo-collab/bug-free-goggles/core"
	"github.com/grumpycatyo-collab/bug-free-goggles/models"
	"sync"
	"time"
)

type CachedItems struct {
	lock     sync.RWMutex
	Projects []models.Project
	Users    []models.User
}

func main() {

	intervals := map[string]time.Duration{
		"every 5 minutes":  time.Minute * 5,
		"every 30 seconds": time.Second * 30,
	}

	projectsChannel := make(chan []models.Project)
	usersChannel := make(chan []models.User)
	for intervalName, interval := range intervals {
		go extract(interval, projectsChannel, usersChannel)
		fmt.Printf("Started fetching data %s\n", intervalName)
	}

	for {
		projectsData := <-projectsChannel
		fmt.Println(projectsData)
		usersData := <-usersChannel
		fmt.Println(usersData)
	}
}

var projects []models.Project
var users []models.User

func extract(interval time.Duration, chP chan<- []models.Project, chU chan<- []models.User) {
	for {
		projects, _ = core.GetProjects()
		users, _ = core.GetUsers()
		chP <- projects
		chU <- users
		time.Sleep(interval)
	}

}
