package ali_mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type AliMNSTopic interface {
	Name() string
	GenerateQueueEndpoint(queueName string) string
	GenerateMailEndpoint(mailAddress string) string

	PublishMessage(message MessagePublishRequest) (resp MessageSendResponse, err error)

	Subscribe(subscriptionName string, message MessageSubsribeRequest) (err error)
	SetSubscriptionAttributes(subscriptionName string, notifyStrategy NotifyStrategyType) (err error)
	GetSubscriptionAttributes(subscriptionName string) (attr SubscriptionAttribute, err error)
	Unsubscribe(subscriptionName string) (err error)
	ListSubscriptionByTopic(nextMarker string, retNumber int32, prefix string) (subscriptions Subscriptions, err error)
	ListSubscriptionDetailByTopic(nextMarker string, retNumber int32, prefix string) (subscriptionDetails SubscriptionDetails, err error)
}

type MNSTopic struct {
	name    string
	client  MNSClient
	decoder MNSDecoder

	qpsMonitor *QPSMonitor
}

func NewMNSTopic(name string, client MNSClient, qps ...int32) AliMNSTopic {
	if name == "" {
		panic("ali_mns: topic name could not be empty")
	}

	topic := new(MNSTopic)
	topic.client = client
	topic.name = name
	topic.decoder = NewAliMNSDecoder()

	qpsLimit := DefaultTopicQPSLimit
	if qps != nil && len(qps) == 1 && qps[0] > 0 {
		qpsLimit = qps[0]
	}
	topic.qpsMonitor = NewQPSMonitor(5, qpsLimit)
	return topic
}

func (p *MNSTopic) Name() string {
	return p.name
}

func (p *MNSTopic) GenerateQueueEndpoint(queueName string) string {
	return "acs:mns:" + p.client.getRegion() + ":" + p.client.getAccountID() + ":queues/" + queueName
}

func (p *MNSTopic) GenerateMailEndpoint(mailAddress string) string {
	return "mail:directmail:" + mailAddress
}

func (p *MNSTopic) PublishMessage(message MessagePublishRequest) (resp MessageSendResponse, err error) {
	p.qpsMonitor.checkQPS()
	_, err = send(p.client, p.decoder, POST, nil, message, fmt.Sprintf("topics/%s/%s", p.name, "messages"), &resp)
	return
}

func (p *MNSTopic) Subscribe(subscriptionName string, message MessageSubsribeRequest) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	p.qpsMonitor.checkQPS()

	var code int
	code, err = send(p.client, p.decoder, PUT, nil, message, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": subscriptionName})
		return
	}
	return
}

func (p *MNSTopic) SetSubscriptionAttributes(subscriptionName string, notifyStrategy NotifyStrategyType) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	message := SetSubscriptionAttributesRequest{
		NotifyStrategy: notifyStrategy,
	}

	p.qpsMonitor.checkQPS()
	_, err = send(p.client, p.decoder, PUT, nil, message, fmt.Sprintf("topics/%s/subscriptions/%s?metaoverride=true", p.name, subscriptionName), nil)
	return
}

func (p *MNSTopic) GetSubscriptionAttributes(subscriptionName string) (attr SubscriptionAttribute, err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	_, err = send(p.client, p.decoder, GET, nil, nil, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), &attr)

	return
}

func (p *MNSTopic) Unsubscribe(subscriptionName string) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	_, err = send(p.client, p.decoder, DELETE, nil, nil, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), nil)

	return
}

func (p *MNSTopic) ListSubscriptionByTopic(nextMarker string, retNumber int32, prefix string) (subscriptions Subscriptions, err error) {
	header := map[string]string{}

	marker := strings.TrimSpace(nextMarker)
	if len(marker) > 0 {
		if marker != "" {
			header["x-mns-marker"] = marker
		}
	}

	if retNumber > 0 {
		if retNumber >= 1 && retNumber <= 1000 {
			header["x-mns-ret-number"] = strconv.Itoa(int(retNumber))
		} else {
			err = ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
			return
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	_, err = send(p.client, p.decoder, GET, header, nil, fmt.Sprintf("topics/%s/subscriptions", p.name), &subscriptions)

	return
}

func (p *MNSTopic) ListSubscriptionDetailByTopic(nextMarker string, retNumber int32, prefix string) (subscriptionDetails SubscriptionDetails, err error) {
	header := map[string]string{}

	marker := strings.TrimSpace(nextMarker)
	if len(marker) > 0 {
		if marker != "" {
			header["x-mns-marker"] = marker
		}
	}

	if retNumber > 0 {
		if retNumber >= 1 && retNumber <= 1000 {
			header["x-mns-ret-number"] = strconv.Itoa(int(retNumber))
		} else {
			err = ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
			return
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	header["x-mns-with-meta"] = "true"

	_, err = send(p.client, p.decoder, GET, header, nil, fmt.Sprintf("topics/%s/subscriptions", p.name), &subscriptionDetails)

	return
}
