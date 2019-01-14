package main

import "github.com/deckarep/golang-set"

type providerImporter interface {
	getProviderData(arg ...string) map[string]interface{}
	getResourceConnections() map[string]map[string][]string
	getNotInfraService() mapset.Set
	getAccount() string
	getName() string
}
