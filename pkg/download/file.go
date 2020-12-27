package download

//FileDownload represents file data being downloaded
type FileDownload struct {
	ID                 uint64
	Name               string
	Extension          string
	Path               string
	FromPersonalFolder bool
	Owner              string
}
