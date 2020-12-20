package permissions

import "strings"

//DownladPermissions is interface returned by downloadpermissions
type DownladPermissions interface {
	//CanDownload checks if particular user can download requested file or folder
	CanDownload(userID, path string) error
}

//Repository is interface thath provided repository has to implement
type Repository interface{}

//NewDownloadPermissions returns sruct implementing DownladPermissions interface
func NewDownloadPermissions() DownladPermissions {
	return &downloadPermissions{}
}

type downloadPermissions struct{}

func (p *downloadPermissions) CanDownload(userID, path string) error {
	pathSlice := strings.Split(path, "/")
	if len(path) < 1 {
		return ErrInvalidPath
	}

	if pathSlice[0] != userID {
		return ErrPermissionDenied
	}

	return nil
}
