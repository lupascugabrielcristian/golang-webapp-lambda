package persistance

type GetRobotsDAO struct {
	datasource RobotsDataSource
}

func GetRobotsDAOFactory(ds RobotsDataSource) *GetRobotsDAO {
	return &GetRobotsDAO{
		datasource: ds,
	}
}

func (dao *GetRobotsDAO) GetRobots(userId *string) map[string]string {
	return dao.datasource.GetRobots(userId)
}
