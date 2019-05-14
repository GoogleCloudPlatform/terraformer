package openstack

import (
	"fmt"
	"strings"
)

func resourceNetworkingRouterRouteV2BuildID(routerID, dstCIDR, nextHop string) string {
	return fmt.Sprintf("%s-route-%s-%s", routerID, dstCIDR, nextHop)
}

func resourceNetworkingRouterRouteV2ParseID(routeID string) (string, string, string, error) {
	routeIDAllParts := strings.Split(routeID, "-route-")
	if len(routeIDAllParts) != 2 {
		return "", "", "", fmt.Errorf("invalid ID format: %s", routeID)
	}

	routeIDLastPart := routeIDAllParts[1]
	routeIDLastParts := strings.Split(routeIDLastPart, "-")
	if len(routeIDLastParts) != 2 {
		return "", "", "", fmt.Errorf("invalid last part format for %s: %s", routeID, routeIDLastPart)
	}

	return routeIDAllParts[0], routeIDLastParts[0], routeIDLastParts[1], nil
}
