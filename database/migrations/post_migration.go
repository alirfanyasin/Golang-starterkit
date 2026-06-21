package migrations

import "github.com/alirfanyasin/golang-starterkit/packages/post"

// GetPostModels returns models related to the Post feature for migration
func GetPostModels() []interface{} {
	return []interface{}{
		&post.Post{},
	}
}
