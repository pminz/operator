// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gingersnap-project/operator/api/v1alpha1"
)

// DataSourceSpecApplyConfiguration represents an declarative configuration of the DataSourceSpec type for use
// with apply.
type DataSourceSpecApplyConfiguration struct {
	DbType               *v1alpha1.DBType                 `json:"dbType,omitempty"`
	ConnectionProperties map[string]string                `json:"connectionProperties,omitempty"`
	ServiceBindingRef    *NamespacedRefApplyConfiguration `json:"serviceBindingRef,omitempty"`
}

// DataSourceSpecApplyConfiguration constructs an declarative configuration of the DataSourceSpec type for use with
// apply.
func DataSourceSpec() *DataSourceSpecApplyConfiguration {
	return &DataSourceSpecApplyConfiguration{}
}

// WithDbType sets the DbType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DbType field is set to the value of the last call.
func (b *DataSourceSpecApplyConfiguration) WithDbType(value v1alpha1.DBType) *DataSourceSpecApplyConfiguration {
	b.DbType = &value
	return b
}

// WithConnectionProperties puts the entries into the ConnectionProperties field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ConnectionProperties field,
// overwriting an existing map entries in ConnectionProperties field with the same key.
func (b *DataSourceSpecApplyConfiguration) WithConnectionProperties(entries map[string]string) *DataSourceSpecApplyConfiguration {
	if b.ConnectionProperties == nil && len(entries) > 0 {
		b.ConnectionProperties = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ConnectionProperties[k] = v
	}
	return b
}

// WithServiceBindingRef sets the ServiceBindingRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceBindingRef field is set to the value of the last call.
func (b *DataSourceSpecApplyConfiguration) WithServiceBindingRef(value *NamespacedRefApplyConfiguration) *DataSourceSpecApplyConfiguration {
	b.ServiceBindingRef = value
	return b
}