# Context

https://github.com/backendcloud/example/tree/master/golang/go-routine-context


# Context控制多个goroutine

使用Context控制一个goroutine的例子如上，非常简单，下面我们看看控制多个goroutine的例子，其实也比较简单。

示例中启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，当我们使用cancel函数通知取消时，这3个goroutine都会被结束。这就是Context的控制能力，它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，这时就可以进行清理操作了，最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题。
