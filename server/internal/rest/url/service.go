package url

import "net/http"

type Service interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) Get(w http.ResponseWriter, r *http.Request) {
	// s.repo.Get()
}
