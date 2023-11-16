package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	. "github.com/achilsh/zap_log_demo/demo_one"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

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
	if cfg ==nil {
		return 
	}
	NewSugaredZapLogHandler(cfg)

	var r = gin.New()	

	r.GET("test_demo1", func(c *gin.Context) {
		resp, _ := CallPeerServer()
		
		c.String(http.StatusOK, resp)
	})

	r.Run(strings.Join([]string{"", "8080"}, ":"))
}


func CallPeerServer() (r string, e error) {
	peerIp := os.Getenv("req_ip")
	if len(peerIp) <= 0 {
		Infof(context.Background(), "set default ip")
		peerIp = "172.18.0.1"
	}
	cli := resty.New()
	rr, ee := cli.R().SetBody("this request one stop...").Post("http://" + peerIp + ":8081/test2/push")

	if ee != nil {
		Infof(context.Background(), "is fail from peer node, e: " + ee.Error())
		return "is fail from peer", nil
	}


	Infof(context.Background(), "ret: " + rr.String())
	return rr.String(), nil
}