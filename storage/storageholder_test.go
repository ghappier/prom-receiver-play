package storage

import (
	"testing"
)

func TestAttach(t *testing.T) {
	kafka := new(KafkaStorage)
	mongodb := new(MongodbStorage)

	holder := NewStorageHolder()
	holder.Attach(kafka)
	holder.Attach(mongodb)
	holder.Save()

	if holder.storages.Len() == 0 {
		t.Error("error")
	}
	/*
		if len(holder.storages) == 0 {
			t.Error("error")
		}
	*/
}
