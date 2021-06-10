package public

import (
	"io/ioutil"
	"os"
	"strings"
)

type FileEntity interface {
	Serialize() string
}

type FileEntitySerializer interface {
	Serialize(target FileEntity) string
	Deserialize(target string) FileEntity
}

type FileRepository struct {
	FileName   string
	Serializer FileEntitySerializer
}

func (repository *FileRepository) Save(entity FileEntity) error {
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

func (repository *FileRepository) GetAll() (*[]FileEntity, error) {
	body, err := ioutil.ReadFile(repository.FileName)
	if err != nil {
		return nil, err
	}

	entities := []FileEntity{}
	rows := strings.Split(string(body), "\n")
	for _, row := range rows {
		entities = append(entities, repository.Serializer.Deserialize(row))
	}
	return &entities, nil
}
