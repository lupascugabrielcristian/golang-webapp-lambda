package framework

type DBService struct {
	Robots []map[string]string
}

func (db DBService) GetRobotsForUser(userId string) map[string]string {
	// Filter Robots map so that id == userId
	for _, value := range db.Robots {
		if value["id"] == userId {
			return value
		}
	}

	return nil
}

type RobotsDataGateway struct {
	Db *DBService
}

func (r RobotsDataGateway) GetRobots(id string) map[string]string {
	robotsMap := r.Db.GetRobotsForUser(id)
	return robotsMap
}
