package console

func (c UrlBuilder) lambdaBaseUrl() genericUrlBuilder {
	return c.baseUrl().Path("/lambda/home")
}

func (c UrlBuilder) LambdaDashboard() string {
	return c.lambdaBaseUrl().Hash("/functions").String()
}

func (c UrlBuilder) LambdaFunctionCode(functionName string) string {
	return c.lambdaBaseUrl().Hashf("/functions/%s?tab=code", functionName).String()
}
