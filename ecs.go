package console

func (c UrlBuilder) EcsDashboard() string {
	return c.baseUrl().Path("/ecs/home").String()
}

func (c UrlBuilder) EcsService(cluster string, service string) string {
	return c.baseUrl().Pathf("/ecs/v2/clusters/%s/services/%s/health", cluster, service).String()
}

func (c UrlBuilder) EcsCluster(cluster string) string {
	return c.baseUrl().Pathf("/ecs/v2/clusters/%s/services", cluster).String()
}

func (c UrlBuilder) EcsTask(cluster string, task string) string {
	return c.baseUrl().Pathf("/ecs/v2/clusters/%s/tasks/%s/configuration", cluster, task).String()
}
