package antspool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

func InitPool() *ants.Pool {
	AntsPool, err := ants.NewPool(100000)
	fmt.Println("pool 的值是---->", AntsPool)
	if err != nil {
		panic("进程池连接失败")
	}
	return AntsPool
}
