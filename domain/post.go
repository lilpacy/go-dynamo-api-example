package domain

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"time"
)

// Post is struct for DynamoDB
type Post struct {
	Id       string `json:"id"`
	Date     time.Time `json:"date"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Status   string `json:"status"`
	OgImage  string `json:"og_image"`
	Content  string `json:"content"`
}

var imgRegex = regexp.MustCompile(`http.+.(png|jpg|gif|webp)`)

func New(dto PostDTO) Post {
	imgArr := imgRegex.FindAllStringSubmatch(dto.Content, -1)
	return Post{
		Id:       fmt.Sprintf("%x", sha256.Sum256([]byte(dto.Content))),
		Date:     time.Now(),
		Title:    dto.Title,
		Category: "dummy",
		Status:   dto.Status,
		OgImage:  imgArr[0][0], // first image
		Content:  dto.Content,
	}
}

var mockContent = ``

var MockPost = Post{
	Id:       fmt.Sprintf("%x", sha256.Sum256([]byte(mockContent))),
	Date:     time.Now(),
	Title:    "title",
	Category: "dummy",
	Status:   "publish",
	OgImage:  "https://via.placeholder.com/350x150",
	Content:  mockContent,
}
