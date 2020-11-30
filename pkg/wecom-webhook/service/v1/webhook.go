package v1

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/common/model"
	"github.com/rosenlo/toolkits/log"
	"github.com/rosenlo/wecom-webhook/pkg/wecom-webhook/util"
	"github.com/rosenlo/wecom-webhook/utils"
	"github.com/spf13/viper"
)

const (
	FIELD_RECEIVERS = "receivers"
	FIELD_PRIORITY  = "priority"
)

func (s *Service) Webhook(c *gin.Context) {
	data := new(template.Data)
	if err := c.Bind(data); err != nil {
		log.Error(err)
		util.JSONR(c, http.StatusBadRequest, err)
		return
	}
	for _, alert := range data.Alerts {
		switch alert.Labels[FIELD_PRIORITY] {
		case "0":
			log.Info("P0")
			s.sendVoice(alert)
		case "1":
			log.Info("P1")
		default:
			log.Info("no action")
		}
		content := s.buildIMContent(alert)
		if receivers, exists := alert.Annotations[FIELD_RECEIVERS]; exists {
			s.sendWechatUsers(receivers, content)
		}
		s.sendWechatRobot(content)
	}
	util.JSONR(c, "success")
}

func (s *Service) buildIMContent(alert template.Alert) string {
	title := "恢复通知"
	if alert.Status != string(model.AlertResolved) {
		title = fmt.Sprintf("[P%s]告警", alert.Labels[FIELD_PRIORITY])
	}
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n", title))
	for k, v := range alert.Annotations {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	buffer.WriteString(fmt.Sprintf("Starts At: %s\n", alert.StartsAt.Add(time.Hour*8).Format(utils.TIME_FORMAT)))

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
	err := utils.RestRequest("POST", url, body, headers)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) sendWechatUsers(receivers, content string) {
	body := map[string]interface{}{
		"username": strings.Split(receivers, ","),
		"content":  content,
	}
	url := viper.GetString("user_url")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	err := utils.RestRequest("POST", url, body, headers)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) sendVoice(alert template.Alert) {
}
