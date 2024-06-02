package repos

import (
	"context"
	"fmt"
	"io"
	"meeting-center/src/clients"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type gcsRepo struct {
	gcs *clients.GcsInstance
}

type GcsRepo interface {
	UploadFile(ctx context.Context, file multipart.File, filename string) error
	GetSignedURL(ctx context.Context, objectName string) (string, error)
	DeleteFile(ctx context.Context, filename string) error
}

func NewGcsRepo(gcsArgs ...*clients.GcsInstance) GcsRepo {
	if len(gcsArgs) == 0 {
		return gcsRepo{gcs: clients.GetGCSInstance()}
	} else if len(gcsArgs) == 1 {
		return gcsRepo{gcs: gcsArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (gr gcsRepo) UploadFile(ctx context.Context, file multipart.File, objectName string) error {
	bucket := gr.gcs.Bucket()
	obj := bucket.Object(objectName)
	writer := obj.NewWriter(ctx)

	fmt.Println(objectName)

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}

func (gr gcsRepo) GetSignedURL(ctx context.Context, objectName string) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(5 * time.Minute),
	}

	url, err := gr.gcs.Bucket().SignedURL(objectName, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (gr gcsRepo) DeleteFile(ctx context.Context, objectName string) error {
	bucket := gr.gcs.Bucket()
	obj := bucket.Object(objectName)

	return obj.Delete(ctx)
}
