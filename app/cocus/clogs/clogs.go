package clogs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func Init(region, logGroupName string) (*cloudwatchlogs.CloudWatchLogs, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	if err != nil {
		return nil, err
	}

	cwl := cloudwatchlogs.New(sess)

	err = ensureLogGroupExists(logGroupName, cwl)
	if err != nil {
		return nil, err
	}

	return cwl, nil
}

func ensureLogGroupExists(name string, cwl *cloudwatchlogs.CloudWatchLogs) error {
	resp, err := cwl.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		return err
	}

	for _, logGroup := range resp.LogGroups {
		if *logGroup.LogGroupName == name {
			return nil
		}
	}

	_, err = cwl.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &name,
	})
	if err != nil {
		return err
	}

	_, err = cwl.PutRetentionPolicy(&cloudwatchlogs.PutRetentionPolicyInput{
		RetentionInDays: aws.Int64(14),
		LogGroupName:    &name,
	})

	return err
}
