package models

import "fmt"

type Station struct {
	Id			int			`json: "id"`
	Name 		string		`json: "name"`
}

func (s *Station) Save() {

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