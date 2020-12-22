package permissions

import "cloud/pkg/upload/uploadpb"

//UploadPermissions is interface returned by upload permissions
type UploadPermissions interface {
	//CanDownload checks if particular user can download requested file or folder
	CanUploadToFolder(userID string, in *uploadpb.FileUploadRequest) error

	//CanDeleteFile checks if particular user can delete requested file or folder
	CanDeleteFile(userID string, req *uploadpb.FileDeleteRequest) error
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

func (p *uploadPermissions) CanDeleteFile(userID string, req *uploadpb.FileDeleteRequest) error {
	//Its empty for now but maybe it will be needed in the future
	return nil
}
