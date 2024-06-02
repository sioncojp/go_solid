package dependency_inversion

import "fmt"

func (db MySQL) QuerySomeData() []string {
	return []string{"inf1", "inf2", "inf3"}
}

type MyRepositoryBad struct {
	db MySQL
}

func (r MyRepositoryBad) GetData() []string {
	return r.db.QuerySomeData()
}

func runBad() {
	mysqlDB := MySQL{}
	repo := MyRepositoryBad{db: mysqlDB}
	fmt.Println(repo.GetData())
}
