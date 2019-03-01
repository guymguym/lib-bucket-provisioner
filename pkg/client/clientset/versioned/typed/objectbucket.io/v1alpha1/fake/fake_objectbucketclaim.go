/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/yard-turkey/lib-bucket-provisioner/pkg/apis/objectbucket.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeObjectBucketClaims implements ObjectBucketClaimInterface
type FakeObjectBucketClaims struct {
	Fake *FakeObjectbucketV1alpha1
	ns   string
}

var objectbucketclaimsResource = schema.GroupVersionResource{Group: "objectbucket.io", Version: "v1alpha1", Resource: "objectbucketclaims"}

var objectbucketclaimsKind = schema.GroupVersionKind{Group: "objectbucket.io", Version: "v1alpha1", Kind: "ObjectBucketClaim"}

// Get takes name of the objectBucketClaim, and returns the corresponding objectBucketClaim object, and an error if there is any.
func (c *FakeObjectBucketClaims) Get(name string, options v1.GetOptions) (result *v1alpha1.ObjectBucketClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectbucketclaimsResource, c.ns, name), &v1alpha1.ObjectBucketClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectBucketClaim), err
}

// List takes label and field selectors, and returns the list of ObjectBucketClaims that match those selectors.
func (c *FakeObjectBucketClaims) List(opts v1.ListOptions) (result *v1alpha1.ObjectBucketClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectbucketclaimsResource, objectbucketclaimsKind, c.ns, opts), &v1alpha1.ObjectBucketClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectBucketClaimList{ListMeta: obj.(*v1alpha1.ObjectBucketClaimList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectBucketClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectBucketClaims.
func (c *FakeObjectBucketClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectbucketclaimsResource, c.ns, opts))

}

// Create takes the representation of a objectBucketClaim and creates it.  Returns the server's representation of the objectBucketClaim, and an error, if there is any.
func (c *FakeObjectBucketClaims) Create(objectBucketClaim *v1alpha1.ObjectBucketClaim) (result *v1alpha1.ObjectBucketClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectbucketclaimsResource, c.ns, objectBucketClaim), &v1alpha1.ObjectBucketClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectBucketClaim), err
}

// Update takes the representation of a objectBucketClaim and updates it. Returns the server's representation of the objectBucketClaim, and an error, if there is any.
func (c *FakeObjectBucketClaims) Update(objectBucketClaim *v1alpha1.ObjectBucketClaim) (result *v1alpha1.ObjectBucketClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectbucketclaimsResource, c.ns, objectBucketClaim), &v1alpha1.ObjectBucketClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectBucketClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectBucketClaims) UpdateStatus(objectBucketClaim *v1alpha1.ObjectBucketClaim) (*v1alpha1.ObjectBucketClaim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objectbucketclaimsResource, "status", c.ns, objectBucketClaim), &v1alpha1.ObjectBucketClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectBucketClaim), err
}

// Delete takes name of the objectBucketClaim and deletes it. Returns an error if one occurs.
func (c *FakeObjectBucketClaims) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectbucketclaimsResource, c.ns, name), &v1alpha1.ObjectBucketClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectBucketClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectbucketclaimsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectBucketClaimList{})
	return err
}

// Patch applies the patch and returns the patched objectBucketClaim.
func (c *FakeObjectBucketClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ObjectBucketClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectbucketclaimsResource, c.ns, name, data, subresources...), &v1alpha1.ObjectBucketClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectBucketClaim), err
}
