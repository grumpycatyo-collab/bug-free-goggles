package core

import (
	"encoding/json"
	"fmt"
	"github.com/grumpycatyo-collab/bug-free-goggles/models"

	//"github.com/grumpycatyo-collab/bug-free-goggles/models"
	"io"
	"net/http"
)

func GetProjects() error {
	url := "https://app.asana.com/api/1.0/projects?workspace=1207269339457439"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1207269339457428/1207269055600904:ad5baba08d5372e63291a8cf8181976d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	pingJSON := make(map[string][]models.User)

	err := json.Unmarshal([]byte(string(body)), &pingJSON)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n %v", pingJSON)
	return nil
}

func GetUsers() error {
	url := "https://app.asana.com/api/1.0/users?workspace=1207269339457439"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1207269339457428/1207269055600904:ad5baba08d5372e63291a8cf8181976d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	pingJSON := make(map[string][]models.User)

	err := json.Unmarshal([]byte(string(body)), &pingJSON)
	if err != nil {
		panic(err)
	}

	for _, val := range pingJSON {
		fmt.Println(val)
	}
	return nil
}
