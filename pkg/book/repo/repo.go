package repo

import "sync"

type Repo struct {
	Mtx sync.Mutex
}

func NewRepo() *Repo {
	return &Repo{}
}
