package console

import (
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EcsTaskArn(t *testing.T) {
	a, err := arn.Parse("arn:aws:ecs:eu-west-1:123456890:task/MyCluster/223340300f34487f9fc1bff3d6afcddd")
	assert.NoError(t, err)

	c := NewConsoleUrlBuilder("eu-west-1")
	s, err := c.ARN(a)
	assert.NoError(t, err)
	assert.Equal(t, "https://eu-west-1.console.aws.amazon.com/ecs/v2/clusters/MyCluster/tasks/223340300f34487f9fc1bff3d6afcddd/configuration?region=eu-west-1", s)
}
