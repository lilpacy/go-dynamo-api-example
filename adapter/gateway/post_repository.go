package gateway

import (
	"fmt"
	"github.com/guregu/dynamo"
	"github.com/lilpacy/gin-nginx-next-docker-min/domain"
	"golang.org/x/xerrors"
)

type (
	PostRepository struct {
		Table dynamo.Table
	}
)

func (r *PostRepository) Store(post domain.Post) (id int, err error) {
	err = r.Table.Put(post).Run()
	if err != nil {
		e := xerrors.Errorf("error in store method: %v", err)
		fmt.Printf("Failed to get post[%+v]\n", e)
		return
	}
	return int(0), err
}

func (r *PostRepository) FindAll() (posts []domain.Post, err error) {
	if err = r.Table.Scan().All(&posts); err != nil {
		e := xerrors.Errorf("error in store method: %v", err)
		fmt.Printf("Failed to get post[%+v]\n", e)
	}
	return posts, err
}

func (r *PostRepository) Find(id string) (post domain.Post, err error) {
	err = r.Table.Get("Id", id).One(&post)
	return post, err
}
