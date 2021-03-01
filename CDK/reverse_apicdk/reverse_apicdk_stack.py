import json
from aws_cdk import (
    core,
    aws_iam as iam,
    aws_ec2 as ec2,
    aws_rds as rds,
    aws_secretsmanager as secretsmanager,
)


class ReverseApicdkStack(core.Stack):

    def __init__(self, scope: core.Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        vpc = ec2.Vpc(
            scope=self,
            id="aurora-VPC",
            cidr="10.10.0.0/16"
        )

        db_secret = secretsmanager.Secret(
            scope=self,
            id="templated-secret",
            generate_secret_string=secretsmanager.SecretStringGenerator(
                secret_string_template=json.dumps(
                    {"username": "testuser"}),
                generate_string_key="password",
                exclude_punctuation=True,
            )
        )

        cluster = rds.ServerlessCluster(
            scope=self,
            id="Cluster",
            engine=rds.DatabaseClusterEngine.AURORA_MYSQL,
            vpc=vpc,
            enable_data_api=True,
            default_database_name ="Racing",
            credentials=rds.Credentials.from_secret(db_secret),
            # removal_policy= core
            scaling=rds.ServerlessScalingOptions(
                # default is to pause after 5 minutes of idle time
                auto_pause=core.Duration.minutes(10),
                # min_capacity=rds.AuroraCapacityUnit.ACU_8,
                # max_capacity=rds.AuroraCapacityUnit.ACU_32
            )
        )
