package public

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type FileEntity interface {
	EqualsId(id int) bool
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

func (repository *FileRepository) FindById(id int) (FileEntity, error) {
	entities, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	for _, entity := range *entities {
		if entity.EqualsId(id) {
			return entity, nil
		}
	}

	return nil, errors.New("Not Found")
}
