package repositories

import (
	"database/sql"
	"fmt"
	"go-api-ddd/consumer/internal/domain"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
    dsn := fmt.Sprintf("root:root@tcp(db:3306)/appdb")

    var err error
    for i := 0; i < 10; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil && db.Ping() == nil {
            log.Println("Connected to MySQL")
            return
        }

        log.Printf("Waiting for MySQL (%d/10)...", i+1)
        time.Sleep(2 * time.Second)
    }

    log.Fatal("Could not connect to MySQL:", err)
}

func SaveMessageToDB(msg domain.Message) error {
    _, err := db.Exec("INSERT INTO messages (content) VALUES (?)", msg.Content)
    return err
}
