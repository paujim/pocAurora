#!/usr/bin/env python3

from aws_cdk import core

from reverse_apicdk.reverse_apicdk_stack import ReverseApicdkStack


app = core.App()
ReverseApicdkStack(app, "reverse-apicdk")

app.synth()
