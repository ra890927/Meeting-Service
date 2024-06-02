package domains

import (
	"context"
	"meeting-center/src/repos"
	"mime/multipart"
)

type gcsDoamin struct {
	gcsRepo repos.GcsRepo
}

type GcsDomain interface {
	UploadFile(ctx context.Context, file multipart.File, objectID string, objExt string) error
	GetSignedURL(ctx context.Context, objectID string, objExt string) (string, error)
	DeleteFile(ctx context.Context, objectID string, objExt string) error
}

func NewGcsDomain(gcsRepoArgs ...repos.GcsRepo) GcsDomain {
	if len(gcsRepoArgs) == 0 {
		return gcsDoamin{gcsRepo: repos.NewGcsRepo()}
	} else if len(gcsRepoArgs) == 1 {
		return gcsDoamin{gcsRepo: gcsRepoArgs[0]}
	} else {
		panic("Too many arguments")
	}
}

func (gd gcsDoamin) UploadFile(ctx context.Context, file multipart.File, objectID string, objExt string) error {
	objectName := objectID + objExt
	return gd.gcsRepo.UploadFile(ctx, file, objectName)
}

func (gd gcsDoamin) GetSignedURL(ctx context.Context, objectID string, objExt string) (string, error) {
	objectName := objectID + objExt
	return gd.gcsRepo.GetSignedURL(ctx, objectName)
}

func (gd gcsDoamin) DeleteFile(ctx context.Context, objectID string, objExt string) error {
	objectName := objectID + objExt
	return gd.gcsRepo.DeleteFile(ctx, objectName)
}
