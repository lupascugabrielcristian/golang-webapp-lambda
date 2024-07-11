package application

import "fmt"

type Robot struct {
	RobotId string
	Name    string
}

func (r Robot) String() string {
	return fmt.Sprintf("RobotId: %s, Name: %s", r.RobotId, r.Name)
}
