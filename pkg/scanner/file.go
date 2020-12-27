package scanner

//ScannedFile is struct that scanerr service is using
type ScannedFile struct {
	ID        string
	Path      string
	Owner     string
	Name      string
	Extension string
}
