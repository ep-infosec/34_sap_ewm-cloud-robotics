// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/main/LICENSE)
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/ewm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRobotConfigurations implements RobotConfigurationInterface
type FakeRobotConfigurations struct {
	Fake *FakeEwmV1alpha1
	ns   string
}

var robotconfigurationsResource = schema.GroupVersionResource{Group: "ewm.sap.com", Version: "v1alpha1", Resource: "robotconfigurations"}

var robotconfigurationsKind = schema.GroupVersionKind{Group: "ewm.sap.com", Version: "v1alpha1", Kind: "RobotConfiguration"}

// Get takes name of the robotConfiguration, and returns the corresponding robotConfiguration object, and an error if there is any.
func (c *FakeRobotConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RobotConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(robotconfigurationsResource, c.ns, name), &v1alpha1.RobotConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotConfiguration), err
}

// List takes label and field selectors, and returns the list of RobotConfigurations that match those selectors.
func (c *FakeRobotConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RobotConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(robotconfigurationsResource, robotconfigurationsKind, c.ns, opts), &v1alpha1.RobotConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RobotConfigurationList{ListMeta: obj.(*v1alpha1.RobotConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.RobotConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested robotConfigurations.
func (c *FakeRobotConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(robotconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a robotConfiguration and creates it.  Returns the server's representation of the robotConfiguration, and an error, if there is any.
func (c *FakeRobotConfigurations) Create(ctx context.Context, robotConfiguration *v1alpha1.RobotConfiguration, opts v1.CreateOptions) (result *v1alpha1.RobotConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(robotconfigurationsResource, c.ns, robotConfiguration), &v1alpha1.RobotConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotConfiguration), err
}

// Update takes the representation of a robotConfiguration and updates it. Returns the server's representation of the robotConfiguration, and an error, if there is any.
func (c *FakeRobotConfigurations) Update(ctx context.Context, robotConfiguration *v1alpha1.RobotConfiguration, opts v1.UpdateOptions) (result *v1alpha1.RobotConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(robotconfigurationsResource, c.ns, robotConfiguration), &v1alpha1.RobotConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRobotConfigurations) UpdateStatus(ctx context.Context, robotConfiguration *v1alpha1.RobotConfiguration, opts v1.UpdateOptions) (*v1alpha1.RobotConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(robotconfigurationsResource, "status", c.ns, robotConfiguration), &v1alpha1.RobotConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotConfiguration), err
}

// Delete takes name of the robotConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeRobotConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(robotconfigurationsResource, c.ns, name), &v1alpha1.RobotConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRobotConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(robotconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RobotConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched robotConfiguration.
func (c *FakeRobotConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RobotConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(robotconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RobotConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RobotConfiguration), err
}
