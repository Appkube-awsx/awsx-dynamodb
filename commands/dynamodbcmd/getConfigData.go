/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package dynamodbcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-dynamodb/authenticater"
	"github.com/Appkube-awsx/awsx-dynamodb/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			table, _ := cmd.Flags().GetString("table")
			getDynamoDbTableDetail(region, crossAccountRoleArn, acKey, secKey, table, externalId)
		}
	},
}

func getDynamoDbTableDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, table string, externalId string) (*dynamodb.DescribeTableOutput, error) {
	log.Println("Getting dynamodb table data")
	dynamodbClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(table),
	}

	tableData, err := dynamodbClient.DescribeTable(input)
	if err != nil {
		log.Fatalln("Error: in getting dynamodb table data", err)
	}

	log.Println(tableData)
	return tableData, err
}

func init() {
	GetConfigDataCmd.Flags().StringP("table", "t", "", "dynamodb table name")

	if err := GetConfigDataCmd.MarkFlagRequired("table"); err != nil {
		fmt.Println(err)
	}
}
