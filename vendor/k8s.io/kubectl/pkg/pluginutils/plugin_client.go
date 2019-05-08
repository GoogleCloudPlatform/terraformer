/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pluginutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// InitClientAndConfig uses the KUBECONFIG environment variable to create
// a new rest client and config object based on the existing kubectl config
// and options passed from the plugin framework via environment variables
func InitClientAndConfig() (*restclient.Config, clientcmd.ClientConfig, error) {
	// resolve kubeconfig location, prioritizing the --config global flag,
	// then the value of the KUBECONFIG env var (if any), and defaulting
	// to ~/.kube/config as a last resort.
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	kubeconfig := filepath.Join(home, ".kube", "config")

	kubeconfigEnv := os.Getenv("KUBECONFIG")
	if len(kubeconfigEnv) > 0 {
		kubeconfig = kubeconfigEnv
	}

	configFile := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CONFIG")
	kubeConfigFile := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_KUBECONFIG")
	if len(configFile) > 0 {
		kubeconfig = configFile
	} else if len(kubeConfigFile) > 0 {
		kubeconfig = kubeConfigFile
	}

	if len(kubeconfig) == 0 {
		return nil, nil, fmt.Errorf("error initializing config. The KUBECONFIG environment variable must be defined.")
	}

	config, err := configFromPath(kubeconfig)
	if err != nil {
		return nil, nil, fmt.Errorf("error obtaining kubectl config: %v", err)
	}
	client, err := config.ClientConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("the provided credentials %q could not be used: %v", kubeconfig, err)
	}

	err = applyGlobalOptionsToConfig(client)
	if err != nil {
		return nil, nil, fmt.Errorf("error processing global plugin options: %v", err)
	}

	return client, config, nil
}

func configFromPath(path string) (clientcmd.ClientConfig, error) {
	rules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: path}
	credentials, err := rules.Load()
	if err != nil {
		return nil, fmt.Errorf("the provided credentials %q could not be loaded: %v", path, err)
	}

	overrides := &clientcmd.ConfigOverrides{
		Context: clientcmdapi.Context{
			Namespace: os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_NAMESPACE"),
		},
	}

	var cfg clientcmd.ClientConfig
	context := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CONTEXT")
	if len(context) > 0 {
		rules := clientcmd.NewDefaultClientConfigLoadingRules()
		cfg = clientcmd.NewNonInteractiveClientConfig(*credentials, context, overrides, rules)
	} else {
		cfg = clientcmd.NewDefaultClientConfig(*credentials, overrides)
	}

	return cfg, nil
}

func applyGlobalOptionsToConfig(config *restclient.Config) error {
	// impersonation config
	impersonateUser := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_AS")
	if len(impersonateUser) > 0 {
		config.Impersonate.UserName = impersonateUser
	}

	impersonateGroup := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_AS_GROUP")
	if len(impersonateGroup) > 0 {
		impersonateGroupJSON := []string{}
		err := json.Unmarshal([]byte(impersonateGroup), &impersonateGroupJSON)
		if err != nil {
			return errors.New(fmt.Sprintf("error parsing global option %q: %v", "--as-group", err))
		}
		if len(impersonateGroupJSON) > 0 {
			config.Impersonate.Groups = impersonateGroupJSON
		}
	}

	// tls config
	caFile := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CERTIFICATE_AUTHORITY")
	if len(caFile) > 0 {
		config.TLSClientConfig.CAFile = caFile
	}

	clientCertFile := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CLIENT_CERTIFICATE")
	if len(clientCertFile) > 0 {
		config.TLSClientConfig.CertFile = clientCertFile
	}

	clientKey := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CLIENT_KEY")
	if len(clientKey) > 0 {
		config.TLSClientConfig.KeyFile = clientKey
	}

	cluster := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CLUSTER")
	if len(cluster) > 0 {
		// TODO(jvallejo): figure out how to override kubeconfig options
	}

	user := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_USER")
	if len(user) > 0 {
		// TODO(jvallejo): figure out how to override kubeconfig options
	}

	// user / misc request config
	requestTimeout := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_REQUEST_TIMEOUT")
	if len(requestTimeout) > 0 {
		t, err := time.ParseDuration(requestTimeout)
		if err != nil {
			return errors.New(fmt.Sprintf("%v", err))
		}
		config.Timeout = t
	}

	server := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_SERVER")
	if len(server) > 0 {
		config.ServerName = server
	}

	token := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_TOKEN")
	if len(token) > 0 {
		config.BearerToken = token
	}

	username := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_USERNAME")
	if len(username) > 0 {
		config.Username = username
	}

	password := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_PASSWORD")
	if len(password) > 0 {
		config.Password = password
	}

	return nil
}
