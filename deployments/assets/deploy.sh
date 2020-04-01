#!/bin/sh

wait()
{
    ./wait-for-it/wait-for-it.sh -t 120 $1
}

# Make sure each service that we depend on is online
# Database
wait data:5432
# API Gateway
wait aws:4567
# Cloudformation
wait aws:4581
# Cloudwatch
wait aws:4582
# DynamoDB
wait aws:4569
# IAM
wait aws:4593
# Lambda
wait aws:4574
# S3
wait aws:4572
# SNS
wait aws:4575
# SQS
wait aws:4576
# STS
wait aws:4592

# Push to local stage
serverless deploy --stage local