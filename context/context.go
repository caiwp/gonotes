package context

import (
	"context"
	"fmt"
	"time"
)

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了")
				return

			default:
				fmt.Println("goroutine 监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

func run1() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "监控一")
	go watch(ctx, "监控二")
	go watch(ctx, "监控三")

	time.Sleep(10 * time.Second)

	fmt.Println("通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " 监控退出")
			return
		default:
			fmt.Println(name, " goroutine 监控中")
			time.Sleep(2 * time.Second)
		}
	}
}

func run2() {
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	go watch(ctx, "监控一")
	go watch(ctx, "监控二")

	time.Sleep(10 * time.Second)

	fmt.Println("通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}
