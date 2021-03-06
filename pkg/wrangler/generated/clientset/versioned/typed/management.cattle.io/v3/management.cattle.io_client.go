/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"github.com/rancher/rancher/pkg/wrangler/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/types/apis/management.cattle.io/v3"
	rest "k8s.io/client-go/rest"
)

type ManagementV3Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
	UsersGetter
}

// ManagementV3Client is used to interact with features provided by the management.cattle.io group.
type ManagementV3Client struct {
	restClient rest.Interface
}

func (c *ManagementV3Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *ManagementV3Client) Users() UserInterface {
	return newUsers(c)
}

// NewForConfig creates a new ManagementV3Client for the given config.
func NewForConfig(c *rest.Config) (*ManagementV3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ManagementV3Client{client}, nil
}

// NewForConfigOrDie creates a new ManagementV3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ManagementV3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ManagementV3Client for the given RESTClient.
func New(c rest.Interface) *ManagementV3Client {
	return &ManagementV3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ManagementV3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
