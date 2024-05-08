package core

import (
	"encoding/json"
	"errors"
	"github.com/grumpycatyo-collab/bug-free-goggles/models"
	"golang.org/x/time/rate"

	//"github.com/grumpycatyo-collab/bug-free-goggles/models"
	"io"
	"net/http"
)

func GetProjects() ([]models.Project, error) {
	url := "https://app.asana.com/api/1.0/projects?workspace=1207269339457439"
	// A very primitive limiter
	limiter := rate.NewLimiter(15, 2)
	if !limiter.Allow() {
		return nil, errors.New("the API is at capacity, try again later")
	}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1207269339457428/1207269055600904:ad5baba08d5372e63291a8cf8181976d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var projects []models.Project
	projectsJSON := make(map[string][]models.Project)

	err := json.Unmarshal([]byte(string(body)), &projectsJSON)
	if err != nil {
		return nil, err
	}

	for _, project := range projectsJSON {
		projects = append(projects, project...)
	}

	return projects, nil
}

func GetUsers() ([]models.User, error) {
	url := "https://app.asana.com/api/1.0/users?workspace=1207269339457439"

	limiter := rate.NewLimiter(15, 2)
	if !limiter.Allow() {
		return nil, errors.New("the API is at capacity, try again later")
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1207269339457428/1207269055600904:ad5baba08d5372e63291a8cf8181976d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var users []models.User
	usersJSON := make(map[string][]models.User)

	err := json.Unmarshal([]byte(string(body)), &usersJSON)
	if err != nil {
		return nil, err
	}

	for _, project := range usersJSON {
		users = append(users, project...)
	}

	return users, nil
}
