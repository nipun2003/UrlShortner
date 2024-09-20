package services

import (
	"fmt"
	"strconv"

	"github.com/go-zookeeper/zk"
	"github.com/nipun2003/url-shortner/db"
)

var shortIDGenerateService ShortIDGenerateService

type ShortIDGenerateService interface {
	GenerateUniqueID() (int, error)
}

type ShortIDGenerateServiceImpl struct {
	conn *zk.Conn
}

func NewShortIDGenerateService() ShortIDGenerateService {
	if shortIDGenerateService == nil {
		shortIDGenerateService = &ShortIDGenerateServiceImpl{conn: db.GetZookeeperConnection()}
	}
	return shortIDGenerateService
}
func (s *ShortIDGenerateServiceImpl) GenerateUniqueID() (int, error) {
	// Increment the ID by updating the znode's data
	data, _, err := s.conn.Get(db.GetZnodePath())
	if err != nil {
		return 0, err
	}

	// Convert current ID to integer, increment it and update the znode
	currentID := atoi(string(data))
	newID := currentID + 1

	_, err = s.conn.Set(db.GetZnodePath(), []byte(fmt.Sprintf("%d", newID)), -1)
	if err != nil {
		return 0, err
	}
	return newID, err
}

// Helper function to convert string to integer
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
