package job

import (
	"bbs/app/model/posts"
	"github.com/gogf/gf/frame/g"
	"github.com/uniplaces/carbon"
)

type activeUserJob struct{}

var ActiveUserJob = &activeUserJob{}

var users = make(map[int]int)

const (
	//	配置信息
	TOPIC_WEIGHT = 4 // 话题权重
	REPLY_WEIGHT = 1 // 回复权重
	PASS_DAY     = 17 // 多少天内发表过内容
	USER_NUM     = 6 // 取出来多少用户

	// 缓存相关配置
	CACHE_KEY               = "gfbbs_active_users"
	CACHE_EXPIRE_IN_SECONDS = 65 * 60
)

func (a *activeUserJob) GetActiveUsers() {

}

func (a *activeUserJob) CalculateTopicScore()  {
	g.Dump(carbon.Now().SubDays(PASS_DAY).DateTimeString())
	postUsers, _ := g.DB().Table(posts.Table).
		Fields("uid, COUNT(*) as post_count").
		Where("create_at >", carbon.Now().SubDays(PASS_DAY).DateTimeString()).
		Group("uid").
		All()
	g.Dump(postUsers)
	//for _, postUser := range postUsers {
	//
	//	//_, ok := users[postUser.Map()["uid"]]
	//
	//
	//}
}
