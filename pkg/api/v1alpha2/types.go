/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DeschedulerPolicy struct {
	metav1.TypeMeta `json:",inline"`

	// Profiles
	Profiles []DeschedulerProfile `json:"profiles,omitempty"`

	// NodeSelector for a set of nodes to operate over
	NodeSelector *string `json:"nodeSelector,omitempty"`

	// MaxNoOfPodsToEvictPerNode restricts maximum of pods to be evicted per node.
	MaxNoOfPodsToEvictPerNode *uint `json:"maxNoOfPodsToEvictPerNode,omitempty"`

	// MaxNoOfPodsToEvictPerNamespace restricts maximum of pods to be evicted per namespace.
	MaxNoOfPodsToEvictPerNamespace *uint `json:"maxNoOfPodsToEvictPerNamespace,omitempty"`

	// MaxNoOfPodsToTotal restricts maximum of pods to be evicted total.
	MaxNoOfPodsToEvictTotal *uint `json:"maxNoOfPodsToEvictTotal,omitempty"`

	// MetricsCollector configures collection of metrics for actual resource utilization
	MetricsCollector MetricsCollector `json:"metricsCollector,omitempty"`

	// Prometheus enables metrics collection through Prometheus
	Prometheus Prometheus `json:"prometheus,omitempty"`
}

type DeschedulerProfile struct {
	Name          string         `json:"name"`
	PluginConfigs []PluginConfig `json:"pluginConfig"`
	Plugins       Plugins        `json:"plugins"`
}

type Plugins struct {
	PreSort           PluginSet `json:"presort"`
	Sort              PluginSet `json:"sort"`
	Deschedule        PluginSet `json:"deschedule"`
	Balance           PluginSet `json:"balance"`
	Filter            PluginSet `json:"filter"`
	PreEvictionFilter PluginSet `json:"preevictionfilter"`
}

type PluginConfig struct {
	Name string               `json:"name"`
	Args runtime.RawExtension `json:"args"`
}

type PluginSet struct {
	Enabled  []string `json:"enabled"`
	Disabled []string `json:"disabled"`
}

// MetricsCollector configures collection of metrics about actual resource utilization
type MetricsCollector struct {
	// Enabled metrics collection from kubernetes metrics.
	// Later, the collection can be extended to other providers.
	Enabled bool `json:"enabled,omitempty"`
}

type Prometheus struct {
	URL                string    `json:"url,omitempty"`
	AuthToken          AuthToken `json:"authToken,omitempty"`
	InsecureSkipVerify bool      `json:"insecureSkipVerify,omitempty"`
}

type AuthToken struct {
	// raw for a raw authentication token
	Raw string `json:"raw,omitempty"`
	// secretReference references an authentication token.
	// secrets are expected to be created under the descheduler's namespace.
	SecretReference SecretReference `json:"secretReference,omitempty"`
}

// SecretReference holds a reference to a Secret
type SecretReference struct {
	// namespace is the namespace of the secret.
	Namespace string `json:"namespace,omitempty"`
	// name is the name of the secret.
	Name string `json:"name,omitempty"`
}
