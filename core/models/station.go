package models

import "fmt"

type Station struct {
	Id			int64			`json: "id"`
	Name 		string		`json: "name"`
}

func (s *Station) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO station(name) values(?)")
	checkErr(err)

	res, err := stmt.Exec(s.Name)
	checkErr(err)

	return res.LastInsertId()
}

func (s *Station) Update() (err error) {

	stmt, err := db.Prepare("UPDATE station SET name=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(s.Name, s.Id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func (s *Station) Remove() (err error) {
	return StationRemove(s.Id)
}

func StationRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM station WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func StationGetById(id int) (station *Station, err error) {

	station = new(Station)
	rows, err := db.Query(fmt.Sprintf("SELECT name FROM station WHERE id=%d LIMIT 1", id))
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			var name string
			rows.Scan(&name)
			station.Name = name
		}
	}

	return
}

func StationGetAll() (stations []*Station, err error) {

	stations = make([]*Station, 0)	//[]

	rows, err := db.Query("SELECT * FROM station")
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			station := new(Station)
			rows.Scan(&station.Id, &station.Name)
			stations = append(stations, station)
		}
	}

	return
}