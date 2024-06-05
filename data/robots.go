package data

type GetRobotsRemote struct {
}

func (g GetRobotsRemote) GetRobots() []string {
	return []string{"Robot 1", "Robot2"}
}
