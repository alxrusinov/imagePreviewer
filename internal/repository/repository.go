package repository

type Repo interface {
	Save() error
	Get() error
}
