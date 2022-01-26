package upload

//File represents file being downloaded
type File struct {
	ID         uint64
	SearchTags string
	Name       string
	Path       string
	Owner      string
}
