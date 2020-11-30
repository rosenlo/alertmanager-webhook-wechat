module github.com/rosenlo/wecom-webhook

go 1.13

replace github.com/Sirupsen/logrus v1.6.0 => github.com/sirupsen/logrus v1.6.0

require (
	github.com/Sirupsen/logrus v1.6.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/prometheus/alertmanager v0.21.0
	github.com/prometheus/common v0.10.0
	github.com/rosenlo/toolkits v0.0.0-20201130030658-d7eb83aaddae
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)
