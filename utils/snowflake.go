package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// @file      : snowflake.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

var node *sf.Node

// SnowFlakeInit 基于雪花算法生成用户ID
func SnowFlakeInit(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenSnowFlakeID() int64 {
	return node.Generate().Int64()
}
