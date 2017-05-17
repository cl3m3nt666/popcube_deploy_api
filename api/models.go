package main

type Create_stack struct {
	Name  string `json:"name"`
  Description  string `json:"description"`
}


type Response struct {
	Message  string `json:"Message"`
  Url  string `json:"url"`
  Stack_name  string `json:"stack_name"`
}
