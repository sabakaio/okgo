package registry

import (
	"github.com/docker/libkv/store"
)

type Model interface {
	GetName() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// Common models storage
type Registry struct {
	store      store.Store
	namePrefix string
}

// Return prefixed model name
func (r *Registry) getModelName(name string) string {
	if r.namePrefix == "" {
		return name
	}
	return r.namePrefix + "/" + name
}

// Return key-value storage write options
// TODO should be configuren on specific registry creation
func (r *Registry) options() *store.WriteOptions {
	return &store.WriteOptions{IsDir: true}
}

// Put model data into key-value storage in JSON-serialized format
func (r *Registry) Put(model Model) (err error) {
	data, err := model.Marshal()
	if err != nil {
		return
	}
	err = r.store.Put(r.getModelName(model.GetName()), data, r.options())
	return
}

// Get value fon key-value storage. You should define `Get` function
// for each model registry to unmarshall value into Model
func (r *Registry) get(name string) (v []byte, err error) {
	pair, err := r.store.Get(r.getModelName(name))
	if err != nil {
		return
	}
	return pair.Value, nil
}

// List all values in registry. You should define `List` function
// for each model registry to unmarshall value into Model
func (r *Registry) list() (values []*[]byte, err error) {
	pairs, err := r.store.List(r.getModelName(""))
	if err != nil {
		return
	}
	for _, p := range pairs {
		values = append(values, &p.Value)
	}
	return
}

// Delete a model from key-value storage by it's name
func (r *Registry) Delete(name string) error {
	return r.store.Delete(r.getModelName(name))
}

// Delete all model from key-value storage
func (r *Registry) DeleteAll() error {
	return r.store.DeleteTree(r.getModelName(""))
}
