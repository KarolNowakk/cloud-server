package permissions

import "cloud/pkg/upload/uploadpb"

//UploadPermissions is interface returned by upload permissions
type UploadPermissions interface {
	//CanDownload checks if particular user can download requested file or folder
	CanUploadToFolder(userID string, req *uploadpb.FileUploadRequest) error
}

// //Repository is interface thath provided repository has to implement
// type Repository interface{}

//NewUploadPermissions returns sruct implementing UploadPermissions interface
func NewUploadPermissions() UploadPermissions {
	return &uploadPermissions{}
}

type uploadPermissions struct{}

func (p *uploadPermissions) CanUploadToFolder(userID string, req *uploadpb.FileUploadRequest) error {
	if req.GetInfo().ToPersonalFolder {
		return nil
	}

	return ErrPermissionDenied
}
