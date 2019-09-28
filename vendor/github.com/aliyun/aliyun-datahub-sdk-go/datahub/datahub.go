package datahub

import (
	"fmt"
	"time"
)

type DataHub struct {
	Client *RestClient
}

func New(accessId, accessKey, endpoint string) *DataHub {
	return &DataHub{
		Client: NewRestClient(endpoint, DefaultUserAgent(), DefaultHttpClient(), NewAliyunAccount(accessId, accessKey)),
	}
}

func NewClientWithConfig(endpoint string, config *Config, account Account) *DataHub {
	return &DataHub{
		Client: NewRestClient(endpoint, config.UserAgent, DefaultHttpClient(), account),
	}
}

// ListProjects list all projects
// It returns all project names
func (datahub *DataHub) ListProjects() (projects *Projects, err error) {
	path := PROJECTS
	projects = &Projects{}
	err = datahub.Client.Get(path, projects)
	return
}

// CreateProject create new project (Added at 2018.9)
func (datahub *DataHub) CreateProject(projectName, comment string) error {
	path := fmt.Sprintf(PROJECT, projectName)
	project := &Project{
		Comment: comment,
	}
	err := datahub.Client.Post(path, project)
	return err
}

// UpdateProject update project (Added at 2018.9)
func (datahub *DataHub) UpdateProject(projectName, comment string) error {
	path := fmt.Sprintf(PROJECT, projectName)
	project := &Project{
		Comment: comment,
	}
	err := datahub.Client.Put(path, project)
	return err
}

// DeleteProject delete project (Added at 2018.9)
func (datahub *DataHub) DeleteProject(projectName string) error {
	path := fmt.Sprintf(PROJECT, projectName)
	project := &Project{}
	err := datahub.Client.Delete(path, project)
	return err
}

// GetProject get a project deatil named the given name
// It returns Project
func (datahub *DataHub) GetProject(projectName string) (project *Project, err error) {
	path := fmt.Sprintf(PROJECT, projectName)
	project = &Project{}
	err = datahub.Client.Get(path, project)
	return
}

// ListTopics list all topic of the project named projectName
// It returns Topics
func (datahub *DataHub) ListTopics(projectName string) (topics *Topics, err error) {
	path := fmt.Sprintf(TOPICS, projectName)
	topics = &Topics{}
	err = datahub.Client.Get(path, topics)
	return
}

// GetTopic get a topic detail named the given name of the project named projectName (Changed at 2018.9)
// It return Topic
func (datahub *DataHub) GetTopic(projectName, topicName string) (topic *Topic, err error) {
	path := fmt.Sprintf(TOPIC, projectName, topicName)
	topic = &Topic{
		ProjectName: projectName,
		TopicName:   topicName,
	}
	err = datahub.Client.Get(path, topic)
	return
}

// CreateTopic create new topic
// It receives a Topic object
func (datahub *DataHub) CreateTopic(topic *Topic) error {
	path := fmt.Sprintf(TOPIC, topic.ProjectName, topic.TopicName)
	err := datahub.Client.Post(path, topic)
	return err
}

// CreateTupleTopic create new tuple topic (Added at 2018.9)
func (datahub *DataHub) CreateTupleTopic(projectName, topicName, comment string, shardCount, lifecycle int, recordSchema *RecordSchema) error {
	path := fmt.Sprintf(TOPIC, projectName, topicName)
	topic := &Topic{
		ProjectName:  projectName,
		TopicName:    topicName,
		ShardCount:   shardCount,
		Lifecycle:    lifecycle,
		RecordSchema: recordSchema,
		RecordType:   TUPLE,
		Comment:      comment,
	}
	err := datahub.Client.Post(path, topic)
	return err
}

// CreateBlobTopic create new blob topic (Added at 2018.9)
func (datahub *DataHub) CreateBlobTopic(projectName, topicName, comment string, shardCount, lifecycle int) error {
	path := fmt.Sprintf(TOPIC, projectName, topicName)
	topic := &Topic{
		ProjectName: projectName,
		TopicName:   topicName,
		ShardCount:  shardCount,
		Lifecycle:   lifecycle,
		RecordType:  BLOB,
		Comment:     comment,
	}
	err := datahub.Client.Post(path, topic)
	return err
}

// UpdateTopic update a topic (Changed at 2018.9)
func (datahub *DataHub) UpdateTopic(projectName, topicName string, lifecycle int, comment string) error {
	path := fmt.Sprintf(TOPIC, projectName, topicName)
	topic := &Topic{
		ProjectName: projectName,
		TopicName:   topicName,
		Lifecycle:   lifecycle,
		Comment:     comment,
	}
	err := datahub.Client.Put(path, topic)
	return err
}

// DeleteTopic delete a topic (Changed at 2018.9)
func (datahub *DataHub) DeleteTopic(projectName, topicName string) error {
	path := fmt.Sprintf(TOPIC, projectName, topicName)
	topic := &Topic{
		ProjectName: projectName,
		TopicName:   topicName,
	}
	err := datahub.Client.Delete(path, topic)
	return err
}

// ListShards list all shards of the given topic
// It returns []Shard
func (datahub *DataHub) ListShards(projectName, topicName string) ([]Shard, error) {
	path := fmt.Sprintf(SHARDS, projectName, topicName)
	shards := &Shards{}
	err := datahub.Client.Get(path, shards)
	if err != nil {
		return nil, err
	}
	return shards.ShardList, nil
}

// WaitAllShardsReady wait all shards ready util timeout
// If timeout < 0, it will block util all shards ready
func (datahub *DataHub) WaitAllShardsReady(projectName, topicName string, timeout int) bool {
	ready := make(chan bool)
	if timeout > 0 {
		go func(timeout int) {
			time.Sleep(time.Duration(timeout) * time.Second)
			ready <- false
		}(timeout)
	}
	go func(datahub *DataHub) {
		for {
			shards, err := datahub.ListShards(projectName, topicName)
			if err != nil {
				time.Sleep(1 * time.Microsecond)
				continue
			}
			ok := true
			for _, shard := range shards {
				switch shard.State {
				case ACTIVE, CLOSED:
					continue
				default:
					ok = false
					break
				}
			}
			if ok {
				break
			}
		}
		ready <- true
	}(datahub)

	return <-ready
}

// MergeShard merge two adjacent shards
// It returns the new shard after merged
func (datahub *DataHub) MergeShard(projectName, topicName, shardId, adjShardId string) (*ShardAbstract, error) {
	path := fmt.Sprintf(SHARDS, projectName, topicName)
	mergedShards := &MergeShard{
		Id:              shardId,
		AdjacentShardId: adjShardId,
	}
	err := datahub.Client.Post(path, mergedShards)
	if err != nil {
		return nil, err
	}
	return &mergedShards.NewShard, nil
}

// SplitShard split a shard to two adjacent shards
// It returns two new shards after split
func (datahub *DataHub) SplitShard(projectName, topicName, shardId, splitKey string) ([]ShardAbstract, error) {
	path := fmt.Sprintf(SHARDS, projectName, topicName)
	splitedShards := &SplitShard{
		Id:       shardId,
		SplitKey: splitKey,
	}
	err := datahub.Client.Post(path, splitedShards)
	if err != nil {
		return nil, err
	}
	return splitedShards.NewShards, nil
}

// GetCursor get cursor of given shard, if cursor type is "SYSTEM_TIME", the sysTime parameter must be set
// It returns Cursor
func (datahub *DataHub) GetCursor(projectName, topicName, shardId string, ct CursorType, sysTime uint64) (cursor *Cursor, err error) {
	path := fmt.Sprintf(SHARD, projectName, topicName, shardId)
	cursor = &Cursor{
		Type:       ct,
		SystemTime: sysTime,
	}
	err = datahub.Client.Post(path, cursor)
	return
}

// PutRecords put records
func (datahub *DataHub) PutRecords(projectName, topicName string, records []IRecord) (*PutResult, error) {
	path := fmt.Sprintf(SHARDS, projectName, topicName)
	recordsToPut := &PutRecords{
		Records: make([]IRecord, 0, len(records)),
	}
	for _, r := range records {
		if r != nil {
			recordsToPut.Records = append(recordsToPut.Records, r)
		}
	}
	err := datahub.Client.Post(path, recordsToPut)
	return recordsToPut.Result, err
}

// GetRecords get records
func (datahub *DataHub) GetRecords(topic *Topic, shardId, cursor string, limitNum int) (*GetResult, error) {
	path := fmt.Sprintf(SHARD, topic.ProjectName, topic.TopicName, shardId)
	records := &GetRecords{
		Cursor:       cursor,
		Limit:        limitNum,
		RecordSchema: topic.RecordSchema,
	}
	err := datahub.Client.Post(path, records)
	return records.Result, err
}

// ListSubscriptions list all subscriptions of specified topic
// It returns all subscriptions of specified topic
func (datahub *DataHub) ListSubscriptions(projectName, topicName string) (subscriptions *Subscriptions, err error) {
	path := fmt.Sprintf(SUBSCRIPTIONS, projectName, topicName)
	subscriptions = &Subscriptions{}
	err = datahub.Client.Post(path, subscriptions)
	return
}

// CreateSubscription create new subscription (Added at 2018.9)
// It returns subId
func (datahub *DataHub) CreateSubscription(projectName, topicName, comment string) (SubId string, err error) {
	path := fmt.Sprintf(SUBSCRIPTIONS, projectName, topicName)
	subscription := &Subscription{
		Comment: comment,
	}
	err = datahub.Client.Post(path, subscription)
	SubId = subscription.SubId
	return
}

// UpdateSubscription update subscription (Added at 2018.9)
func (datahub *DataHub) UpdateSubscription(projectName, topicName, subId, comment string) error {
	path := fmt.Sprintf(SUBSCRIPTION, projectName, topicName, subId)
	subscription := &Subscription{
		SubId:   subId,
		Comment: comment,
	}
	err := datahub.Client.Put(path, subscription)
	return err
}

// UpdateSubscription update subscription state (Added at 2018.9)
func (datahub *DataHub) UpdateSubscriptionState(projectName, topicName, subId string, state SubscriptionState) error {
	path := fmt.Sprintf(SUBSCRIPTION, projectName, topicName, subId)
	subscription := &Subscription{
		SubId: subId,
		State: state,
	}
	err := datahub.Client.Put(path, subscription)
	return err
}

// DeleteSubscription delete subscription (Added at 2018.9)
func (datahub *DataHub) DeleteSubscription(projectName, topicName, subId string) error {
	path := fmt.Sprintf(SUBSCRIPTION, projectName, topicName, subId)
	subscription := &Subscription{
		SubId: subId,
	}
	err := datahub.Client.Delete(path, subscription)
	return err
}

// GetSubscription get a subscription detail (Added at 2018.9)
// It returns Subscription
func (datahub *DataHub) GetSubscription(projectName, topicName, subId string) (subscription *Subscription, err error) {
	path := fmt.Sprintf(SUBSCRIPTION, projectName, topicName, subId)
	subscription = &Subscription{
		SubId: subId,
	}
	err = datahub.Client.Get(path, subscription)
	return
}
