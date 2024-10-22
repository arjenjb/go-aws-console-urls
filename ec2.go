package console

import "fmt"

func (c UrlBuilder) EC2Dashboard() string {
	return c.baseUrl().Path("/ec2/home").Hash("Home:").String()
}

func (c UrlBuilder) EC2Instances() string {
	return c.baseUrl().Path("/ec2/home").Hash("Instances:").String()
}

func (c UrlBuilder) EC2LoadBalancers() string {
	return c.baseUrl().Path("/ec2/home").Hash("LoadBalancers:").String()
}

func (c UrlBuilder) EC2TargetGroups() string {
	return c.baseUrl().Path("/ec2/home").Hash("TargetGroups:").String()
}

func (c UrlBuilder) EC2SecurityGroups() string {
	return c.baseUrl().Path("/ec2/home").Hash("SecurityGroups:").String()
}

func (c UrlBuilder) EC2Instance(instanceID string) string {
	return c.baseUrl().Path("/ec2/home").Hash(fmt.Sprintf("InstanceDetails:instanceId=%s", instanceID)).String()
}

func (c UrlBuilder) EC2SecurityGroup(securityGroupID string) string {
	return c.baseUrl().Path("/ec2/home").Hash(fmt.Sprintf("SecurityGroup:groupId=%s", securityGroupID)).String()
}

func (c UrlBuilder) EC2LoadBalancer(loadBalancerARN string) string {
	return c.baseUrl().Path("/ec2/home").Hash(fmt.Sprintf("LoadBalancer:loadBalancerArn=%s", loadBalancerARN)).String()
}
