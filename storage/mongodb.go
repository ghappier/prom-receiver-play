package storage

import (
	"fmt"
)

type MongodbStorage struct {
}

func (c *MongodbStorage) Save() {
	fmt.Println("store to mongodb")
}
