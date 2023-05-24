package poststore

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	posts       = "posts/%s/%s"
	postsLabels = "posts/%s/%s/%s"
	all         = "posts"
)

func generateKey(string) string {
	id := uuid.New().String()

	return fmt.Sprintf(posts, id)

}

func constructKey(id string) string {

	return fmt.Sprintf(posts, id)

}
