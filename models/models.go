package models

type User struct {
	GID          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type Project struct {
	GID          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}
