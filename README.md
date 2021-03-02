# POC Aurora
An example API with gin and Aurora Serverless

To run set AURORA_ARN and SECRET_ARN as environmental variables (mac or linux)

```
export AURORA_ARN="AWS_AURORA_ARN"
export SECRET_ARN="AWS_SECRET_ARN"
```

Then build and run the Gin server

go build .

The CDK folder has the the CDK project used to create an aurora cluster.

# Docker 

To build 

```
docker build -t poc:1.0 .
```

To run (if running in ECS you have to set a role that allows RDS and secretmanager )
```
docker run -e AURORA_ARN='AURORA_ARN_FROM_CONSOLE' -e SECRET_ARN='SECRET_ARN_FROM_CONSOLE' -p 8080:8080 poc:1.0
```

To run (if running in Locally you must specify your AWS credentials)
```
docker run -e AURORA_ARN='AURORA_ARN_FROM_CONSOLE' -e SECRET_ARN='SECRET_ARN_FROM_CONSOLE' -e AWS_REGION='REGION' -e AWS_ACCESS_KEY_ID='USER_CRED' -e AWS_SECRET_ACCESS_KEY='USER_CRED' -p 8080:8080 poc:1.0
```