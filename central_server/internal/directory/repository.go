package directory

import (
	"errors"
	"fmt"
)

var (
	ErrExists = errors.New("directory: index already exists")
)

import "fmt"

type DirectoryRepository interface {
	SaveIndex(index string) error
	GetIndexTable() (map[int]string, error)
	SearchFile(filename string) (error)
}
type defaultMapRepo struct {
	indexTable        *map[string]string
	
}

func (d defaultMapRepo) SaveIndex(index string) error {
	table := *d.indexTable
	
	for key, indexTable := range table {
		if indexTable == index {
			return fmt.Errorf("%w id:%d", ErrExists, key)
		}
	}
	table[] = index
}

func (d defaultMapRepo) GetIndexTable() (map[int]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultMapRepo) SearchFile(filename string) error {
	//TODO implement me
	panic("implement me")
}
