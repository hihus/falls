package icars

import (
	"database/sql"
	_ "mysql"
)

const (
	STEP         = 10000
	DB_NAME      = "icars"
	DB_USER_NAME = "root"
	DB_USER_PASS = "wanghe"
	DB_HOST      = "42.96.165.150"
	//type : samll,suv,huo | rent_day : 多值 | price : 单值 | time
	SQL = "select user_id,car_id,city_code,rent_day,type as car_type,price,UNIX_TIMESTAMP(time) as time from cars"
)

type DataSource struct {
	Step int
	sql  string
	db   *sql.DB
}

func InitDataSource() *DataSource {
	db, err := GetDBConnect()
	if err != nil {
		println(err.Error())
		return nil
	}
	return &DataSource{STEP, SQL, db}
}

func (self *DataSource) GetDataIterator() (*sql.Rows, error) {
	rows, err := self.sqlQuery(self.sql)
	if err != nil {
		println("Iterator : " + err.Error())
		return nil, err
	}
	return rows, nil
}

func (self *DataSource) sqlQuery(sql string) (*sql.Rows, error) {
	return self.db.Query(sql)
}

func GetDBConnect() (db *sql.DB, err error) {
	db_addr := DB_USER_NAME + ":" + DB_USER_PASS + "@tcp(" + DB_HOST + ":3306)/" + DB_NAME + "?charset=utf8"
	db, err = sql.Open("mysql", db_addr)
	return
}
