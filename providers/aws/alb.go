// Copyright 2018 The Terraformer Authors.
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

package aws

import (
	"context"
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var AlbAllowEmptyValues = []string{"tags.", "^condition."}

type AlbGenerator struct {
	AWSService
}

func (g *AlbGenerator) loadLB(svc *elasticloadbalancingv2.Client) error {
	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(svc, &elasticloadbalancingv2.DescribeLoadBalancersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, lb := range page.LoadBalancers {
			resourceName := StringValue(lb.LoadBalancerName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*lb.LoadBalancerArn,
				resourceName,
				"aws_lb",
				"aws",
				AlbAllowEmptyValues,
			))
			err := g.loadLBListener(svc, lb.LoadBalancerArn)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func (g *AlbGenerator) loadLBListener(svc *elasticloadbalancingv2.Client, loadBalancerArn *string) error {
	p := elasticloadbalancingv2.NewDescribeListenersPaginator(svc, &elasticloadbalancingv2.DescribeListenersInput{LoadBalancerArn: loadBalancerArn})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, ls := range page.Listeners {
			resourceName := *ls.ListenerArn
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_lb_listener",
				"aws",
				AlbAllowEmptyValues,
			))
			err := g.loadLBListenerRule(svc, ls.ListenerArn)
			if err != nil {
				log.Println(err)
			}
			err = g.loadLBListenerCertificate(svc, &ls)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func (g *AlbGenerator) loadLBListenerRule(svc *elasticloadbalancingv2.Client, listenerArn *string) error {
	var marker *string
	for {
		lsrs, err := svc.DescribeRules(context.TODO(), &elasticloadbalancingv2.DescribeRulesInput{
			ListenerArn: listenerArn,
			Marker:      marker,
			PageSize:    aws.Int32(400)},
		)
		if err != nil {
			return err
		}
		for _, lsr := range lsrs.Rules {
			if !lsr.IsDefault {
				resourceName := *lsr.RuleArn
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					resourceName,
					resourceName,
					"aws_lb_listener_rule",
					"aws",
					AlbAllowEmptyValues,
				))
			}
		}
		marker = lsrs.NextMarker
		if marker == nil {
			break
		}
	}
	return nil
}

func (g *AlbGenerator) loadLBListenerCertificate(svc *elasticloadbalancingv2.Client, loadBalancer *types.Listener) error {
	lcs, err := svc.DescribeListenerCertificates(context.TODO(), &elasticloadbalancingv2.DescribeListenerCertificatesInput{
		ListenerArn: loadBalancer.ListenerArn,
	})
	if err != nil {
		return err
	}
	for _, lc := range lcs.Certificates {
		certificateArn := *lc.CertificateArn
		if certificateArn == *loadBalancer.Certificates[0].CertificateArn { // discard default certificate
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			certificateArn,
			certificateArn,
			"aws_lb_listener_certificate",
			"aws",
			map[string]string{
				"listener_arn":    *loadBalancer.ListenerArn,
				"certificate_arn": certificateArn,
			},
			AlbAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return err
}

func (g *AlbGenerator) loadLBTargetGroup(svc *elasticloadbalancingv2.Client) error {
	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(svc, &elasticloadbalancingv2.DescribeTargetGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, tg := range page.TargetGroups {
			resourceName := StringValue(tg.TargetGroupName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*tg.TargetGroupArn,
				resourceName,
				"aws_lb_target_group",
				"aws",
				AlbAllowEmptyValues,
			))
			err := g.loadTargetGroupTargets(svc, tg.TargetGroupArn)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func (g *AlbGenerator) loadTargetGroupTargets(svc *elasticloadbalancingv2.Client, targetGroupArn *string) error {
	targetHealths, err := svc.DescribeTargetHealth(context.TODO(), &elasticloadbalancingv2.DescribeTargetHealthInput{
		TargetGroupArn: targetGroupArn,
	})
	if err != nil {
		return err
	}
	for _, tgh := range targetHealths.TargetHealthDescriptions {
		id := resource.PrefixedUniqueId(fmt.Sprintf("%s-", *targetGroupArn))
		g.Resources = append(g.Resources, terraformutils.NewResource(
			id,
			id,
			"aws_lb_target_group_attachment",
			"aws",
			map[string]string{
				"target_id":        *tgh.Target.Id,
				"target_group_arn": *targetGroupArn,
			},
			AlbAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

// Generate TerraformResources from AWS API,
func (g *AlbGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := elasticloadbalancingv2.NewFromConfig(config)
	if err := g.loadLB(svc); err != nil {
		return err
	}
	if err := g.loadLBTargetGroup(svc); err != nil {
		return err
	}
	return nil
}

func (g *AlbGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.Address.Type != "aws_lb_listener" {
			continue
		}
		if r.InstanceState.Value.GetAttr("default_action").AsValueSlice()[0].GetAttr("order").AsString() == "0" {
			instanceStateMap := r.InstanceState.Value.AsValueMap()
			rootBlockDeviceMap := instanceStateMap["default_action"].AsValueSlice()[0].AsValueMap()
			delete(rootBlockDeviceMap, "order")
			instanceStateMap["default_action"] = cty.ListVal([]cty.Value{cty.ObjectVal(rootBlockDeviceMap)})
			r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
		}
	}

	for _, r := range g.Resources {
		if r.Address.Type != "aws_lb_listener_rule" {
			continue
		}
		if r.InstanceState.Value.GetAttr("action").AsValueSlice()[0].GetAttr("order").AsString() == "0" {
			instanceStateMap := r.InstanceState.Value.AsValueMap()
			rootBlockDeviceMap := instanceStateMap["action"].AsValueSlice()[0].AsValueMap()
			delete(rootBlockDeviceMap, "order")
			instanceStateMap["action"] = cty.ListVal([]cty.Value{cty.ObjectVal(rootBlockDeviceMap)})
			r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
		}
		for _, lb := range g.Resources {
			if lb.Address.Type != "aws_lb_listener_certificate" {
				continue
			}
			if r.InstanceState.Value.GetAttr("certificate_arn").AsString() == r.InstanceState.Value.GetAttr("arn").AsString() {
				instanceStateMap := r.InstanceState.Value.AsValueMap()
				instanceStateMap["certificate_arn"] = cty.StringVal("${aws_lb_listener_certificate." + lb.Address.Name + ".arn}")
				r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
			}
		}
	}

	for _, r := range g.Resources {
		if r.Address.Type != "aws_lb" {
			continue
		}
		if r.InstanceState.Value.GetAttr("access_logs").AsValueSlice()[0].GetAttr("enabled").AsString() == "false" {
			instanceStateMap := r.InstanceState.Value.AsValueMap()
			delete(instanceStateMap, "access_logs")
			r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
		}
	}
	return nil
}
