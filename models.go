package main

type Skill struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Params []int  `json:"params"`
}

type Character struct {
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Skills []Skill `json:"skills"`
}
