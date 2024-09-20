package db

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
)

var conn *zk.Conn
var path string

func InitZookeeper() {
	host := os.Getenv("ZK_HOST")
	port := os.Getenv("ZK_PORT")
	path = "/url_shortener/id"
	// Initialize Zookeeper
	var err error
	var events <-chan zk.Event
	conn, events, err = zk.Connect([]string{host + ":" + port}, time.Second*20, zk.WithLogInfo(false))
	if err != nil {
		panic(err)
	}
	// Wait for the connection to be established
	for event := range events {
		if event.State == zk.StateHasSession {
			fmt.Println("Connected to ZooKeeper")
			break
		}
	}
	// Ensure the node exists
	err = createPathRecursively(path)
	if err != nil {
		panic(err)
	}
}

func GetZookeeperConnection() *zk.Conn {
	return conn
}

func CloseZookeeperConnection() {
	if conn != nil {
		conn.Close()
	}
}

func createPathRecursively(path string) error {
	parts := strings.Split(path, "/")
	for i := 1; i < len(parts); i++ {
		subPath := strings.Join(parts[:i+1], "/")
		exists, _, err := conn.Exists(subPath)
		if err != nil {
			return err
		}
		if !exists {
			_, err := conn.Create(subPath, []byte(""), zk.FlagPersistent, zk.WorldACL(zk.PermAll))
			if err != nil && err != zk.ErrNodeExists {
				return err
			}
		}
	}
	return nil
}

func GetZnodePath() string {
	return path
}
