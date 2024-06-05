package framework

type DBService struct {
}

func (db DBService) GetRobotsForUser(userId string) map[string]string {
	var robotsMap = map[string]string{
		"id":   "abc1",
		"name": "Robot1",
	}
	return robotsMap
}

type RobotsDataGateway struct {
	Db *DBService
}

func (r RobotsDataGateway) GetRobots(id string) map[string]string {
	robotsMap := r.Db.GetRobotsForUser(id)
	return robotsMap
}
