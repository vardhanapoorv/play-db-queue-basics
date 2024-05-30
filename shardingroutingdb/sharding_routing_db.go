package shardingroutingdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

func Sharding_Routing_DB() {
	//db, err := sql.Open("mysql", "/mysql")
	//if err != nil {
	//panic(err)
	//}

	start := time.Now()
	db, err := sql.Open("mysql", "root:@tcp/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db2, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()
	if err := db2.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to week1 - Sharding Routing DB")
	// Users table exists on both shards (DB servers)
	// Range-based sharding is done to divide data
	// All records with username starting with greater than 'm' below to second shard

	// Other sharding options -
	// - Hash-based - Using a hash function
	// - Geo-based - Using location info
	// - Directory-based - Uses a look-up table to find the appropriate shard

	username := "johndoe" // example username
	shardNumber := getShardNumber(username)
	dbShardConn := db
	if shardNumber == 2 {
		dbShardConn = db2
	}

	var (
		userID    int
		email     string
		createdAt string // assuming datetime can be scanned as a string for simplicity
	)

	query := "SELECT user_id, email, created_at FROM Users WHERE username = ?"
	row := dbShardConn.QueryRow(query, username)

	// Scan for the values from the row
	if err := row.Scan(&userID, &email, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No results found.")
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("Shard%d has it - User ID: %d, Email: %s, Created At: %s\n", shardNumber, userID, email, createdAt)

	duration := time.Since(start)
	fmt.Println("Time spent", duration)
}

func getShardNumber(username string) int {
	firstChar := strings.ToLower(string(username[0]))

	// Compare the first character with 'm'
	if firstChar > "m" {
		return 2
	}
	return 1
}
