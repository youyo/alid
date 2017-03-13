package alid

import (
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
)

const usEast1 string = "us-east-1"

// Service is wrapping configuration *ec2.EC2
type Service struct {
	*ec2.EC2
}

// NewConfig returns a new Config pointer that can be chained with builder methods to set multiple configuration values inline without using pointers.
//  c := simplebackupec2.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
func NewConfig(region string) *aws.Config {
	return &aws.Config{Region: aws.String(region)}
}

// NewService creates a new instance of the EC2 client with a session.
func NewService(c *aws.Config) (*Service, error) {
	sess, err := session.NewSession(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new session")
	}
	return &Service{ec2.New(sess)}, nil
}

// FetchLatestAmiInfo is request to aws api, and fetch latest ami information.
func (s *Service) FetchLatestAmiInfo() (amiInfo *ec2.Image, err error) {
	params := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String("amazon"),
		},
		Filters: []*ec2.Filter{
			{
				Name: aws.String("image-type"),
				Values: []*string{
					aws.String("machine"),
				},
			},
			{
				Name: aws.String("architecture"),
				Values: []*string{
					aws.String("x86_64"),
				},
			}, {Name: aws.String("root-device-type"), Values: []*string{aws.String("ebs")}}, {Name: aws.String("virtualization-type"), Values: []*string{
				aws.String("hvm"),
			},
			},
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}
	resp, err := s.DescribeImages(params)
	if err != nil {
		return
	}

	var amis []*ec2.Image

	for _, v := range resp.Images {
		if m := strings.HasPrefix(*v.ImageLocation, "amazon/amzn-ami-hvm"); m {
			amis = append(amis, v)
		}
	}
	sort.Slice(amis, func(i, j int) bool {
		return *amis[i].CreationDate > *amis[j].CreationDate
	})
	amiInfo = amis[0]
	return
}

// SelectRegion returns the region to use.
func SelectRegion(optRegion string) (region string, err error) {
	if optRegion == "OS Environment 'AWS_REGION'" {
		if os.Getenv("AWS_REGION") != "" {
			return os.Getenv("AWS_REGION"), err
		}
		return usEast1, err
	}
	return optRegion, err
}
