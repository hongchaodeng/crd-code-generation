/*
Copyright 2019 The Openshift Evangelists

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
package v1

import (
	v1 "github.com/openshift-evangelists/crd-code-generation/pkg/apis/example.com/v1"
	scheme "github.com/openshift-evangelists/crd-code-generation/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatabasesGetter has a method to return a DatabaseInterface.
// A group's client should implement this interface.
type DatabasesGetter interface {
	Databases(namespace string) DatabaseInterface
}

// DatabaseInterface has methods to work with Database resources.
type DatabaseInterface interface {
	Create(*v1.Database) (*v1.Database, error)
	Update(*v1.Database) (*v1.Database, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Database, error)
	List(opts meta_v1.ListOptions) (*v1.DatabaseList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Database, err error)
	DatabaseExpansion
}

// databases implements DatabaseInterface
type databases struct {
	client rest.Interface
	ns     string
}

// newDatabases returns a Databases
func newDatabases(c *ExampleV1Client, namespace string) *databases {
	return &databases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the database, and returns the corresponding database object, and an error if there is any.
func (c *databases) Get(name string, options meta_v1.GetOptions) (result *v1.Database, err error) {
	result = &v1.Database{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Databases that match those selectors.
func (c *databases) List(opts meta_v1.ListOptions) (result *v1.DatabaseList, err error) {
	result = &v1.DatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested databases.
func (c *databases) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("databases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a database and creates it.  Returns the server's representation of the database, and an error, if there is any.
func (c *databases) Create(database *v1.Database) (result *v1.Database, err error) {
	result = &v1.Database{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("databases").
		Body(database).
		Do().
		Into(result)
	return
}

// Update takes the representation of a database and updates it. Returns the server's representation of the database, and an error, if there is any.
func (c *databases) Update(database *v1.Database) (result *v1.Database, err error) {
	result = &v1.Database{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databases").
		Name(database.Name).
		Body(database).
		Do().
		Into(result)
	return
}

// Delete takes name of the database and deletes it. Returns an error if one occurs.
func (c *databases) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *databases) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched database.
func (c *databases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Database, err error) {
	result = &v1.Database{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("databases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
