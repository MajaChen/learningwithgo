package golang

import(
"context"
"fmt"
"math/rand"
"sync"
"time"
)

type stringType string
const num stringType = "luckynum"
var lock sync.Mutex

func hitBall() {

	rand.Seed(time.Now().Unix())

	ctx, canFunc := context.WithCancel(context.Background())
	defer canFunc()
	luckyNum, isHit := 666, make(chan int)
	defer close(isHit)

	ctx = context.WithValue(ctx, num, luckyNum)
	subCanFuncs := make([]context.CancelFunc, 0, 10)
	for i := 0; i < 10; i++ {
		subCtx, subCanFun := context.WithCancel(ctx)
		subCanFuncs = append(subCanFuncs, subCanFun)
		go func(subCtx context.Context, isHit chan int) {
			// 从context中获取幸运数
			luckyNum := subCtx.Value(num).(int)
			for {
				select {
				case <-subCtx.Done():
					fmt.Printf("the ball has been hitten\n")
					return
				default:
					// 生成随机数,如果命中幸运数，向“爸爸”报告
					if ranndom := rand.Intn(1000); ranndom == luckyNum {
						isHit <- 1
					}
				}
			}
		}(subCtx, isHit)
	}

	<-isHit
	fmt.Print("one of my childern has hit the bool\n")
	for _, subCanFunc := range subCanFuncs {
		subCanFunc()
	}

	// 假设父goroutine后面还有耗时工作
	time.Sleep(time.Second * 10)

	fmt.Println("Done!")
}

func hitBall2() {

	ctx, canFunc := context.WithCancel(context.Background())
	defer canFunc()

	for i := 0; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			ctx = context.WithValue(ctx, i, i)
			fmt.Printf("the addr is %p\n", &ctx)
			fmt.Printf("the key is %v,the val is %v\n", i, ctx.Value(i).(int))
			lock.Unlock()
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("the addr is %p\n", &ctx)
	for i := 0; i < 10; i++ {
		fmt.Printf("the key is %v,the val is %v\n", i, ctx.Value(i).(int))
	}

	fmt.Println("Done!")
}