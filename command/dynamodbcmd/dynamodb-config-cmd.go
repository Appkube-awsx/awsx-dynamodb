/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package dynamodbcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "Config data of dynamodb",
	Long:  `Config data of dynamodb`,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)

		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			table, _ := cmd.Flags().GetString("table")
			if table != "" {
				GetDynamoDbTableDetail(table, *clientAuth)
			} else {
				log.Fatalln("table name not provided. program exit")
			}
		}

	},
}

func GetDynamoDbTableDetail(table string, auth client.Auth) (*dynamodb.DescribeTableOutput, error) {
	log.Println("Getting dynamodb table data")
	dynamodbClient := client.GetClient(auth, client.DYNAMODB_CLIENT).(*dynamodb.DynamoDB)
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
