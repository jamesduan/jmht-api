package websocket

import (
	"log"
	"net/http"
	"strings"
	"time"

	"jmht-api/g"
	"jmht-api/redi"

	"github.com/go-redis/redis"
	"golang.org/x/net/websocket"
)

var (
	migrationConns         map[*websocket.Conn]string
	migrationCountConns    map[*websocket.Conn]string
	relatinMapConns        map[*websocket.Conn]string
	relatinMapPointConns   map[*websocket.Conn]string
	relationMapPubsub      *redis.PubSub
	relationMapPointPubsub *redis.PubSub
)

func checkAndDel(conns map[*websocket.Conn]string, conn *websocket.Conn, errMsg string) {
	if strings.Contains(errMsg, "broken pipe") {
		delete(conns, conn)
	}
}

func pingRelationMapPoint() {
	for {
		for key, _ := range relatinMapPointConns {
			errMarshl := websocket.Message.Send(key, "ping relation map point")
			if errMarshl != nil {
				log.Println(errMarshl.Error())
				// if strings.Contains(errMarshl.Error(), "broken pipe") {
				delete(relatinMapPointConns, key)
				// }
			}
		}
		time.Sleep(time.Millisecond * 30000)
	}
}

func pingMigration() {
	for {
		for key, _ := range migrationConns {
			errMarshl := websocket.Message.Send(key, "ping migrations")
			if errMarshl != nil {
				log.Println(errMarshl.Error())
				// if strings.Contains(errMarshl.Error(), "broken pipe") {
				delete(migrationConns, key)
				// }
			}
		}
		time.Sleep(time.Millisecond * 30000)
	}

}

func pingRelationMap() {
	for {
		for key, _ := range relatinMapConns {
			errMarshl := websocket.Message.Send(key, "ping relation map")
			if errMarshl != nil {
				log.Println(errMarshl.Error())
				// if strings.Contains(errMarshl.Error(), "broken pipe") || strings.Contains(errMarshl.Error(), "closed") || strings.Contains(errMarshl.Error(), "protocol wrong") {
				delete(relatinMapConns, key)
				// }
			}
		}
		time.Sleep(time.Millisecond * 30000)
	}
}

func pingMigrationCount() {
	for {
		for key, _ := range migrationCountConns {
			errMarshl := websocket.Message.Send(key, "ping migrations count")
			if errMarshl != nil {
				log.Println(errMarshl.Error())
				// if strings.Contains(errMarshl.Error(), "broken pipe") {
				delete(migrationCountConns, key)
				// }
			}
		}
		time.Sleep(time.Millisecond * 30000)
	}
}

func migrationHandler(ws *websocket.Conn) {
	migrationConns[ws] = ""
	var migrationQ = g.Config().Redis.MigrationMapQueue
	log.Println(migrationQ)

	for {
		// var migrations []*model.Migration

		data, err := redi.BlPop(migrationQ)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 200)
		for key, _ := range migrationConns {
			errMarshl := websocket.Message.Send(key, data)
			if errMarshl != nil {
				log.Println(errMarshl)
			}
		}
	}
}

func relationMapHandler(ws *websocket.Conn) {
	relatinMapConns[ws] = ""
	// var relationQ = g.Config().Redis.RelationMapQueue
	// log.Println(relationQ)
	log.Println(relatinMapConns)

	for {
		// msg, err := relationMapPubsub.ReceiveTimeout(20000 * time.Millisecond)
		msg := <-relationMapPubsub.Channel()

		for key, _ := range relatinMapConns {
			errMarshl := websocket.Message.Send(key, msg.Payload)
			if errMarshl != nil {
				log.Println(errMarshl)
				delete(relatinMapConns, key)
				// checkAndDel(relatinMapConns, key, errMarshl.Error())
			}
		}
	}
}

func relationMapPointHandler(ws *websocket.Conn) {
	relatinMapPointConns[ws] = ""
	// var relationPointQ = g.Config().Redis.RelationMapPointQueue

	// log.Println(relationPointQ)
	log.Println(relatinMapPointConns)
	// var message = make(chan *redis.Message)
	for {
		msg := <-relationMapPointPubsub.Channel()

		// log.Println(msg)
		// log.Println(msg.Payload)
		for key, _ := range relatinMapPointConns {
			errMarshl := websocket.Message.Send(key, msg.Payload)
			if errMarshl != nil {
				log.Println(errMarshl)
				delete(relatinMapPointConns, key)
			}
		}
	}

}

func migrationCountHandler(ws *websocket.Conn) {
	migrationCountConns[ws] = ""
	ws.MaxPayloadBytes = 20480
	var migrationCountQ = g.Config().Redis.MigrationCountQueue
	log.Println(migrationCountQ)

	for {
		data, err := redi.BlPop(migrationCountQ)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 200)
		for key, _ := range migrationCountConns {
			errMarshl := websocket.Message.Send(key, data)
			if errMarshl != nil {
				log.Println(errMarshl)
			}
		}
	}
}

func reloadHandler() {
	for {
		for key, _ := range relatinMapConns {
			errMarshl := websocket.Message.Send(key, "reload")
			if errMarshl != nil {
				log.Println(errMarshl.Error())
				// if strings.Contains(errMarshl.Error(), "broken pipe") || strings.Contains(errMarshl.Error(), "closed") || strings.Contains(errMarshl.Error(), "protocol wrong") {
				delete(relatinMapConns, key)
				// }
			}
		}
		time.Sleep(time.Hour * 3)
	}
}

func StartServer() {

	//初始化
	migrationConns = make(map[*websocket.Conn]string)
	relatinMapConns = make(map[*websocket.Conn]string)
	migrationCountConns = make(map[*websocket.Conn]string)
	relatinMapPointConns = make(map[*websocket.Conn]string)
	log.Println(g.RedisClient)
	relationMapPubsub = g.RedisClient.Subscribe(g.Config().Redis.RelationMapQueue)
	relationMapPointPubsub = g.RedisClient.Subscribe(g.Config().Redis.RelationMapPointQueue)

	http.Handle("/migration", websocket.Handler(migrationHandler))
	http.Handle("/migration/count", websocket.Handler(migrationCountHandler))
	http.Handle("/relationMap", websocket.Handler(relationMapHandler))
	http.Handle("/relationMap/point", websocket.Handler(relationMapPointHandler))
	http.Handle("/", http.FileServer(http.Dir("./templates")))

	go pingMigration()
	go pingMigrationCount()
	go pingRelationMap()
	go pingRelationMapPoint()
	go reloadHandler()

	err := http.ListenAndServe(g.Config().WebSocket.Port, nil)
	log.Println("WebSocket Listening on %s.", g.Config().WebSocket.Port)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	defer relationMapPubsub.Close()
	defer relationMapPointPubsub.Close()
}
