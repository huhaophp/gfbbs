package service

import (
	"bbs/app/model/posts"
	"github.com/gogf/gf/frame/g"
)

// PublishPostsReqEntity
type PublishPostsReqEntity struct {
	Nid     string `p:"nid" v:"required#请选择节点ID"`
	Title   string `p:"title" v:"required#请填写帖子标题"`
	Content string `p:"content" v:"required#请填写帖子内容"`
}

var PostsService = newPostsService()

// newPostsService Initialize the service
func newPostsService() *postsService {
	return &postsService{}
}

// postsService
type postsService struct{}

// Publish Post a post.
func (s *postsService) Publish(publisher int, req *PublishPostsReqEntity) int64 {
	res, err := g.DB().Table(posts.Table).Insert(g.Map{
		"title":       req.Title,
		"nid":         req.Nid,
		"content":     req.Content,
		"uid":         publisher,
		"view_num":    0,
		"comment_num": 0,
	})
	if err != nil {
		g.Log().Error("Posting error:", err)
		return 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		g.Log().Error("Error getting insert ID:", err)
		return 0
	}
	if id <= 0 {
		g.Log().Error("Posting error: LastInsertId is 0.")
		return 0
	}
	return id
}
