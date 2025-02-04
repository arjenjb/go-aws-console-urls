package console

import (
	"fmt"
	arn2 "github.com/aws/aws-sdk-go-v2/aws/arn"
	"strings"
)

type UrlBuilder struct {
	region string
}

func NewUrlBuilder(region string) UrlBuilder {
	return UrlBuilder{region: region}
}

func (c UrlBuilder) globalBaseUrl() genericUrlBuilder {
	return c.constructBaseUrl("us-east-1", c.region)
}
func (c UrlBuilder) baseUrl() genericUrlBuilder {
	return c.constructBaseUrl(c.region, c.region)
}

func (c UrlBuilder) constructBaseUrl(serviceRegion, userRegion string) genericUrlBuilder {
	return genericUrlBuilder{
		host: serviceRegion + ".console.aws.amazon.com",
	}.Query("region", userRegion)
}

func (c UrlBuilder) ARN(target arn2.ARN) (string, error) {
	parts := strings.Split(target.Resource, "/")
	head := parts[0]
	tail := ""
	last := parts[len(parts)-1]

	if len(parts) > 1 {
		tail = target.Resource[len(parts)+1:]
	}

	switch target.Service {
	case "iam":
		switch head {
		case "role":
			return c.IAMRole(tail), nil
		case "policy":
			return c.IAMPolicy(target), nil
		case "user":
			return c.IAMUser(last), nil
		}

	case "ecs":
		switch head {
		case "service":
			return c.EcsService(parts[1], parts[2]), nil
		case "cluster":
			return c.EcsCluster(tail), nil
		case "task":
			return c.EcsTask(parts[1], parts[2]), nil
		}

	case "ecr":
		switch head {
		case "repository":
			return c.EcrRepository(target.AccountID, tail), nil
		}

	case "lambda":
		if strings.HasPrefix(target.Resource, "function:") {
			functionName := strings.TrimPrefix(target.Resource, "function:")
			return c.LambdaFunctionCode(functionName), nil
		}

	case "logs":
		parts := strings.Split(target.Resource, ":")
		if parts[0] == "log-group" && len(parts) == 3 && parts[2] == "*" {
			return c.CloudwatchLogGroup(parts[1]), nil
		}

	case "rds":
		parts := strings.Split(target.Resource, ":")
		if len(parts) == 2 && parts[0] == "db" {
			return c.RDSDbInstance(parts[1]), nil
		} else if len(parts) == 2 && parts[0] == "cluster" {
			return c.RDSDbCluster(parts[1]), nil
		}

	case "s3":
		parts := strings.Split(target.Resource, ":")
		if len(parts) == 1 {
			return c.S3Bucket(parts[0]), nil
		} else if len(parts) == 1 {
			return c.S3Object(parts[0], parts[1]), nil
		}
	}

	return "", fmt.Errorf("cannot figure out console url for this ARN %s", target)
}

func (c UrlBuilder) HomeUrl() string {
	return c.baseUrl().String()
}
