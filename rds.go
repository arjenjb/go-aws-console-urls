package console

func (c UrlBuilder) RDSDashboard() string {
	return c.baseUrl().Path("/rds/home").String()
}

func (c UrlBuilder) RDSDbInstance(instanceName string) string {
	return c.baseUrl().Path("/rds/home").Hashf("database:id=%s;is-cluster=false;tab=connectivity", instanceName).String()
}

func (c UrlBuilder) RDSDbCluster(clusterName string) string {
	return c.baseUrl().Path("/rds/home").Hashf("database:id=%s;is-cluster=true;tab=connectivity", clusterName).String()
}
