package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-dynamodb/command"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func GetDynamodbTableNamesByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*dynamodb.ListTablesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetDynamodbTableNamesByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbTableNamesByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*dynamodb.ListTablesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetDynamodbTableNamesByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbTableNamesByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*dynamodb.ListTablesOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetDynamoDbTableList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetDynamodbTableNames(clientAuth *client.Auth) (*dynamodb.ListTablesOutput, error) {
	response, err := command.GetDynamoDbTableList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetDynamodbDetailsByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*dynamodb.DescribeTableOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetDynamodbDetailsByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbDetailssByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*dynamodb.DescribeTableOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetDynamodbDetailsByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbDetailsByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*dynamodb.DescribeTableOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetDynamoDbTableDetails(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetDynamodbDetails(clientAuth *client.Auth) ([]*dynamodb.DescribeTableOutput, error) {
	response, err := command.GetDynamoDbTableDetails(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
