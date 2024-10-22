package console

func (c UrlBuilder) SsmParameters() string {
	return c.baseUrl().Path("/systems-manager/parameters").String()
}

func (c UrlBuilder) SsmParameter(parameter string) string {
	return c.baseUrl().Pathf("/systems-manager/parameters/%s/description", parameter).String()
}
