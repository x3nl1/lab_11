package main

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0,lte=120"`
}

type Numbers struct {
	Values []int `json:"values" binding:"required"`
}