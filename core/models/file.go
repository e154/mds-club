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

	stmt, err := db.Prepare(`INSERT INTO file(book_id, name, size, url) values(?,?,?,?)`)
	if err != nil {
		checkErr(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(f.Book_id, f.Name, f.Size, f.Url)
	if err != nil {
		checkErr(err)
		return
	}

	id, err = res.LastInsertId()
	f.Id = id

	return
}

func (f *File) Update() (err error) {

	_, err = db.Exec(fmt.Sprintf(`UPDATE file SET book_id=%d, name="%s", size="%s", url="%s" WHERE id=%d`, f.Book_id, f.Name, f.Size, f.Url, f.Id))
	checkErr(err)

	return
}

func (f *File) Remove() (err error) {
	return FileRemove(f.Id)
}

func FileRemove(id int64) (err error) {

	stmt, err := db.Prepare("DELETE FROM file WHERE id=?")
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

func FileGetAll() (files []*File, err error) {

	files = make([]*File, 0)	//[]

	rows, err := db.Query("SELECT * FROM file")
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

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
	if err != nil {
		checkErr(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if rows != nil {
			file := new(File)
			rows.Scan(&file.Id, &file.Book_id, &file.Name, &file.Size, &file.Url)
			files = append(files, file)
		}
	}

	return
}

func FileExist(name, url string) (int64, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id FROM file WHERE name='%s' AND url='%s'", name, url))
	if err != nil {
		checkErr(err)
		return 0, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}