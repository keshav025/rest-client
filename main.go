package main

import (
	"rest-client/restservice"

	"github.com/sirupsen/logrus"
)

type Post struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostList []Post

func main() {
	c := restservice.GetRestClient("https://jsonplaceholder.typicode.com/posts", "", "")
	m := PostList{}
	err := c.Get(nil, &m)
	if err != nil {
		logrus.Errorf("error in get call : %v", err)
	}

	logrus.Info(m)
}
