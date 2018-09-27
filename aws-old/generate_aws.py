import os
REGIONS = ['eu-west-1', 'us-east-1']

SERVICES = [
#    'alb',  # ALB
#    'asg',  # AutoScaling Group
#    'cwa',  # CloudWatch Alarm
#    'dbpg',  # Database Parameter Group
#    'dbsg',  # Database Security Group
#    'dbsn',  # Database Subnet Group
    'ec2',  # EC2
    'ecc',  # ElastiCache Cluster
    'ecsn',  # ElastiCache Subnet Group
#    'efs',  # EFS File System
#    'eip',  # EIP
#    'elb',  # ELB
#    'iamg',  # IAM Group
#    'iamgm',  # IAM Group Membership
#    'iamgp',  # IAM Group Policy
#    'iamip',  # IAM Instance Profile
#    'iamp',  # IAM Policy
#    'iampa',  # IAM Policy Attachment
#    'iamr',  # IAM Role
#    'iamrp',  # IAM Role Policy
#    'iamu',  # IAM User
#    'iamup',  # IAM User Policy
    'igw',  # Internet Gateway
#    'kmsa',  # KMS Key Alias
#    'kmsk',  # KMS Key
#    'lc',  # Launch Configuration
    'nacl',  # Network ACL
    'nat',  # NAT Gateway
    'nif',  # Network Interface
    'r53r',  # Route53 Record
    'r53z',  # Route53 Hosted Zone
#    'rds',  # RDS
#    'rs',  # Redshift
    'rt',  # Route Table
    'rta',  # Route Table Association
#    's3',  # S3
    'sg',  # Security Group
    'sn',  # Subnet
    'snss',  # SNS Subscription
    'snst',  # SNS Topic
    'sqs',  # SQS
    'vgw',  # VPN Gateway
    'vpc',  # VPC
]


def main():
    currentPath = os.getcwd()
    for region in REGIONS:
        os.system('rm -rf ' + region)
        os.mkdir(region)
        for service in SERVICES:
            if service in ['vpc']:#['igw', 'sg', 'vpc', 'nacl', 'rt', 'rta',  'sn', 'vgw']:
                print(service)
                os.system('rm -rf ' + region + '/' + service)
                os.makedirs(region + '/' + service)
                os.chdir(region + '/' + service)
                os.system('terraforming {service} --region={region}  > {service}.tf'.format(service=service, region=region))
                os.system('terraforming {service} --region={region} --tfstate > terraform.tfstate'.format(service=service, region=region))
                os.system('terraform init')
                os.system('AWS_REGION={region} terraform refresh'.format(region=region))
                os.chdir(currentPath)



if __name__ == '__main__':
    main()
