package public

import (
	"io/ioutil"
	"os"
	"strings"
)

type Serializable interface {
	Serialize() string
}

type FileEntitySerializer interface {
	Serialize(target Serializable) string
	Deserialize(target string) interface{}
}

type FileRepository struct {
	FileName   string
	Serializer FileEntitySerializer
}

func (repository *FileRepository) Save(entity Serializable) error {
	f, err := os.OpenFile(repository.FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(repository.Serializer.Serialize(entity)); err != nil {
		return err
	}
	return nil
}

func (repository *FileRepository) GetAll() (*[]interface{}, error) {
	body, err := ioutil.ReadFile(repository.FileName)
	if err != nil {
		return nil, err
	}

	entities := []interface{}{}
	rows := strings.Split(string(body), "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		entities = append(entities, repository.Serializer.Deserialize(row))
	}
	return &entities, nil
}
