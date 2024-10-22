package console

func (c UrlBuilder) EcrDashboard() string {
	return c.baseUrl().Path("/ecr/repositories").String()
}

func (c UrlBuilder) EcrRepository(registryId, name string) string {
	return c.baseUrl().Pathf("/ecr/repositories/private/%s/%s", registryId, name).String()
}
