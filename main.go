package main

import (
	"fmt"
	"github.com/grumpycatyo-collab/bug-free-goggles/core"
	"github.com/grumpycatyo-collab/bug-free-goggles/models"
)

var projects []models.Project
var users []models.User

func main() {
	projects, _ = core.GetProjects()
	fmt.Printf("%v", projects)
	users, _ = core.GetUsers()
	fmt.Printf("\n%v", users)

}
