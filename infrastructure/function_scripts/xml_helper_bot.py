import boto3
import time
import json
import os

instance_id = os.environ.get("INSTANCE_ID")


def command(ssm, instance, commands):

    r = ssm.send_command(
        InstanceIds=[instance],
        DocumentName="AWS-RunShellScript",
        Parameters={
            "commands": commands
        }
    )
    command_id = r['Command']['CommandId']

    while True:
        time.sleep(1)
        res = ssm.list_command_invocations(CommandId=command_id)

        invocations = res['CommandInvocations']
        if len(invocations) <= 0:
            continue

        output = invocations[0]['StandardOutputUrl']

        status = invocations[0]['Status']
        if status == 'Success':
            return True


def main(event, context):
    client = boto3.client('ssm', region_name='ap-northeast-1')
    res = command(client, instance_id, [
                  "cd /home/ubuntu/code-database", "./xml_update"])
    statusCode = 500
    if res:
        statusCode = 200
    return {
        'statusCode': statusCode,
        'body': json.dumps('succeeded!')
    }
