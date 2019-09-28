package ali_mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type AliTopicManager interface {
	CreateSimpleTopic(topicName string) (err error)
	CreateTopic(topicName string, maxMessageSize int32, loggingEnabled bool) (err error)
	SetTopicAttributes(topicName string, maxMessageSize int32, loggingEnabled bool) (err error)
	GetTopicAttributes(topicName string) (attr TopicAttribute, err error)
	DeleteTopic(topicName string) (err error)
	ListTopic(nextMarker string, retNumber int32, prefix string) (topics Topics, err error)
	ListTopicDetail(nextMarker string, retNumber int32, prefix string) (topicDetails TopicDetails, err error)
}

type MNSTopicManager struct {
	cli     MNSClient
	decoder MNSDecoder
}

func checkTopicName(topicName string) (err error) {
	if len(topicName) > 256 {
		err = ERR_MNS_TOPIC_NAME_IS_TOO_LONG.New()
		return
	}
	return
}

func NewMNSTopicManager(client MNSClient) AliTopicManager {
	return &MNSTopicManager{
		cli:     client,
		decoder: NewAliMNSDecoder(),
	}
}

func (p *MNSTopicManager) CreateSimpleTopic(topicName string) (err error) {
	return p.CreateTopic(topicName, 65536, false)
}

func (p *MNSTopicManager) CreateTopic(topicName string, maxMessageSize int32, loggingEnabled bool) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	if err = checkMaxMessageSize(maxMessageSize); err != nil {
		return
	}

	message := CreateTopicRequest{
		MaxMessageSize: maxMessageSize,
		LoggingEnabled: loggingEnabled,
	}

	var code int
	code, err = send(p.cli, p.decoder, PUT, nil, &message, "topics/"+topicName, nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": topicName})
		return
	}

	return
}

func (p *MNSTopicManager) SetTopicAttributes(topicName string, maxMessageSize int32, loggingEnabled bool) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	if err = checkMaxMessageSize(maxMessageSize); err != nil {
		return
	}

	message := CreateTopicRequest{
		MaxMessageSize: maxMessageSize,
		LoggingEnabled: loggingEnabled,
	}

	_, err = send(p.cli, p.decoder, PUT, nil, &message, fmt.Sprintf("topics/%s?metaoverride=true", topicName), nil)
	return
}

func (p *MNSTopicManager) GetTopicAttributes(topicName string) (attr TopicAttribute, err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, GET, nil, nil, "topics/"+topicName, &attr)

	return
}

func (p *MNSTopicManager) DeleteTopic(topicName string) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, DELETE, nil, nil, "topics/"+topicName, nil)

	return
}

func (p *MNSTopicManager) ListTopic(nextMarker string, retNumber int32, prefix string) (topics Topics, err error) {

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

	_, err = send(p.cli, p.decoder, GET, header, nil, "topics", &topics)

	return
}

func (p *MNSTopicManager) ListTopicDetail(nextMarker string, retNumber int32, prefix string) (topicDetails TopicDetails, err error) {

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

	_, err = send(p.cli, p.decoder, GET, header, nil, "topics", &topicDetails)

	return
}
