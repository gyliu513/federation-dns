/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/kubernetes-sigs/federation-v2/pkg/apis/multiclusterdns"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	multiclusterdnsMultiClusterDNSLbStorage = builders.NewApiResource( // Resource status endpoint
		multiclusterdns.InternalMultiClusterDNSLb,
		MultiClusterDNSLbSchemeFns{},
		func() runtime.Object { return &MultiClusterDNSLb{} },     // Register versioned resource
		func() runtime.Object { return &MultiClusterDNSLbList{} }, // Register versioned resource list
		&MultiClusterDNSLbStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("multiclusterdns.k8s.io", "v1alpha1").WithResources(
		multiclusterdnsMultiClusterDNSLbStorage,
		builders.NewApiResource( // Resource status endpoint
			multiclusterdns.InternalMultiClusterDNSLbStatus,
			MultiClusterDNSLbSchemeFns{},
			func() runtime.Object { return &MultiClusterDNSLb{} },     // Register versioned resource
			func() runtime.Object { return &MultiClusterDNSLbList{} }, // Register versioned resource list
			&MultiClusterDNSLbStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// MultiClusterDNSLb Functions and Structs
//
// +k8s:deepcopy-gen=false
type MultiClusterDNSLbSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type MultiClusterDNSLbStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type MultiClusterDNSLbStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterDNSLbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiClusterDNSLb `json:"items"`
}