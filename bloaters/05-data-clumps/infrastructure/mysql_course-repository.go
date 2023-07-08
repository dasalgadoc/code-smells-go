package infrastructure

import (
	"dasalgadoc.com/code_smell_go/bloaters/05-data-clumps/domain"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	user     = "root"
	password = "root"
	host     = "localhost"
	port     = "3306"
	database = "go_testing"
)

type MySQLCourseRepository struct {
	db *sql.DB
}

func NewMySQLCourseRepository() *MySQLCourseRepository {
	connectionString := gerConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return &MySQLCourseRepository{
		db: db,
	}
}

func gerConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}

func (m *MySQLCourseRepository) Save(course domain.Course) error {
	stmt, err := m.db.Prepare("INSERT INTO course(id, name, duration) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(course.Id(), course.Name(), course.Duration())
	if err != nil {
		return err
	}

	return nil
}
