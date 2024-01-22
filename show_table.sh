#!/usr/bin/bash

SOURCE="http://localhost:8000"
aws dynamodb list-tables --endpoint-url $SOURCE
