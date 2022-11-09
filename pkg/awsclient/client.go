package awsclient

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	Region     string = "ap-northeast-1"
	MaxResults int32  = 100
)

func MustConfig(assumeRole *AssumeRole) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(Region),
	)
	if err != nil {
		log.Panicf("failed to load default configuration, %v", err)
	}
	if assumeRole != nil {
		creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(cfg), assumeRole.Arn)
		cfg.Region = assumeRole.Region
		cfg.Credentials = aws.NewCredentialsCache(creds)
		cred, err := creds.Retrieve(context.Background())
		if err != nil {
			log.Panicf("failed to Retrieve credentials %s", err)
		}
		log.Printf("assume role: %s region: %s expires: %s", assumeRole.Arn, assumeRole.Region, cred.Expires)
	}
	return cfg
}

type Client struct {
	Role *AssumeRole
	cfg  aws.Config
	ecs  *ecs.Client
}

type AssumeRole struct {
	Arn     string
	Region  string
	Account string
}

func MustNew(assumeRole *AssumeRole) *Client {
	cfg := MustConfig(assumeRole)
	return &Client{
		Role: assumeRole,
		cfg:  cfg,
		ecs:  ecs.NewFromConfig(cfg),
	}
}
