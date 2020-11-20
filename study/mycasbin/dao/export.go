package dao

var (

	User   = &UserDao{}
	Casbin = &CasbinDao{}
)

const (
	DBNAME = "mysql"
	DNS = "lambdacal:yskj2407@tcp(10.0.1.58:3306)/wzz_db"
)

