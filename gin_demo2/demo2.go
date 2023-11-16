package main

import (
	"context"
	"net/http"

	. "github.com/achilsh/zap_log_demo/demo_one"
	"github.com/gin-gonic/gin"
)

//recv other req message.

func main() {

	cfgStr := `{
	"path_file": "logs/test.log",
	"file_max_size_mb": 10,
	"old_file_remain_day": 1,
	"old_file_nums": 2,
	"old_file_compress": false,
	"log_level": "info"
}`
	cfg := ParseCfg(cfgStr)
	if cfg == nil {
		return
	}
	NewSugaredZapLogHandler(cfg)

	r := gin.New()
	r.POST("/test2/push", func(c *gin.Context) {
		if c.Request == nil {
			Infof(context.Background(), "is nil for request.")
			return
		}
		// 模拟 panic and auto stop this process.
		// var vv *int = nil
		// *vv = 1000
		// //
		data := ""
		e := c.ShouldBind(&data)
		if e != nil {
			Infof(context.Background(), "get data fail, e: "+e.Error())
			c.String(http.StatusOK, "is fail")
			return
		}

		Infof(context.Background(), "recv msg body: "+string(data))
		c.String(http.StatusOK, "is succ")
	})

	r.Run(":8081")
}
