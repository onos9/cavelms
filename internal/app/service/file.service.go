package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
	"github.com/gin-gonic/gin"
)

const MAX_UPLOAD_SIZE = 1000 * 1024 * 1024 // 1GB

type fileService interface {
	CreateFile(ctx context.Context, input model.NewFile) (*model.File, error)
	UpdateFile(ctx context.Context, data interface{}) (*model.File, error)
	GetFiles(ctx context.Context) ([]*model.File, error)
	GetFileByID(ctx context.Context, id string) (*model.File, error)
}

type file struct {
	*repository.Repository
}

func newFileService(repo *repository.Repository) fileService {
	return &file{
		Repository: repo,
	}
}

func (d *file) CreateFile(ctx context.Context, input model.NewFile) (*model.File, error) {
	file := model.File{
		Filename:    input.File.Filename,
		UserID:      input.UserID,
		ContentType: input.File.ContentType,
		Size:        input.File.Size,
		Category:    input.Category,
	}

	err := os.MkdirAll("./uploads/"+file.UserID, os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = d.DB.Create(&file)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s_%s", file.ID, file.Filename)
	path := fmt.Sprintf("./uploads/%s/%s", file.UserID, url)
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, input.File.File)
	if err != nil {
		return nil, err
	}

	user := model.User{ID: input.UserID}
	user.Files = append(user.Files, file.ID)

	err = d.DB.UpdateOne(user)
	if err != nil {
		return nil, err
	}

	file.Path = path
	file.URL = url
	err = d.DB.UpdateOne(&file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (f *file) UpdateFile(ctx context.Context, data interface{}) (*model.File, error) {
	file := model.File{}

	ub, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ub, &file)
	if err != nil {
		return nil, err
	}

	err = f.DB.UpdateOne(&file)
	if err != nil {
		return nil, err
	}

	err = f.DB.FetchByID(&file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
func (f *file) GetFiles(ctx context.Context) ([]*model.File, error) {
	file := new(model.File)
	files := []model.File{}
	err := f.DB.FetchAll(&files, file)
	if err != nil {
		return nil, err
	}

	fileList := []*model.File{}
	for i := 0; i < len(files); i++{
		fileList = append(fileList, &files[i])
	}

	return fileList, nil
}

func (f *file) GetFileByID(ctx context.Context, id string) (*model.File, error) {
	file := new(model.File)
	file.ID = id
	err := f.DB.FetchByID(file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *file) DownloadFile(ctx context.Context) error {
	c := ctx.Value(apiCtx("apiCtx")).(*gin.Context)

	filename := c.Param("filename")
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// set the default MIME type to send
	mime := http.DetectContentType(buffer)
	fileSize := len(string(buffer))

	reader := bytes.NewReader(buffer)

	// Generate the server headers
	c.Request.Header.Set("Content-Type", mime)
	c.Request.Header.Set("Content-Disposition", "attachment; filename="+filename+"")
	c.Request.Header.Set("Expires", "0")
	c.Request.Header.Set("Content-Transfer-Encoding", "binary")
	c.Request.Header.Set("Content-Length", strconv.Itoa(fileSize))
	c.Request.Header.Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		return err
	}

	return nil
}
