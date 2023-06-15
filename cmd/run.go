package cmd

var (
	// 需要调用mdels的InitData的方法初始化数据的时候传 true
	initData = GetEnvDefaultBool("INIT", false)

	// gin-mode 发版时需要传release
	mode = GetEnvDefaultString("MODE", "debug")
)

func Execute() {
	//添加prometheus监控
	// recordMetrics()
	//启动服务
	run()
}

// func recordMetrics() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total",
// 		Help: "The total number of processed events",
// 	})
// )
