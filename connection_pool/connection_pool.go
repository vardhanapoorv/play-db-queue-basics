package connection_pool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

type SafeConns struct {
	mu      sync.Mutex
	dbConns []*sql.DB
}

func Connection_pool() {
	//db, err := sql.Open("mysql", "/mysql")
	//if err != nil {
	//panic(err)
	//}
	conns := SafeConns{dbConns: []*sql.DB{}}
	for i := 0; i < 10; i++ {
		db, _ := sql.Open("mysql", "root:@tcp/mysql")
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		conns.dbConns = append(conns.dbConns, db)
	}
	fmt.Println("Welcome to week1 - Connection Pool")
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//fmt.Println(dbConns)
			conns.mu.Lock()
			val := conns.dbConns
			for len(val) == 0 {
				conns.mu.Unlock()
				time.Sleep(8 * time.Microsecond)
				conns.mu.Lock()
				val = conns.dbConns
			}
			dbConn := conns.dbConns[0]
			conns.dbConns = conns.dbConns[1:]
			conns.mu.Unlock()
			dbConn.Exec("Select SLEEP(0.01);")
			conns.mu.Lock()
			conns.dbConns = append(conns.dbConns, dbConn)
			conns.mu.Unlock()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Time spent", duration)
}
