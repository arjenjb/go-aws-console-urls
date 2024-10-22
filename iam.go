package console

import arn2 "github.com/aws/aws-sdk-go-v2/aws/arn"

func (c UrlBuilder) IAMDashboard() string {
	return c.globalBaseUrl().Path("/iamv2/home").Hash("/home").String()
}

func (c UrlBuilder) IAMRoles() string {
	return c.globalBaseUrl().Path("/iamv2/home").Hash("/roles").String()
}

func (c UrlBuilder) IAMPolicies() string {
	return c.globalBaseUrl().Path("/iamv2/home").Hash("/policies").String()
}

func (c UrlBuilder) IAMUsers() string {
	return c.globalBaseUrl().Path("/iamv2/home").Hash("/users").String()
}

func (c UrlBuilder) IAMRole(roleName string) string {
	return c.globalBaseUrl().Path("/iamv2/home").Hashf("/roles/details/%s", roleName).String()
}

func (c UrlBuilder) IAMPolicy(policyArn arn2.ARN) string {
	return c.globalBaseUrl().Path("/iamv2/home").Hashf("/policies/details/%s", policyArn.String()).String()
}

func (c UrlBuilder) IAMUser(userName string) string {
	return c.globalBaseUrl().Path("/iamv2/home").Hashf("/users/details/%s", userName).String()
}
