# DynamoDB CLI Documentation

- [What is awsx-dynamodb](#awsx-dynamodb)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [How it works](#How-it-works)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx dynamodb
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture.svg)


# How to write plugin subcommand 
Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug  / publish and integrate into the main commmand.

# How to build / Test
            go run main.go
                - Program will print Calling aws-dynamodb on console 

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-dynamodb) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-dynamodb --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2


# How it works

The `dynamodb` command-line interface (CLI) is a tool for managing DynamoDB tables. This document provides instructions for using the `dynamodb` CLI to retrieve a list of DynamoDB tables and retrieve the configuration details of a specific table.

## List Tables

To list all the DynamoDB tables in an account, run the following command:

    awsx-dynamodb --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>
  
    awsx-dynamodb --vaultUrl <vaultUrl> --accountId <accountId> 


where:
- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a table. This is an optional parameter. 
- `--accountId` specifies the AWS account ID that the DynamoDB table belongs to.
- `--zone` specifies the AWS region where the DynamoDB table is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a table in another account. This is an optional parameter.

Example:

    awsx-dynamodb --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012 
  
    awsx-dynamodb --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

## Get Table Configuration

To retrieve the configuration details of a specific DynamoDB table, run the following command:

    awsx-dynamodb getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>
   
    awsx-dynamodb getConfigData -t <table> --vaultUrl <vaultUrl> --accountId <accountId> 

where:
- `-t` or `--table` is the shorthand for specifying the name of the DynamoDB table. This parameter is mandatory.
- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a table. This is an optional parameter. 
- `--accountId` specifies the AWS account ID that the DynamoDB table belongs to.
- `--zone` specifies the AWS region where the DynamoDB table is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a table in another account. This is an optional parameter.

Example:

    awsx-dynamodb getConfigData -t my-table --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012 
 
    awsx-dynamodb getConfigData -t my-table --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

This command returns the configuration details of the specified DynamoDB table in JSON format.
