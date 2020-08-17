/*

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
	v1alpha1 "github.com/kubeflow/arena/pkg/operators/edl-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScaleInSpecs implements ScaleInSpecInterface
type FakeScaleInSpecs struct {
	Fake *FakeEdlV1alpha1
	ns   string
}

var scaleinspecsResource = schema.GroupVersionResource{Group: "edl", Version: "v1alpha1", Resource: "scaleinspecs"}

var scaleinspecsKind = schema.GroupVersionKind{Group: "edl", Version: "v1alpha1", Kind: "ScaleInSpec"}

// Get takes name of the scaleInSpec, and returns the corresponding scaleInSpec object, and an error if there is any.
func (c *FakeScaleInSpecs) Get(name string, options v1.GetOptions) (result *v1alpha1.ScaleInSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scaleinspecsResource, c.ns, name), &v1alpha1.ScaleInSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaleInSpec), err
}

// List takes label and field selectors, and returns the list of ScaleInSpecs that match those selectors.
func (c *FakeScaleInSpecs) List(opts v1.ListOptions) (result *v1alpha1.ScaleInSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scaleinspecsResource, scaleinspecsKind, c.ns, opts), &v1alpha1.ScaleInSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaleInSpecList), err
}

// Watch returns a watch.Interface that watches the requested scaleInSpecs.
func (c *FakeScaleInSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scaleinspecsResource, c.ns, opts))

}

// Create takes the representation of a scaleInSpec and creates it.  Returns the server's representation of the scaleInSpec, and an error, if there is any.
func (c *FakeScaleInSpecs) Create(scaleInSpec *v1alpha1.ScaleInSpec) (result *v1alpha1.ScaleInSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scaleinspecsResource, c.ns, scaleInSpec), &v1alpha1.ScaleInSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaleInSpec), err
}

// Update takes the representation of a scaleInSpec and updates it. Returns the server's representation of the scaleInSpec, and an error, if there is any.
func (c *FakeScaleInSpecs) Update(scaleInSpec *v1alpha1.ScaleInSpec) (result *v1alpha1.ScaleInSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scaleinspecsResource, c.ns, scaleInSpec), &v1alpha1.ScaleInSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaleInSpec), err
}

// Delete takes name of the scaleInSpec and deletes it. Returns an error if one occurs.
func (c *FakeScaleInSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scaleinspecsResource, c.ns, name), &v1alpha1.ScaleInSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScaleInSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scaleinspecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScaleInSpecList{})
	return err
}

// Patch applies the patch and returns the patched scaleInSpec.
func (c *FakeScaleInSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScaleInSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scaleinspecsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScaleInSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaleInSpec), err
}
