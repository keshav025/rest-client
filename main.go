package main

import (
	"rest-client/restservice"

	"github.com/sirupsen/logrus"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
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

	logrus.Infof("post length: %v", len(m))
}
