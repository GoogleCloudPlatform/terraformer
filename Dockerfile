FROM golang:alpine

RUN apk --update --no-cache add bash musl-dev gcc py3-pip jq
RUN wget https://releases.hashicorp.com/terraform/0.12.29/terraform_0.12.29_linux_amd64.zip \
    && unzip terraform_0.12.29_linux_amd64.zip \
    && mv terraform /usr/local/bin/
RUN pip3 install --upgrade pip && pip3 install awscli

WORKDIR ./terraformer
COPY "." "."
RUN go mod download
RUN go run build/main.go google
RUN go run build/main.go azure
RUN go run build/main.go aws

RUN terraform init
RUN chmod 777 run.sh

ENTRYPOINT ["./run.sh"]
