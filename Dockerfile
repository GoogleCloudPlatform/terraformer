FROM golang:alpine

RUN apk --update --no-cache add bash musl-dev gcc git terraform py3-pip jq
RUN pip3 install --upgrade pip && pip3 install awscli

RUN git clone https://github.com/bridgecrewio/terraformer.git
WORKDIR ./terraformer

RUN go mod download
RUN go run build/main.go google
RUN go run build/main.go azure

COPY "." "."
RUN terraform init
RUN chmod 777 run.sh

ENTRYPOINT ["./run.sh"]
