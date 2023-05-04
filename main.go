package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/vitorestevam/simple/simple"
)

type homePage struct {
	Posts []simple.Post
}

func main() {
	dataPost := `
[header]
title: First Post Ever
tags: demo, test, go
description: The first post ever made
[/header]

# Creating a blog should not be that hard

Even having many options of ways to create a blog such as blogger, tumblr, dev.com, Hugo and others, as a developer it should be obvious that i would try to make my own.
`

	post := simple.ParsePost(dataPost)
	temp := template.Must(template.ParseFiles("./templates/post.html"))
	f, err := os.Create(fmt.Sprintf("./output/posts/%s.html", post.Title))
	fmt.Println(err)
	err = temp.Execute(f, post)
	fmt.Println(err)

	data := homePage{
		Posts: []simple.Post{post, post, post, post},
	}
	temp = template.Must(template.ParseFiles("./templates/home.html"))
	f, err = os.Create("./output/index.html")
	fmt.Println(err)
	err = temp.Execute(f, data)
	fmt.Println(err)
}
