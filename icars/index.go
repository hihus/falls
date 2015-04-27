package icars

import (
	"indexer"
	_ "time"
)

type DataFields struct {
	user_id   int
	car_id    int
	city_code int
	car_type  int
	rent_day  string
	price     float32
	time      int
}

type DataItem map[int]DataFields

type IcarsIndex struct {
	I *indexer.Indexer
	D DataItem
}

func InitIcarsIndex() *IcarsIndex {
	return &IcarsIndex{new(indexer.Indexer), make(DataItem)}
}

func BuildIcarsIndex() {
	index := InitIcarsIndex()
	data := InitDataSource()
	rows, err := data.GetDataIterator()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		item := DataFields{}
		err := rows.Scan(&item.user_id,
			&item.car_id,
			&item.city_code,
			&item.rent_day,
			&item.car_type,
			&item.price,
			&item.time)
		if nil != err {
			panic(err)
		}
		index.D[item.car_id] = item
	}
	for _, v := range index.D {
		print(v.car_id, " | ")
		print(v.user_id, " | ")
		print(v.city_code, " | ")
		print(v.car_type, " | ")
		print(v.rent_day, " | ")
		print(int(v.price), " | ")
		print(v.time)
		println("")

	}
}
