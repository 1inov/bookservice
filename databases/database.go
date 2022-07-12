package databases

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Интерфейс Database ...
type Database interface {
	GetDriver() string
	GetDataSrc(string) string
}

// Подключение к бд Open(название базы,объект базы)
func Open(database_name string, dbase Database) (*sql.DB, error) {
	driver := dbase.GetDriver()
	src := dbase.GetDataSrc(database_name)
	db, err := sql.Open(driver, src)
	Takeerr(err, "connect to database")
	return db, nil
}

// Закрыть подключение Close(объект базы)
func Close(db *sql.DB) error {
	err := db.Close()
	Takeerr(err, "disconnect from database")
	return nil
}

// Поиск по автору DBSearchByAuthor(объект базы,имя автора,режим поиска(1 или 0))
func DBSearchByAuthor(db *sql.DB, s string, like bool) ([]string, error) {
	var title []string
	i := 0
	s = setMode(s, like)
	sel, err := db.Query("SELECT t1.title FROM books t1  LEFT JOIN authors t2 on t1.authorid = t2.authorid WHERE t2.name " + s)
	Takeerr(err, "search by author")
	defer sel.Close()
	for sel.Next() {
		var st string
		err := sel.Scan(&st)
		title = append(title, st)
		i++
		Takeerr(err, "search by author")
	}
	if len(title) == 0 {
		title = append(title, "No authors found")
	}
	return title, nil
}

// Поиск по книге DBSearchByAuthor(объект базы,название книги,режим поиска(1 или 0))
func DBSearchByBook(db *sql.DB, s string, like bool) ([]string, error) {
	var name []string
	i := 0
	s = setMode(s, like)
	sel, err := db.Query("SELECT t2.name FROM books t1 LEFT JOIN authors t2 on t1.authorid = t2.authorid WHERE t1.title " + s)
	Takeerr(err, "search by book")
	defer sel.Close()
	for sel.Next() {
		var st string
		err := sel.Scan(&st)
		name = append(name, st)
		i++
		Takeerr(err, "search by book")
	}
	if len(name) == 0 {
		name = append(name, "No books found")
	}
	return name, nil
}

// Задать режим поиска
func setMode(s string, like bool) string {
	var mode string
	if like {
		mode = "LIKE '%" + s + "%'"
	} else {
		mode = "= '" + s + "'"
	}
	return mode
}

// Обработчик ошибок
func Takeerr(err error, method string) {
	if err != nil {
		log.Fatalf("Service can't "+method+" : %s", err)
	}
}
