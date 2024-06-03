/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/gke-networking-api/crd/apis/network/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
)

// FakeGKENetworkParamSetLists implements GKENetworkParamSetListInterface
type FakeGKENetworkParamSetLists struct {
	Fake *FakeNetworkingV1alpha1
}

var gkenetworkparamsetlistsResource = v1alpha1.SchemeGroupVersion.WithResource("gkenetworkparamsetlists")

var gkenetworkparamsetlistsKind = v1alpha1.SchemeGroupVersion.WithKind("GKENetworkParamSetList")

// Get takes name of the gKENetworkParamSetList, and returns the corresponding gKENetworkParamSetList object, and an error if there is any.
func (c *FakeGKENetworkParamSetLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GKENetworkParamSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(gkenetworkparamsetlistsResource, name), &v1alpha1.GKENetworkParamSetList{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GKENetworkParamSetList), err
}
