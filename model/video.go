package model

import (
	"github.com/jinzhu/gorm"
	"kilikili/cache"
	"strconv"
)

type Video struct {
	gorm.Model
	User          User   ` gorm:"ForeignKey:Uid"`
	Uid           uint   `gorm:"not null"`       //用户id
	Title         string `gorm:"index;not null"` //标题
	Uptime        int64  `gorm:"autoCreateTime"` //上传时间
	Content       string `gorm:"not null"`       //视频
	IsPass        int64  `gorm:"default:0"`      //是否通过审核
	CommentNum    int64  `gorm:"default:0"`      //评论数
	CollectionNum int64  `gorm:"default:0"`      //收藏数
	LikeNum       int64  `gorm:"default:0"`      //点赞数
	ForwardNum    int64  `gorm:"default:0"`      //转发数
}

func (video *Video) View() uint64 {
	//增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(video.ID)).Result()
	//string 转 int
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (video *Video) AddView() {
	//Redis Incr 命令将 key 中储存的数字值增一。
	//如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
	cache.RedisClient.Incr(cache.TaskViewKey(video.ID)) //增加视频点击数
	//Redis
	//Zincrby 命令对有序集合中指定成员的分数加上增量 increment
	//可以通过传递一个负数值 increment ，让分数减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
	//当 key 不存在，或分数不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
	//zadd 命令
	//添加元素到集合，元素在集合中存在则更新对应score
	//zadd key score member
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(video.ID))) //增加排行点击数
}
