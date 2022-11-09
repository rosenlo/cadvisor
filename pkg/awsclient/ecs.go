package awsclient

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

func (c *Client) DescribeECSTask(cluster *string, task string) (*types.Task, error) {
	input := &ecs.DescribeTasksInput{
		Cluster: cluster,
		Tasks:   []string{task},
	}
	result, err := c.ecs.DescribeTasks(context.TODO(), input)
	if err != nil {
		log.Printf("failed to retrieving about ECS Task: %s", err)
		return nil, err
	}

	if len(result.Tasks) == 0 {
		return nil, errors.New("record not found")
	}

	return &result.Tasks[0], nil

}
