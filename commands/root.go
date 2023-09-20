/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-dynamodb/commands/dynamodbcmd"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
)

// AwsxDynamoDbCmd represents the base command when called without any subcommands
var AwsxDynamoDbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "get dynamodb Details command gets resource counts",
	Long:  `get dynamodb Details command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command dynamodb started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetDynamoDbList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
		
	},
}

func GetDynamoDbList(auth client.Auth) (*dynamodb.ListTablesOutput, error) {

	log.Println("Getting dynamodb list summary")

	dynamodbClient := client.GetClient(auth, client.DYNAMODB_CLIENT).(*dynamodb.DynamoDB)

	input := &dynamodb.ListTablesInput{}

	tableList, err := dynamodbClient.ListTables(input)

	if err != nil {
		log.Fatalln("Error: in getting lambda list", err)
	}

	log.Println(tableList)
	return tableList, err
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxDynamoDbCmd.Execute()

	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxDynamoDbCmd.AddCommand(dynamodbcmd.GetConfigDataCmd)
	
	AwsxDynamoDbCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxDynamoDbCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxDynamoDbCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxDynamoDbCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxDynamoDbCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxDynamoDbCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxDynamoDbCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxDynamoDbCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
