package random

/*
分布式 ID 生成器

ref:
	- https://zhuanlan.zhihu.com/p/38308576
	- Twitter Snowflake: 雪花算法
		- 有序+递增
		- 纯数字
		- https://github.com/beinan/fastid
		- https://github.com/bwmarrin/snowflake
	- https://github.com/edwingeng/wuid
*/

import (
	"github.com/beinan/fastid"
)

// 18 bit 纯数字 ID: 有序+递增+唯一
func SnowFlakeID() uint64 {
	// 18bit
	uid := fastid.CommonConfig.GenInt64ID()
	return uint64(uid)
}
