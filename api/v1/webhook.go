package v1

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/common/model"
	"github.com/rosenlo/alertmanager-webhook-wechat/util"
	"github.com/spf13/viper"
)

type LabelField string

const (
	Priority LabelField = "priority"
	Describe LabelField = "describe"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func (s *Service) Webhook(c *gin.Context) {
	data := new(template.Data)
	if err := c.Bind(data); err != nil {
		log.Printf("[error] %v", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	for _, alert := range data.Alerts {
		s.sendWechatRobot(s.buildIMContent(alert))
	}
	c.String(http.StatusOK, "success")
}

func (s *Service) buildIMContent(alert template.Alert) string {
	title := "恢复通知"
	if alert.Status != string(model.AlertResolved) {
		title = fmt.Sprintf("[P%s]告警", alert.Labels[string(Priority)])
	}
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n", title))
	for k, v := range alert.Annotations {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	buffer.WriteString(fmt.Sprintf("Starts At: %s\n", alert.StartsAt.Add(time.Hour*8).Format(TimeFormat)))

	return buffer.String()
}

func (s *Service) sendWechatRobot(content string) {
	body := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": content,
		},
	}
	url := viper.GetString("robot_url")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	respBody, err := util.RestRequest("POST", url, body, headers)
	if err != nil {
		log.Printf("[error] %v", err)
		return
	}
	log.Printf("[debug] response: %s", string(respBody))
}
