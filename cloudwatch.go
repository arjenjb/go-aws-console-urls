package console

import (
	"net/url"
	"strings"
)

func (c UrlBuilder) CloudwatchHome() string {
	return c.baseUrl().Path("/cloudwatch/home").Hash("dashboards/").String()
}

func (c UrlBuilder) CloudwatchLogGroups() string {
	return c.baseUrl().Path("/cloudwatch/home").Hash("logsV2:log-groups").String()
}

func (c UrlBuilder) CloudwatchLogGroup(s string) string {
	encoded := strings.Replace(url.QueryEscape(s), "%", "$25", -1)
	return c.baseUrl().Path("/cloudwatch/home").Hashf("logsV2:log-groups/log-group/%s", encoded).String()
}

func (c UrlBuilder) CloudwatchMetrics() string {
	return c.baseUrl().Path("/cloudwatch/home").Hash("metricsV2").String()
}
