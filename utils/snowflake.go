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

func SnowFlakeInit(startTime string, machindID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machindID)
	return
}

func GenSnowFlakeID() int64 {
	return node.Generate().Int64()
}
