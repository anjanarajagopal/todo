package data
import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func Init() error {
	os.Remove("./list.db")
	db, err := sql.Open("sqlite3", "./list.db")
	if err != nil {
		return err
	}

	sq := `create table task(id integer, description text);
	delete from task;`

	_, err = db.Exec(sq)
	if err != nil {
		return err
	}
	return nil
}