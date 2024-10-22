package console

func (c UrlBuilder) CloudfrontDistributions() string {
	return c.globalBaseUrl().Path("/cloudfront/v3/home").Hash("/distributions").String()
}

func (c UrlBuilder) CloudfrontDistribution(distribution string) string {
	return c.globalBaseUrl().Path("/cloudfront/v3/home").Hashf("/distributions/%s", distribution).String()
}
