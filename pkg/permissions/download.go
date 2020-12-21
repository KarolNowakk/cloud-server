package permissions

import (
	"cloud/pkg/download/downloadpb"
)

//DownladPermissions is interface returned by downloadpermissions
type DownladPermissions interface {
	//CanDownload checks if particular user can download requested file or folder
	CanDownload(userID string, req *downloadpb.FileDownloadRequest) error
}

// //Repository is interface thath provided repository has to implement
// type Repository interface{}

//NewDownloadPermissions returns sruct implementing DownladPermissions interface
func NewDownloadPermissions() DownladPermissions {
	return &downloadPermissions{}
}

type downloadPermissions struct{}

func (p *downloadPermissions) CanDownload(userID string, req *downloadpb.FileDownloadRequest) error {
	if req.FromPersonalFolder {
		return nil
	}

	return ErrPermissionDenied
}
