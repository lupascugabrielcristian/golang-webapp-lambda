package application

import "fmt"

type Robot struct {
	Id   string
	Name string
}

func (r Robot) String() string {
	return fmt.Sprintf("Id: %s, Name: %s", r.Id, r.Name)
}
