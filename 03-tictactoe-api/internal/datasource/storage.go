package datasource

import (
	"sync"
)

type Storage struct {
	data sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}
