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
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/aws/aws-sdk-go/aws"
)

var AlbAllowEmptyValues = []string{"tags.", "^condition."}

type AlbGenerator struct {
	AWSService
}

func (g *AlbGenerator) loadLB(svc *elbv2.ELBV2) error {
	err := svc.DescribeLoadBalancersPages(&elbv2.DescribeLoadBalancersInput{}, func(lbs *elbv2.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, lb := range lbs.LoadBalancers {
			resourceName := aws.StringValue(lb.LoadBalancerName)
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(lb.LoadBalancerArn),
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
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadLBListener(svc *elbv2.ELBV2, loadBalancerArn *string) error {
	err := svc.DescribeListenersPages(&elbv2.DescribeListenersInput{LoadBalancerArn: loadBalancerArn}, func(lcs *elbv2.DescribeListenersOutput, lastPage bool) bool {
		for _, ls := range lcs.Listeners {
			resourceName := aws.StringValue(ls.ListenerArn)
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
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
			err = g.loadLBListenerCertificate(svc, ls.ListenerArn)
			if err != nil {
				log.Println(err)
			}
		}
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadLBListenerRule(svc *elbv2.ELBV2, listenerArn *string) error {
	lsrs, err := svc.DescribeRules(&elbv2.DescribeRulesInput{ListenerArn: listenerArn})
	if err != nil {
		return err
	}
	for _, lsr := range lsrs.Rules {
		if !aws.BoolValue(lsr.IsDefault) {
			resourceName := aws.StringValue(lsr.RuleArn)
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_lb_listener_rule",
				"aws",
				AlbAllowEmptyValues,
			))
		}
	}
	return err
}

func (g *AlbGenerator) loadLBListenerCertificate(svc *elbv2.ELBV2, listenerArn *string) error {
	lcs, err := svc.DescribeListenerCertificates(&elbv2.DescribeListenerCertificatesInput{ListenerArn: listenerArn})
	if err != nil {
		return err
	}
	for _, lc := range lcs.Certificates {
		resourceName := aws.StringValue(lc.CertificateArn)
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			resourceName,
			resourceName,
			"aws_lb_listener_certificate",
			"aws",
			AlbAllowEmptyValues,
		))
	}
	return err
}

func (g *AlbGenerator) loadLBTargetGroup(svc *elbv2.ELBV2) error {
	err := svc.DescribeTargetGroupsPages(&elbv2.DescribeTargetGroupsInput{}, func(tgs *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
		for _, tg := range tgs.TargetGroups {
			resourceName := aws.StringValue(tg.TargetGroupName)
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(tg.TargetGroupArn),
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
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadTargetGroupTargets(svc *elbv2.ELBV2, targetGroupArn *string) error {
	targetHealths, err := svc.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{TargetGroupArn: targetGroupArn})
	if err != nil {
		return err
	}
	for _, tgh := range targetHealths.TargetHealthDescriptions {
		id := resource.PrefixedUniqueId(fmt.Sprintf("%s-", aws.StringValue(targetGroupArn)))
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			id,
			id,
			"aws_lb_target_group_attachment",
			"aws",
			map[string]string{
				"target_id":        aws.StringValue(tgh.Target.Id),
				"target_group_arn": aws.StringValue(targetGroupArn),
			},
			AlbAllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

// Generate TerraformResources from AWS API,
func (g *AlbGenerator) InitResources() error {
	sess := g.generateSession()
	svc := elbv2.New(sess)
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
		if r.InstanceInfo.Type != "aws_lb_listener" {
			continue
		}
		if r.InstanceState.Attributes["default_action.0.order"] == "0" {
			delete(r.Item["default_action"].([]interface{})[0].(map[string]interface{}), "order")
		}
	}

	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lb_listener_rule" {
			continue
		}
		if r.InstanceState.Attributes["action.0.order"] == "0" {
			delete(r.Item["action"].([]interface{})[0].(map[string]interface{}), "order")
		}
		for _, lb := range g.Resources {
			if lb.InstanceInfo.Type != "aws_lb_listener_certificate" {
				continue
			}
			if r.InstanceState.Attributes["certificate_arn"] == lb.InstanceState.Attributes["arn"] {
				g.Resources[i].Item["certificate_arn"] = "${aws_lb_listener_certificate." + lb.ResourceName + ".arn}"
			}
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lb" {
			continue
		}
		if val, ok := r.InstanceState.Attributes["access_logs.0.enabled"]; ok && val == "false" {
			delete(r.Item, "access_logs")
		}
	}
	return nil
}
