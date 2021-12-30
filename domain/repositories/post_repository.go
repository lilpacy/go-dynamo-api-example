package repositories

import "github.com/lilpacy/gin-nginx-next-docker-min/domain"

type PostRepository interface {
	Store(domain.Post) (int, error)
	FindAll() ([]domain.Post, error)
	Find(id string) (domain.Post, error)
}
