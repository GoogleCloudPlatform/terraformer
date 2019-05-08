package kubernetes

import (
	gversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/version"
)

// Flatteners

func flattenServicePort(in []v1.ServicePort) []interface{} {
	att := make([]interface{}, len(in), len(in))
	for i, n := range in {
		m := make(map[string]interface{})
		m["name"] = n.Name
		m["protocol"] = string(n.Protocol)
		m["port"] = int(n.Port)
		m["target_port"] = n.TargetPort.String()
		m["node_port"] = int(n.NodePort)

		att[i] = m
	}
	return att
}

func flattenServiceSpec(in v1.ServiceSpec) []interface{} {
	att := make(map[string]interface{})
	if len(in.Ports) > 0 {
		att["port"] = flattenServicePort(in.Ports)
	}
	if len(in.Selector) > 0 {
		att["selector"] = in.Selector
	}
	if in.ClusterIP != "" {
		att["cluster_ip"] = in.ClusterIP
	}
	if in.Type != "" {
		att["type"] = string(in.Type)
	}
	if len(in.ExternalIPs) > 0 {
		att["external_ips"] = newStringSet(schema.HashString, in.ExternalIPs)
	}
	if in.SessionAffinity != "" {
		att["session_affinity"] = string(in.SessionAffinity)
	}
	if in.LoadBalancerIP != "" {
		att["load_balancer_ip"] = in.LoadBalancerIP
	}
	if len(in.LoadBalancerSourceRanges) > 0 {
		att["load_balancer_source_ranges"] = newStringSet(schema.HashString, in.LoadBalancerSourceRanges)
	}
	if in.ExternalName != "" {
		att["external_name"] = in.ExternalName
	}
	return []interface{}{att}
}

func flattenLoadBalancerIngress(in []v1.LoadBalancerIngress) []interface{} {
	out := make([]interface{}, len(in), len(in))
	for i, ingress := range in {
		att := make(map[string]interface{})

		att["ip"] = ingress.IP
		att["hostname"] = ingress.Hostname

		out[i] = att
	}
	return out
}

// Expanders

func expandServicePort(l []interface{}) []v1.ServicePort {
	if len(l) == 0 || l[0] == nil {
		return []v1.ServicePort{}
	}
	obj := make([]v1.ServicePort, len(l), len(l))
	for i, n := range l {
		cfg := n.(map[string]interface{})
		obj[i] = v1.ServicePort{
			Port:       int32(cfg["port"].(int)),
			TargetPort: intstr.Parse(cfg["target_port"].(string)),
		}
		if v, ok := cfg["name"].(string); ok {
			obj[i].Name = v
		}
		if v, ok := cfg["protocol"].(string); ok {
			obj[i].Protocol = v1.Protocol(v)
		}
		if v, ok := cfg["node_port"].(int); ok {
			obj[i].NodePort = int32(v)
		}
	}
	return obj
}

func expandServiceSpec(l []interface{}) v1.ServiceSpec {
	if len(l) == 0 || l[0] == nil {
		return v1.ServiceSpec{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.ServiceSpec{}

	if v, ok := in["port"].([]interface{}); ok && len(v) > 0 {
		obj.Ports = expandServicePort(v)
	}
	if v, ok := in["selector"].(map[string]interface{}); ok && len(v) > 0 {
		obj.Selector = expandStringMap(v)
	}
	if v, ok := in["cluster_ip"].(string); ok {
		obj.ClusterIP = v
	}
	if v, ok := in["type"].(string); ok {
		obj.Type = v1.ServiceType(v)
	}
	if v, ok := in["external_ips"].(*schema.Set); ok && v.Len() > 0 {
		obj.ExternalIPs = sliceOfString(v.List())
	}
	if v, ok := in["session_affinity"].(string); ok {
		obj.SessionAffinity = v1.ServiceAffinity(v)
	}
	if v, ok := in["load_balancer_ip"].(string); ok {
		obj.LoadBalancerIP = v
	}
	if v, ok := in["load_balancer_source_ranges"].(*schema.Set); ok && v.Len() > 0 {
		obj.LoadBalancerSourceRanges = sliceOfString(v.List())
	}
	if v, ok := in["external_name"].(string); ok {
		obj.ExternalName = v
	}
	return obj
}

// Patch Ops

func patchServiceSpec(keyPrefix, pathPrefix string, d *schema.ResourceData, v *version.Info) (PatchOperations, error) {
	ops := make([]PatchOperation, 0, 0)
	if d.HasChange(keyPrefix + "selector") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "selector",
			Value: d.Get(keyPrefix + "selector").(map[string]interface{}),
		})
	}
	if d.HasChange(keyPrefix + "type") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "type",
			Value: d.Get(keyPrefix + "type").(string),
		})
	}
	if d.HasChange(keyPrefix + "session_affinity") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "sessionAffinity",
			Value: d.Get(keyPrefix + "session_affinity").(string),
		})
	}
	if d.HasChange(keyPrefix + "load_balancer_ip") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "loadBalancerIP",
			Value: d.Get(keyPrefix + "load_balancer_ip").(string),
		})
	}
	if d.HasChange(keyPrefix + "load_balancer_source_ranges") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "loadBalancerSourceRanges",
			Value: d.Get(keyPrefix + "load_balancer_source_ranges").(*schema.Set).List(),
		})
	}
	if d.HasChange(keyPrefix + "port") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "ports",
			Value: expandServicePort(d.Get(keyPrefix + "port").([]interface{})),
		})
	}
	if d.HasChange(keyPrefix + "external_ips") {
		k8sVersion, err := gversion.NewVersion(v.String())
		if err != nil {
			return nil, err
		}
		v1_8_0, _ := gversion.NewVersion("1.8.0")
		if k8sVersion.LessThan(v1_8_0) {
			// If we haven't done this the deprecated field would have priority
			ops = append(ops, &ReplaceOperation{
				Path:  pathPrefix + "deprecatedPublicIPs",
				Value: nil,
			})
		}

		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "externalIPs",
			Value: d.Get(keyPrefix + "external_ips").(*schema.Set).List(),
		})
	}
	if d.HasChange(keyPrefix + "external_name") {
		ops = append(ops, &ReplaceOperation{
			Path:  pathPrefix + "externalName",
			Value: d.Get(keyPrefix + "external_name").(string),
		})
	}
	return ops, nil
}
