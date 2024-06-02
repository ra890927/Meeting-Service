package services

import (
	"context"
	"errors"
	"meeting-center/src/domains"
	"meeting-center/src/models"
	"meeting-center/src/utils"
	"mime/multipart"
	"path/filepath"
	"slices"
)

type fileService struct {
	gcsDomain     domains.GcsDomain
	fileDomain    domains.FileDomain
	meetingDomain domains.MeetingDomain
}

type FileServiceArg struct {
	GcsDomain     domains.GcsDomain
	FileDomain    domains.FileDomain
	MeetingDomain domains.MeetingDomain
}

type FileService interface {
	UploadFile(ctx context.Context, operator models.User, file *models.File, object multipart.File) (string, error)
	GetAllSignedURLsByMeeting(ctx context.Context, operator models.User, meetingID string) ([]models.File, []string, error)
	GetSignedURL(ctx context.Context, operater models.User, id string) (string, error)
	DeleteFile(ctx context.Context, operater models.User, id string) error
}

func NewFileService(fileServiceArgs ...FileServiceArg) FileService {
	if len(fileServiceArgs) == 0 {
		return fileService{
			gcsDomain:     domains.NewGcsDomain(),
			fileDomain:    domains.NewFileDomain(),
			meetingDomain: domains.NewMeetingDomain(),
		}
	} else if len(fileServiceArgs) == 1 {
		return fileService{
			gcsDomain:     fileServiceArgs[0].GcsDomain,
			fileDomain:    fileServiceArgs[0].FileDomain,
			meetingDomain: fileServiceArgs[0].MeetingDomain,
		}
	} else {
		panic("too many arguments")
	}
}

func (fs fileService) UploadFile(ctx context.Context, operator models.User, file *models.File, object multipart.File) (string, error) {
	permission, err := fs.getPermissionByFile(operator, *file)
	if err != nil {
		return "", err
	}

	if !utils.CheckPermission(permission, utils.Upload) {
		return "", errors.New("only admin and participant can upload file")
	}

	file.UploaderID = operator.ID
	// if success, it will get file id
	if err := fs.fileDomain.UploadFile(file); err != nil {
		return "", err
	}

	ext := filepath.Ext(file.FileName)
	if err := fs.gcsDomain.UploadFile(ctx, object, file.ID, ext); err != nil {
		return "", err
	}

	return fs.gcsDomain.GetSignedURL(ctx, file.ID, file.FileExt)
}

func (fs fileService) GetAllSignedURLsByMeeting(ctx context.Context, operator models.User, meetingID string) ([]models.File, []string, error) {
	permission, err := fs.getPermissionByMeetingID(operator, meetingID)
	if err != nil {
		return []models.File{}, []string{}, err
	}
	if !utils.CheckPermission(permission, utils.Read) {
		return []models.File{}, []string{}, errors.New("only admin and participant can download file")
	}

	files, err := fs.fileDomain.GetAllFilesByMeetingID(meetingID)
	if err != nil {
		return []models.File{}, []string{}, err
	}

	urls := make([]string, 0)
	for _, file := range files {
		url, err := fs.gcsDomain.GetSignedURL(ctx, file.ID, file.FileExt)
		if err != nil {
			return []models.File{}, []string{}, err
		}
		urls = append(urls, url)
	}

	return files, urls, nil
}

func (fs fileService) GetSignedURL(ctx context.Context, operater models.User, id string) (string, error) {
	file, err := fs.fileDomain.GetFile(id)
	if err != nil {
		return "", err
	}

	permission, err := fs.getPermissionByFile(operater, file)
	if err != nil {
		return "", err
	}

	if !utils.CheckPermission(permission, utils.Read) {
		return "", errors.New("only admin and participant can download file")
	}

	return fs.gcsDomain.GetSignedURL(ctx, file.ID, file.FileExt)
}

func (fs fileService) DeleteFile(ctx context.Context, operater models.User, id string) error {
	file, err := fs.fileDomain.GetFile(id)
	if err != nil {
		return err
	}

	permission, err := fs.getPermissionByFile(operater, file)
	if err != nil {
		return err
	}

	if !utils.CheckPermission(permission, utils.Delete) {
		return errors.New("only admin, organizer, and uploader can delete file")
	}

	if err := fs.gcsDomain.DeleteFile(ctx, file.ID, file.FileExt); err != nil {
		return err
	}

	if err := fs.fileDomain.DeleteFile(file.ID); err != nil {
		return err
	}

	return nil
}

func (fs fileService) getPermissionByFile(operator models.User, file models.File) (utils.Permission, error) {
	permission := utils.Empty

	meeting, err := fs.meetingDomain.GetMeeting(file.MeetingID)
	if err != nil {
		return utils.Empty, nil
	}

	if operator.Role == "admin" || operator.ID == file.UploaderID || operator.ID == meeting.OrganizerID {
		permission = utils.Upload | utils.Delete | utils.Read
	} else if slices.Contains(meeting.Participants, operator.ID) {
		permission = utils.Upload | utils.Read
	}

	return permission, nil
}

func (fs fileService) getPermissionByMeetingID(operator models.User, meetingID string) (utils.Permission, error) {
	permission := utils.Empty

	meeting, err := fs.meetingDomain.GetMeeting(meetingID)
	if err != nil {
		return utils.Empty, nil
	}

	if operator.Role == "admin" || operator.ID == meeting.OrganizerID {
		permission = utils.Upload | utils.Delete | utils.Read
	} else if slices.Contains(meeting.Participants, operator.ID) {
		permission = utils.Upload | utils.Read
	}

	return permission, nil
}
