package repositories

import (
	"enigmacamp.com/gosql/logger"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type IFileRepository interface {
	Save(file multipart.File, fileName string) error
}

type FileRepository struct {
	path string
}

func (fr *FileRepository) Save(file multipart.File, fileName string) error {
	fileLocation := filepath.Join(fr.path, fileName)
	out, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Logger.Fatal().Msg(err.Error())
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logger.Logger.Fatal().Msg(err.Error())
		return err
	}
	return nil
}

func NewFileRepository(path string) IFileRepository {
	return &FileRepository{
		path: path,
	}
}
