package profile

import (
	"sync"

	"fakebook.com/project/models"
)

var (
	list []models.User
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initializeList)
}

func initializeList() {
	// this is simply a placeholder, and to show functionality of requests
	// in the future, we probably would not even need this
	list = []models.User{
		{
			Name:     "Melissa",
			Age:      29,
			HomeTown: "Rochester, MN",
			Job:      "Intern",
			Username: "brownm26csp",
		},
		{
			Name:     "Avery",
			Age:      22,
			HomeTown: "Place, State",
			Job:      "Professional",
			Username: "averytribbett",
		},
		{
			Name:     "Cade",
			Age:      23,
			HomeTown: "different place, different state",
			Job:      "Super professional",
			Username: "cadegithub",
		},
		{
			Name:     "Youssef",
			Age:      28,
			HomeTown: "Yet another place, yet another state",
			Job:      "Super ultra professional extraordinaire",
			Username: "youssefgithub",
		},
	}
}

func Get() []models.User {
	return list
}

/*
endpoint ideas for a user profile:

Requirements:
1. getFriendList
2. getUserPosts
3. probably something regarding logging out / logging in?

Optional:
1. getUserPhotos


*/
