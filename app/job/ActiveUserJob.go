package job

import (
	"bbs/app/funcs/common"
	"bbs/app/model/comments"
	"bbs/app/model/posts"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/uniplaces/carbon"
	"reflect"
	"time"
)

type activeUserJob struct{}

var ActiveUserJob = &activeUserJob{}

const (
	//	配置信息
	TOPIC_WEIGHT = 4  // 话题权重
	REPLY_WEIGHT = 1  // 回复权重
	PASS_DAY     = 17 // 多少天内发表过内容
	USER_NUM     = 8  // 取出来多少用户

	// 缓存相关配置
	CACHE_KEY              = "gfbbs_active_users"
	CACHE_EXPIRE_IN_MINUTE = 60
)

func (a *activeUserJob) GetActiveUsers() interface{} {
	b, _ := gcache.GetOrSetFunc(CACHE_KEY, func() (interface{}, error) {
		return ActiveUserJob.CalculateActiveUsers(), nil
	}, CACHE_EXPIRE_IN_MINUTE*time.Second)

	return b
}

func (a *activeUserJob) CalculateActiveUsers() []int {
	users := make(map[int]int)

	ActiveUserJob.CalculateTopicScore(users)
	ActiveUserJob.CalculateReplyScore(users)

	usersSort := common.SortMap(users)
	if len(usersSort) > USER_NUM {
		usersSort = usersSort[:USER_NUM]
	}

	return usersSort
}

func (a *activeUserJob) CalculateTopicScore(users map[int]int) {
	postUsers, _ := g.DB().Table(posts.Table).
		Fields("uid, COUNT(*) as post_count").
		Where("create_at >", carbon.Now().SubDays(PASS_DAY).DateTimeString()).
		Group("uid").
		All()
	for _, postUser := range postUsers {
		uid := int(reflect.ValueOf(postUser.Map()["uid"]).Int())
		postCount := int(reflect.ValueOf(postUser.Map()["post_count"]).Int())
		_, ok := users[uid]
		if !ok {
			users[uid] = postCount * TOPIC_WEIGHT
		} else {
			users[uid] += postCount * TOPIC_WEIGHT
		}
	}
}

func (a *activeUserJob) CalculateReplyScore(users map[int]int) {
	commentUsers, _ := g.DB().Table(comments.Table).
		Fields("uid, COUNT(*) as comment_count").
		Where("create_at >", carbon.Now().SubDays(PASS_DAY).DateTimeString()).
		Group("uid").
		All()
	for _, commentUser := range commentUsers {
		uid := int(reflect.ValueOf(commentUser.Map()["uid"]).Int())
		commentCount := int(reflect.ValueOf(commentUser.Map()["comment_count"]).Int())
		_, ok := users[uid]
		if !ok {
			users[uid] = commentCount * REPLY_WEIGHT
		} else {
			users[uid] += commentCount * REPLY_WEIGHT
		}
	}
}
