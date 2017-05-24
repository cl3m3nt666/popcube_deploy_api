package main

type Create_stack struct {
	Name  string `json:"name"`
  Description  string `json:"description"`
}


type Response_create struct {
	Message  string `json:"message"`
  Url  string `json:"url"`
  Stack_name  string `json:"stack_name"`
}

type Response_status struct {
	Message  string `json:"message"`
  Service  string `json:"service"`
}


type Response_remove struct {
	Message  string `json:"message"`
}
