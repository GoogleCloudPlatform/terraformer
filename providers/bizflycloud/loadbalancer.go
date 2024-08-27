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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type LoadBalancerGenerator struct {
	BizflyCloudService
}

func (g *LoadBalancerGenerator) listLoadBalancers(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.LoadBalancer, error) {
	opts := &gobizfly.ListOptions{}

	loadBalancers, err := client.CloudLoadBalancer.List(ctx, opts)
	if err != nil {
		return nil, err
	}

	for _, loadBalancer := range loadBalancers {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			loadBalancer.ID,
			loadBalancer.Name,
			"bizflycloud_loadbalancer",
			"bizflycloud",
			[]string{}))
	}

	return loadBalancers, nil
}

func (g *LoadBalancerGenerator) listLoadBalancerListeners(ctx context.Context, client *gobizfly.Client, loadBalancers []*gobizfly.LoadBalancer, pools []*gobizfly.CloudLoadBalancerPool) ([]*gobizfly.CloudLoadBalancerListener, error) {
	list := []*gobizfly.CloudLoadBalancerListener{}

	opts := &gobizfly.ListOptions{}
	l7Policies := []map[string]string{}
	for _, loadBalancer := range loadBalancers {
		listeners, err := client.CloudLoadBalancer.Listeners().List(ctx, loadBalancer.ID, opts)
		if err != nil {
			return nil, err
		}

		list = append(list, listeners...)
		defaultPoolName := ""
		for _, listener := range listeners {
			for _, policy := range listener.L7Policies {
				l7Policies = append(l7Policies, map[string]string{
					"PolicyID":     policy.ID,
					"ListenerName": listener.Name,
				})
			}

			for _, pool := range pools {
				if pool.ID == listener.DefaultPoolID {
					defaultPoolName = pool.Name
					break
				}
			}

			g.Resources = append(g.Resources, terraformutils.NewResource(
				listener.ID,
				listener.Name,
				"bizflycloud_loadbalancer_listener",
				"bizflycloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{
					"load_balancer_id": "${bizflycloud_loadbalancer.tfer--" + loadBalancer.Name + ".id}",
					"default_pool_id":  "${bizflycloud_loadbalancer_pool.tfer--" + defaultPoolName + ".id}",
				}))
		}
	}

	for _, l7Policy := range l7Policies {
		policy, err := client.CloudLoadBalancer.L7Policies().Get(ctx, l7Policy["PolicyID"])
		if err != nil {
			return nil, err
		}
		if policy.Name == "" {
			policy.Name = policy.ID
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			policy.ID,
			policy.Name,
			"bizflycloud_loadbalancer_l7policy",
			"bizflycloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"listener_id": "${bizflycloud_loadbalancer_listener.tfer--" + l7Policy["ListenerName"] + ".id}"}))
	}
	return list, nil

}

func (g *LoadBalancerGenerator) listLoadBalancerPools(ctx context.Context, client *gobizfly.Client, loadBalancers []*gobizfly.LoadBalancer) ([]*gobizfly.CloudLoadBalancerPool, error) {
	list := []*gobizfly.CloudLoadBalancerPool{}

	opts := &gobizfly.ListOptions{}
	for _, loadBalancer := range loadBalancers {
		pools, err := client.CloudLoadBalancer.Pools().List(ctx, loadBalancer.ID, opts)
		if err != nil {
			return nil, err
		}

		list = append(list, pools...)
		for _, pool := range pools {
			additionalFields := map[string]interface{}{"load_balancer_id": "${bizflycloud_loadbalancer.tfer--" + loadBalancer.Name + ".id}"}
			if pool.HealthMonitorID != "" {
				healthmonitor, err := client.CloudLoadBalancer.HealthMonitors().Get(ctx, pool.HealthMonitorID)
				if err != nil {
					return nil, err
				}

				additionalFields["health_monitor"] = map[string]interface{}{
					"name":             healthmonitor.Name,
					"type":             healthmonitor.Type,
					"delay":            healthmonitor.Delay,
					"max_retries":      healthmonitor.MaxRetries,
					"max_retries_down": healthmonitor.MaxRetriesDown,
					"timeout":          healthmonitor.TimeOut,
					"http_method":      healthmonitor.HTTPMethod,
					"url_path":         healthmonitor.URLPath,
					"expected_code":    healthmonitor.ExpectedCodes,
				}
			}

			_members, err := client.CloudLoadBalancer.Members().List(ctx, pool.ID, opts)
			if err != nil {
				return nil, err
			}
			members := []map[string]interface{}{}
			for _, member := range _members {
				members = append(members, map[string]interface{}{
					"name":          member.Name,
					"weight":        member.Weight,
					"address":       member.Address,
					"protocol_port": member.ProtocolPort,
				})
			}
			if len(members) > 0 {
				additionalFields["members"] = members
			}

			g.Resources = append(g.Resources, terraformutils.NewResource(
				pool.ID,
				pool.Name,
				"bizflycloud_loadbalancer_pool",
				"bizflycloud",
				map[string]string{},
				[]string{},
				additionalFields))
		}
	}
	return list, nil
}

func (g *LoadBalancerGenerator) InitResources() error {
	client := g.generateClient()
	loadBalancers, err := g.listLoadBalancers(context.TODO(), client)
	if err != nil {
		return err
	}
	pools, err := g.listLoadBalancerPools(context.TODO(), client, loadBalancers)
	if err != nil {
		return err
	}
	_, err = g.listLoadBalancerListeners(context.TODO(), client, loadBalancers, pools)
	if err != nil {
		return err
	}
	return nil
}
