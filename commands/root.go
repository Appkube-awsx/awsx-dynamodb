/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"github.com/Appkube-awsx/awsx-dynamodb/authenticater"
	"github.com/Appkube-awsx/awsx-dynamodb/client"
	"github.com/Appkube-awsx/awsx-dynamodb/commands/dynamodbcmd"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/cobra"
	"log"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxDynamoDbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "get dynamodb Details command gets resource counts",
	Long:  `get dynamodb Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command dynamodb started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getDynamoDbList(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func getDynamoDbList(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*dynamodb.ListTablesOutput, error) {
	log.Println("Getting dynamodb list summary")
	dynamodbClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &dynamodb.ListTablesInput{}
	tableList, err := dynamodbClient.ListTables(input)
	if err != nil {
		log.Fatalln("Error: in getting lambda list", err)
	}
	log.Println(tableList)
	return tableList, err
}

//func GetConfig(region string, crossAccountRoleArn string, accessKey string, secretKey string) *configservice.GetDiscoveredResourceCountsOutput {
//	return getLambdaList(region, crossAccountRoleArn, accessKey, secretKey)
//}

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
	AwsxDynamoDbCmd.AddCommand(dynamodbcmd.GetCostDataCmd)
	AwsxDynamoDbCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxDynamoDbCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxDynamoDbCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxDynamoDbCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxDynamoDbCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxDynamoDbCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxDynamoDbCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
