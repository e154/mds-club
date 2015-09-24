package models

import (
	"reflect"
	"fmt"
)

type Station struct {
	Id			int64		`json: "id"`
	Name 		string		`json: "name"`
}

func (s *Station) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO station(name) values(?)")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(s.Name)
	if err != nil {
		checkErr(err)
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		checkErr(err)
		return
	}

	s.Id = id

	return
}

func (s *Station) Update() (err error) {

	stmt, err := db.Prepare("UPDATE station SET name=? where id=?")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(s.Name, s.Id)
	if err != nil {
		checkErr(err)
		return
	}

	_, err = res.RowsAffected()

	return
}

func (s *Station) Remove() (err error) {
	return StationRemove(s.Id)
}

func StationRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM station WHERE id=?")
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		checkErr(err)
		return
	}

	_, err = res.RowsAffected()

	return
}

func StationGet(val interface{}) (station *Station, err error) {

	station = new(Station)

	switch reflect.TypeOf(val).Name() {
	case "int":

		id := val.(int)
		rows, err := db.Query(fmt.Sprintf("SELECT name FROM station WHERE id=%d LIMIT 1", id))
		if err != nil {
			checkErr(err)
			return nil, err
		}
		defer rows.Close()

		station.Id = int64(id)
		for rows.Next() {
			if rows != nil {
				rows.Scan(&station.Name)
				return station, nil
			}
		}

	case "string":

		name := val.(string)
		rows, err := db.Query(fmt.Sprintf("SELECT id FROM station WHERE name='%s' LIMIT 1", name))
		if err != nil {
			checkErr(err)
			return nil, err
		}
		defer rows.Close()

		station.Name = name
		for rows.Next() {
			if rows != nil {
				rows.Scan(&station.Id)
				return station, nil
			}
		}
	}

	return nil, fmt.Errorf("station not found")
}

func StationGetAll() (stations []*Station, err error) {

	stations = make([]*Station, 0)	//[]

	rows, err := db.Query("SELECT * FROM station")
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if rows != nil {
			station := new(Station)
			rows.Scan(&station.Id, &station.Name)
			stations = append(stations, station)
		}
	}

	return
}