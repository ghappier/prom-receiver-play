package storage

import (
	"container/list"
)

type StorageHolder struct {
	storages *list.List
}

func NewStorageHolder() *StorageHolder {
	s := new(StorageHolder)
	s.storages = list.New()
	return s
}

func (s *StorageHolder) Save() {
	for ob := s.storages.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Storage).Save()
	}
}

func (s *StorageHolder) Attach(storage Storage) { //注册
	s.storages.PushBack(storage)
}
