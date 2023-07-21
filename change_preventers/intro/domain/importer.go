package domain

type Importer interface {
	Invoke(courseId string) (*Table, error)
}
