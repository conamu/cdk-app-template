package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func GetDynamoDb(stack constructs.Construct, id string) awsdynamodb.Table {

	customerTable := awsdynamodb.NewTable(stack, jsii.String("id"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("customer_id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("Customers"),
	})

	return customerTable
}
