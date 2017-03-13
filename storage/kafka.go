package storage

import (
	"fmt"
)

type KafkaStorage struct {
}

func (c *KafkaStorage) Save() {
	fmt.Println("store to kafka")
}
