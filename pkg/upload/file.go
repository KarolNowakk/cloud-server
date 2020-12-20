package upload

//File represents file being downloaded
type File struct {
	ID        uint64
	Name      string
	Extension string
	FullPath  string
}
