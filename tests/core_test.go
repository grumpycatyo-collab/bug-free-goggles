package tests

import (
	"fmt"
	"github.com/grumpycatyo-collab/bug-free-goggles/core"
	"github.com/grumpycatyo-collab/bug-free-goggles/models"
	"reflect"
	"sync"
	"testing"
	"time"
)

var projects []models.Project
var users []models.User

func BenchmarkCore(b *testing.B) {
	var wg sync.WaitGroup

	stop := make(chan struct{})

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-time.Tick(1 * time.Second):
					projects, _ = core.GetProjects()
					users, _ = core.GetUsers()
					if projects != nil {
						fmt.Println("Projects returned successfully")
					}
					if users != nil {
						fmt.Println("Users returned successfully")
					}

				case <-stop:
					return
				}
			}
		}()
	}

	go func() {
		time.Sleep(10 * time.Second)
		close(stop)
	}()

	wg.Wait()
}

var ProjectTest = []models.Project{
	models.Project{"1207270196287897", "Project 2", "project"},
	models.Project{"1207269063496742", "Project 1", "project"},
}

func TestProjects(t *testing.T) {
	projects, _ = core.GetProjects()

	if reflect.DeepEqual(projects, ProjectTest) {
		t.Errorf("got %q, wanted %q", projects, ProjectTest)
	}
}

var UserTest = []models.User{
	models.User{"1207269034735706", "user3123@gmail.com", "user"},
	models.User{"1207269038212633", "user313@gmail.com", "user"},
	models.User{"1207269038333720", "user312f93@gmail.com", "user"},
}

func TestUsers(t *testing.T) {
	users, _ = core.GetUsers()

	if reflect.DeepEqual(users, UserTest) {
		t.Errorf("got %q, wanted %q", users, UserTest)
	}
}
