/*
Copyright 2023 The KEDA Authors

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

	v1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTriggerAuthentications implements ClusterTriggerAuthenticationInterface
type FakeClusterTriggerAuthentications struct {
	Fake *FakeKedaV1alpha1
}

var clustertriggerauthenticationsResource = v1alpha1.SchemeGroupVersion.WithResource("clustertriggerauthentications")

var clustertriggerauthenticationsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterTriggerAuthentication")

// Get takes name of the clusterTriggerAuthentication, and returns the corresponding clusterTriggerAuthentication object, and an error if there is any.
func (c *FakeClusterTriggerAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustertriggerauthenticationsResource, name), &v1alpha1.ClusterTriggerAuthentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerAuthentication), err
}

// List takes label and field selectors, and returns the list of ClusterTriggerAuthentications that match those selectors.
func (c *FakeClusterTriggerAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterTriggerAuthenticationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustertriggerauthenticationsResource, clustertriggerauthenticationsKind, opts), &v1alpha1.ClusterTriggerAuthenticationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterTriggerAuthenticationList{ListMeta: obj.(*v1alpha1.ClusterTriggerAuthenticationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterTriggerAuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTriggerAuthentications.
func (c *FakeClusterTriggerAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustertriggerauthenticationsResource, opts))
}

// Create takes the representation of a clusterTriggerAuthentication and creates it.  Returns the server's representation of the clusterTriggerAuthentication, and an error, if there is any.
func (c *FakeClusterTriggerAuthentications) Create(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.CreateOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustertriggerauthenticationsResource, clusterTriggerAuthentication), &v1alpha1.ClusterTriggerAuthentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerAuthentication), err
}

// Update takes the representation of a clusterTriggerAuthentication and updates it. Returns the server's representation of the clusterTriggerAuthentication, and an error, if there is any.
func (c *FakeClusterTriggerAuthentications) Update(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.UpdateOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustertriggerauthenticationsResource, clusterTriggerAuthentication), &v1alpha1.ClusterTriggerAuthentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerAuthentication), err
}

// Delete takes name of the clusterTriggerAuthentication and deletes it. Returns an error if one occurs.
func (c *FakeClusterTriggerAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustertriggerauthenticationsResource, name, opts), &v1alpha1.ClusterTriggerAuthentication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTriggerAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustertriggerauthenticationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterTriggerAuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched clusterTriggerAuthentication.
func (c *FakeClusterTriggerAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustertriggerauthenticationsResource, name, pt, data, subresources...), &v1alpha1.ClusterTriggerAuthentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerAuthentication), err
}
