package console

func (c UrlBuilder) CodebuildDashboard() string {
	return c.baseUrl().Pathf("/codesuite/codebuild/projects").String()
}
func (c UrlBuilder) CodebuildProjectHistory(project string) string {
	return c.baseUrl().Pathf("/codesuite/codebuild/projects/%s/history", project).String()
}
