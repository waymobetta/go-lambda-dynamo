#!/bin/zsh

# build go binary for Lambda environ
printf "building.."

{
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" .
} || {
	exit
}

printf "done\n"

# compress go binary
printf "compressing.."

{
	upx go-lambda-dynamo &> /dev/null
} || {
	exit
}

printf "done\n"

# zip binary to prepare for Lambda
printf "zipping.."

{
	zip go-lambda-dynamo.zip go-lambda-dynamo &> /dev/null
} || {
	exit
}

printf "done\n"

# upload to Lambda function
printf "deploying.."

{	
	aws lambda update-function-code \
	--function-name go-lambda-dynamo \
	--zip-file fileb://go-lambda-dynamo.zip &> /dev/null
} || {
	exit
}

wait

printf "done\n"

# remove build files
printf "cleaning up.."
{
	rm go-lambda-dynamo*
} || {
	exit
}

printf "done\n"
