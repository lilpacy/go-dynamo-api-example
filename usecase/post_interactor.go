package usecase

import (
	"github.com/lilpacy/gin-nginx-next-docker-min/domain"
	"github.com/lilpacy/gin-nginx-next-docker-min/domain/repositories"
)

type PostInteractor struct {
	PostRepository repositories.PostRepository
}

func (i *PostInteractor) Add(post domain.Post) (int, error) {
	return i.PostRepository.Store(post)
}

func (i *PostInteractor) FindAll() ([]domain.Post, error){
	return i.PostRepository.FindAll()
}

func (i *PostInteractor) Find(id string) (domain.Post, error){
	return i.PostRepository.Find(id)
}
