package data

// SOURCES
type RobotsDataSourceRemote interface {
	GetRobots(id string) map[string]string
}

// IMPLEMENTATIONS
type GetRobotsRemote struct {
	RemoteDataSource RobotsDataSourceRemote
}

func (g GetRobotsRemote) GetRobots(id string) map[string]string {
	objData := g.RemoteDataSource.GetRobots(id)
	return objData
}
