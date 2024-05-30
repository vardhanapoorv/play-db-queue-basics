package main

import (
	// conn_pool "week-1/connection_pool"
	// srdb "week-1/shardingroutingdb"
	"fmt"
	"time"
	rmq "week-1/rabbitmqtest"
	//sse "week-1/sse"
	proxysql "week-1/test-proxysql"
)

func main() {

	//conn_pool.Connection_pool()
	// srdb.Sharding_Routing_DB()
	// sse.StartServer()
	proxysql.ProxySQLDB()
}

func rmq_play() {

	rmq.Send()
	go func() {
		time.Sleep(time.Minute * 1)

		rmq.Send()
		rmq.Send()
	}()
	rmq.Receive()

	fmt.Println("Won't ever come here")

}
