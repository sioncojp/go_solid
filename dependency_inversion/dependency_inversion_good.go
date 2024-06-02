package dependency_inversion

import "fmt"

type DBConn interface {
	QuerySomeData() []string
}

func (db MySQL) QuerySomeData() []string {
	return []string{"inf1", "inf2", "inf3"}
}

type MyRepositoryGood struct {
	db DBConn
}

func (r MyRepositoryGood) GetData() []string {
	return r.db.QuerySomeData()
}

func runGood() {
	mysqlDB := MySQL{}
	repo := MyRepositoryGood{db: mysqlDB}
	fmt.Println(repo.GetData())
}
