package console

func (c UrlBuilder) Route53Dashboard() string {
	return c.baseUrl().Path("/route53/v2/home").Hash("Dashboard").String()
}

func (c UrlBuilder) Route53HostedZone(zoneId string) string {
	return c.baseUrl().Path("/route53/v2/hostedzones").Hashf("ListRecordSets/%s", zoneId).String()
}
