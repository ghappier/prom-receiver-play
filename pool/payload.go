package pool

import (
	"fmt"

	"github.com/lizq/prom-receiver-play/storage"
)

type PayloadCollection struct {
	//WindowsVersion string    `json:"version"`
	//Token          string    `json:"token"`
	Payloads []Payload `json:"data"`
}

type Payload struct {
	Id   string
	Name string
	// [redacted]
	holder *storage.StorageHolder
}

func (p *Payload) Save() error {
	fmt.Println("save payload")
	p.holder.Save()
	return nil
}
