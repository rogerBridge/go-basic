package main

type Res struct {
	Name string `json:"name"`
	Hobbies []string `json:"hobbies"`
	Property string `json:"property"`
}

type Request struct {
	Code int `json:"code"`
	Property string  `json:"property"`
}
