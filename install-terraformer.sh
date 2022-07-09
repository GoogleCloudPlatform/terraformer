export PROVIDER=all
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-${PROVIDER}-linux-amd64
chmod +x terraformer-${PROVIDER}-linux-amd64
sudo mv terraformer-${PROVIDER}-linux-amd64 /usr/local/bin/terraformer

terraformer version

sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo
sudo yum -y install terraform

terraform version

cd ~/.terraform.d/plugins/linux_amd64/
sudo wget https://releases.hashicorp.com/terraform-provider-aws/4.13.0/terraform-provider-aws_4.13.0_linux_amd64.zip
sudo unzip terraform-provider-aws_4.13.0_linux_amd64.zip

terraformer import aws --resources=ec2_instance,eip,igw,route_table,sg,subnet,vpc,logs --connect=true --regions=ap-southeast-1 --profile=""
terraformer import aws --resources=cloudformation --connect=true --regions=ap-southeast-1 --profile=""

cd ~

aws s3 rm s3://280887266599-bobono-cf/sftp-tf/ --recursive
aws s3 cp generated/aws/ s3://280887266599-bobono-cf/sftp-tf/ --recursive

aws s3 cp s3://280887266599-bobono-cf/sftp-tf/ . --recursive