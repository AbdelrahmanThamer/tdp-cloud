package monitor

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/monitor"

	"github.com/gin-gonic/gin"
)

// 获取监控数据

func getMonitorData(c *gin.Context) {

	var ud = midware.GetUserdata(c)
	var rq monitor.GetMonitorDataRequestParams

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	response, err := monitor.GetMonitorData(ud, &rq)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}