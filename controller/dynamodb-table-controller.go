package controller

import (
	"github.com/Appkube-awsx/awsx-dynamodb/command"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func GetDynamodbByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*dynamodb.ListTablesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetDynamodbByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*dynamodb.ListTablesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetDynamodbByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetDynamodbByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*dynamodb.ListTablesOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetDynamoDbList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetDynamodb(clientAuth *client.Auth) (*dynamodb.ListTablesOutput, error) {
	response, err := command.GetDynamoDbList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
