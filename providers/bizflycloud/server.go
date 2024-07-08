// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bizflycloud

import (
	"context"
	"golang.org/x/exp/slices"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type ServerGenerator struct {
	BizflyCloudService
}

func (g *ServerGenerator) listCloudServerVPCNetworks(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.VPCNetwork, error) {
	networks := []*gobizfly.VPCNetwork{}

	networks, err := client.CloudServer.VPCNetworks().List(ctx)
	if err != nil {
		return networks, err
	}

	for _, network := range networks {
		networkName := network.Name
		if networkName == "" {
			networkName = network.ID
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			network.ID,
			networkName,
			"bizflycloud_vpc_network",
			"bizflycloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{"is_default": network.IsDefault}))
	}

	return networks, nil
}

func (g *ServerGenerator) listCloudServerFirewalls(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.Firewall, error) {
	firewalls := []*gobizfly.Firewall{}
	opts := &gobizfly.ListOptions{}

	firewalls, err := client.CloudServer.Firewalls().List(ctx, opts)
	if err != nil {
		return firewalls, err
	}

	for _, firewall := range firewalls {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			firewall.ID,
			firewall.Name,
			"bizflycloud_firewall",
			"bizflycloud",
			[]string{}))
	}

	return firewalls, nil
}

func (g *ServerGenerator) listCloudServerVPCNetworkInterfaces(ctx context.Context, client *gobizfly.Client, networks []*gobizfly.VPCNetwork, firewalls []*gobizfly.Firewall) ([]*gobizfly.NetworkInterface, error) {
	interfaces := []*gobizfly.NetworkInterface{}
	opts := &gobizfly.ListNetworkInterfaceOptions{}

	interfaces, err := client.CloudServer.NetworkInterfaces().List(ctx, opts)
	if err != nil {
		return interfaces, err
	}

	for _, networkInterface := range interfaces {
		networkID := networkInterface.NetworkID
		firewallIDs := []string{}
		for _, net := range networks {
			if net.ID == networkID {
				if net.Name == "" {
					networkID = "${bizflycloud_vpc_network.tfer--" + net.ID + ".id}"
				}
				networkID = "${bizflycloud_vpc_network.tfer--" + net.Name + ".id}"
			}
		}

		for _, fw := range firewalls {
			if slices.Contains(networkInterface.SecurityGroups, fw.ID) {
				firewallIDs = append(firewallIDs, "${bizflycloud_firewall.tfer--"+fw.Name+".id}")
			}
		}

		if networkInterface.Name == "" {
			networkInterface.Name = networkInterface.ID
		}

		additionalFields := map[string]interface{}{
			"fixed_ip":     networkInterface.FixedIps[0].IPAddress,
			"firewall_ids": firewallIDs,
			"network_id":   networkID,
			"name":         networkInterface.Name,
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			networkInterface.ID,
			networkInterface.Name,
			"bizflycloud_network_interface",
			"bizflycloud",
			map[string]string{},
			[]string{},
			additionalFields))
	}

	return interfaces, nil
}

func (g *ServerGenerator) listCloudServerPublicNetworkInterfaces(ctx context.Context, client *gobizfly.Client, firewalls []*gobizfly.Firewall) ([]*gobizfly.CloudServerPublicNetworkInterface, error) {
	interfaces := []*gobizfly.CloudServerPublicNetworkInterface{}

	interfaces, err := client.CloudServer.PublicNetworkInterfaces().List(ctx)
	if err != nil {
		return interfaces, err
	}

	for _, networkInterface := range interfaces {
		firewallIDs := []string{}
		for _, fw := range firewalls {
			if slices.Contains(networkInterface.SecurityGroups, fw.ID) {
				firewallIDs = append(firewallIDs, "${bizflycloud_firewall.tfer--"+fw.Name+".id}")
			}
		}

		if networkInterface.Name == "" {
			networkInterface.Name = networkInterface.ID
		}

		additionalFields := map[string]interface{}{
			"availability_zone": networkInterface.AvailabilityZone,
			"firewall_ids":      firewallIDs,
			"name":              networkInterface.Name,
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			networkInterface.ID,
			networkInterface.Name,
			"bizflycloud_wan_ip",
			"bizflycloud",
			map[string]string{},
			[]string{},
			additionalFields))
	}

	return interfaces, nil
}

func (g *ServerGenerator) listCloudServerSSHKeys(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.KeyPair, error) {
	keypairs := []*gobizfly.KeyPair{}

	keypairs, err := client.CloudServer.SSHKeys().List(ctx, &gobizfly.ListOptions{})
	if err != nil {
		return keypairs, err
	}

	for _, keypair := range keypairs {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			keypair.SSHKeyPair.Name,
			keypair.SSHKeyPair.Name,
			"bizflycloud_ssh_key",
			"bizflycloud",
			[]string{}))
	}

	return keypairs, nil
}

func (g *ServerGenerator) listServers(ctx context.Context, client *gobizfly.Client, keypairs []*gobizfly.KeyPair) ([]*gobizfly.Server, error) {
	opt := &gobizfly.ServerListOptions{}
	servers, err := client.CloudServer.List(ctx, opt)
	if err != nil {
		return servers, err
	}

	opts := &gobizfly.VolumeListOptions{}
	if err != nil {
		return nil, err
	}
	volumes, err := client.CloudServer.Volumes().List(ctx, opts)

	beingHandle := []string{}
	for _, server := range servers {
		for _, volume := range volumes {
			if slices.Contains(beingHandle, volume.ID) {
				continue
			}

			// Rootdisk availabe
			if len(volume.Attachments) == 0 && volume.Bootable == true {
				continue
			}

			// Datadisk
			volumeName := volume.Name
			if volumeName == "" {
				volumeName = volume.ID
			}
			// Volume availabe
			if len(volume.Attachments) == 0 {
				beingHandle = append(beingHandle, volume.ID)

				g.Resources = append(g.Resources, terraformutils.NewResource(
					volume.ID,
					volumeName,
					"bizflycloud_volume",
					"bizflycloud",
					map[string]string{},
					[]string{},
					map[string]interface{}{"size": volume.Size}))
			}

			// Datadisk in-use
			if len(volume.Attachments) > 0 && volume.AttachedType == "datadisk" {
				beingHandle = append(beingHandle, volume.ID)

				g.Resources = append(g.Resources, terraformutils.NewResource(
					volume.ID,
					volumeName,
					"bizflycloud_volume",
					"bizflycloud",
					map[string]string{},
					[]string{},
					map[string]interface{}{"size": volume.Size}))

				for _, attachment := range volume.Attachments {
					if attachment.ServerID != server.ID {
						continue
					}

					g.Resources = append(g.Resources, terraformutils.NewResource(
						volume.ID,
						server.Name+volumeName,
						"bizflycloud_volume_attachment",
						"bizflycloud",
						map[string]string{},
						[]string{},
						map[string]interface{}{
							"server_id": "${bizflycloud_server.tfer--" + server.Name + ".id}",
							"volume_id": "${bizflycloud_volume.tfer--" + volumeName + ".id}",
						}))
					break
				}
			}

			// Rootdisk
			if len(volume.Attachments) > 0 && volume.AttachedType == "rootdisk" {
				for _, attachment := range volume.Attachments {
					if attachment.ServerID != server.ID {
						continue
					}

					beingHandle = append(beingHandle, volume.ID)

					_keypair := server.KeyName
					for _, keypair := range keypairs {
						if keypair.SSHKeyPair.Name == server.KeyName {
							_keypair = "${bizflycloud_ssh_key.tfer--" + keypair.SSHKeyPair.Name + ".name}"
						}
					}
					g.Resources = append(g.Resources, terraformutils.NewResource(
						server.ID,
						server.Name,
						"bizflycloud_server",
						"bizflycloud",
						map[string]string{},
						[]string{},
						map[string]interface{}{
							"root_disk_size": volume.Size,
							"ssh_key":        _keypair,
						}))

					break
				}
			}
		}
	}
	return servers, nil
}

func (g *ServerGenerator) listNetworkAttachments(ctx context.Context, client *gobizfly.Client, servers []*gobizfly.Server, privateInterfaces []*gobizfly.NetworkInterface, publicInterfaces []*gobizfly.CloudServerPublicNetworkInterface) error {

	for _, server := range servers {
		for _, net := range privateInterfaces {
			if net.AttachedServer.ID == server.ID {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					net.ID,
					server.Name+"-"+net.Name,
					"bizflycloud_network_interface_attachment",
					"bizflycloud",
					map[string]string{},
					[]string{},
					map[string]interface{}{
						"server_id":            "${bizflycloud_server.tfer--" + server.Name + ".id}",
						"network_interface_id": "${bizflycloud_network_interface.tfer--" + net.Name + ".id}",
					}))
			}
		}

		for _, net := range publicInterfaces {
			if net.AttachedServer.ID == server.ID {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					net.ID,
					server.Name+"-"+net.Name,
					"bizflycloud_network_interface_attachment",
					"bizflycloud",
					map[string]string{},
					[]string{},
					map[string]interface{}{
						"server_id":            "${bizflycloud_server.tfer--" + server.Name + ".id}",
						"network_interface_id": "${bizflycloud_wan_ip.tfer--" + net.Name + ".id}",
					}))
			}
		}
	}
	return nil
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	networks, err := g.listCloudServerVPCNetworks(context.TODO(), client)
	firewalls, err := g.listCloudServerFirewalls(context.TODO(), client)
	vpcInterfaces, err := g.listCloudServerVPCNetworkInterfaces(context.TODO(), client, networks, firewalls)
	publicInterfaces, err := g.listCloudServerPublicNetworkInterfaces(context.TODO(), client, firewalls)
	keypairs, err := g.listCloudServerSSHKeys(context.TODO(), client)
	servers, err := g.listServers(context.TODO(), client, keypairs)
	err = g.listNetworkAttachments(context.TODO(), client, servers, vpcInterfaces, publicInterfaces)
	if err != nil {
		return err
	}

	return nil
}
