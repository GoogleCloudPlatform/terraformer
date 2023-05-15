package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type MessageChannelGenerator struct {
	OpalService
}

func (g *MessageChannelGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal message channels: %v", err)
	}

	messageChannels, _, err := client.MessageChannelsApi.GetMessageChannels(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal message channels: %v", err)
	}

	countByName := make(map[string]int)

	for _, channel := range messageChannels.Channels {
		name := normalizeResourceName(*channel.Name)
		if count, ok := countByName[name]; ok {
			countByName[name] = count + 1
			name = normalizeResourceName(fmt.Sprintf("%s_%d", *channel.Name, count+1))
		} else {
			countByName[name] = 1
		}

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			channel.MessageChannelId,
			name,
			"opal_message_channel",
			"opal",
			[]string{},
		))
	}

	return nil
}
