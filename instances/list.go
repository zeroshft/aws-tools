package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"net/url"
)

func main() {

	aws_region := "us-east-1"

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(aws_region)})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}
	for i := range resp.Reservations {
		for _, inst := range resp.Reservations[i].Instances {
			Instname := "None"
			for _, k := range inst.Tags {
				if *k.Key == "Name" {
					Instname = url.QueryEscape(*k.Value)
				}
			}
			fmt.Println(Instname, *inst.InstanceId, *inst.Hypervisor, *inst.InstanceType, *inst.PublicDnsName, *inst.State.Name)

		}
	}
}
