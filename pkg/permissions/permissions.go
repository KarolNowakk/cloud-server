package permissions

//NewDownloadPermissions returns sruct implementing DownladPermissions interface
func NewPermissions() Permissions {
	return Permissions{}
}

type Permissions struct{}

func (Permissions) CanDownload(userID, ownerID string) error {
	if userID != ownerID {
		return ErrPermissionDenied
	}

	return nil
}

func (Permissions) CanDelete(userID, ownerID string) error {
	if userID != ownerID {
		return ErrPermissionDenied
	}

	return nil
}
