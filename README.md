# POC Aurora
An example API with gin and Aurora Serverless

To run set AURORA_ARN and SECRET_ARN as environmental variables (mac or linux)

export AURORA_ARN="AWS_AURORA_ARN"
export SECRET_ARN="AWS_SECRET_ARN"

Then build and run the Gin server

go build 

The CDK folder has the the CDK project used to create an aurora cluster.
