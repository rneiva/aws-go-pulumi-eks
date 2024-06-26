package networking

import (
	"log"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func VpcModule(ctx *pulumi.Context) error {
	cfg := config.New(ctx, "")

	vpcName := cfg.Require("vpcName")
	cidrBlock := cfg.Require("cidrBlock")
	publicSubnetCidrBlock := cfg.Require("publicSubnetCidrBlock")
	privateSubnetCidrBlock := cfg.Require("privateSubnetCidrBlock")

	vpc, err := ec2.NewVpc(ctx, vpcName, &ec2.VpcArgs{
		CidrBlock:       pulumi.String(cidrBlock),
		InstanceTenancy: pulumi.String("default"),
	})
	if err != nil {
		log.Printf("Error creating VPC: %s", err.Error())
		return err
	}

	publicSubnet, err := ec2.NewSubnet(ctx, "publicSubnet", &ec2.SubnetArgs{
		VpcId:     vpc.ID(),
		CidrBlock: pulumi.String(publicSubnetCidrBlock),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("Public"),
		},
	})
	if err != nil {
		log.Printf("Error creating Public Subnet: %s", err.Error())
		return err
	}

	privateSubnet, err := ec2.NewSubnet(ctx, "privateSubnet", &ec2.SubnetArgs{
		VpcId:     vpc.ID(),
		CidrBlock: pulumi.String(privateSubnetCidrBlock),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("Private"),
		},
	})
	if err != nil {
		log.Printf("Error creating Private Subnet: %s", err.Error())
		return err
	}

	ctx.Export("vpcId", vpc.ID())
	ctx.Export("publicSubnetId", publicSubnet.ID())
	ctx.Export("privateSubnetId", privateSubnet.ID())

	return nil
}

func GetVpcId() pulumi.StringOutput {
	return GetVpcId()
}