package models
import "fmt"

type File struct {
	Id			int64		`json: "id"`
	Book_id		int64		`json: "book_id"`
	Name 		string		`json: "name"`
	Size 		string		`json: "size"`
	Url 		string		`json: "url"`
}

func (f *File) Save() (id int64, err error) {

	stmt, err := db.Prepare("INSERT INTO file(book_id, name, size, url) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(f.Book_id, f.Name, f.Size, f.Url)
	checkErr(err)

	id, err = res.LastInsertId()
	f.Id = id

	return
}

func (f *File) Update() (err error) {

	stmt, err := db.Prepare("UPDATE file SET book_id=?, name='?', size='2', url='2' WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(f.Book_id, f.Name, f.Size, f.Url, f.Id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func (f *File) Remove() (err error) {
	return FileRemove(f.Id)
}

func FileRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM file WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()

	return
}

func FileGetAll() (files []*File, err error) {

	files = make([]*File, 0)	//[]

	rows, err := db.Query("SELECT * FROM file")
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			file := new(File)
			rows.Scan(&file.Id, &file.Book_id, &file.Name, &file.Size, &file.Url)
			files = append(files, file)
		}
	}

	return
}

func FileGetAllByBook(book *Book) (files []*File, err error) {

	files = make([]*File, 0)	//[]

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM file WHERE book_id=%d", book.Id))
	checkErr(err)

	for rows.Next() {

		if rows != nil {
			file := new(File)
			rows.Scan(&file.Id, &file.Book_id, &file.Name, &file.Size, &file.Url)
			files = append(files, file)
		}
	}

	return
}