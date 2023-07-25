/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ConnectionProfile
func (mg *ConnectionProfile) GetTerraformResourceType() string {
	return "google_datastream_connection_profile"
}

// GetConnectionDetailsMapping for this ConnectionProfile
func (tr *ConnectionProfile) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"forward_ssh_connectivity[*].password": "spec.forProvider.forwardSshConnectivity[*].passwordSecretRef", "forward_ssh_connectivity[*].private_key": "spec.forProvider.forwardSshConnectivity[*].privateKeySecretRef", "mysql_profile[*].password": "spec.forProvider.mysqlProfile[*].passwordSecretRef", "mysql_profile[*].ssl_config[*].ca_certificate": "spec.forProvider.mysqlProfile[*].sslConfig[*].caCertificateSecretRef", "mysql_profile[*].ssl_config[*].client_certificate": "spec.forProvider.mysqlProfile[*].sslConfig[*].clientCertificateSecretRef", "mysql_profile[*].ssl_config[*].client_key": "spec.forProvider.mysqlProfile[*].sslConfig[*].clientKeySecretRef", "oracle_profile[*].password": "spec.forProvider.oracleProfile[*].passwordSecretRef", "postgresql_profile[*].password": "spec.forProvider.postgresqlProfile[*].passwordSecretRef"}
}

// GetObservation of this ConnectionProfile
func (tr *ConnectionProfile) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConnectionProfile
func (tr *ConnectionProfile) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConnectionProfile
func (tr *ConnectionProfile) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConnectionProfile
func (tr *ConnectionProfile) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConnectionProfile
func (tr *ConnectionProfile) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ConnectionProfile
func (tr *ConnectionProfile) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this ConnectionProfile
func (tr *ConnectionProfile) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this ConnectionProfile using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConnectionProfile) LateInitialize(attrs []byte) (bool, error) {
	params := &ConnectionProfileParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConnectionProfile) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this PrivateConnection
func (mg *PrivateConnection) GetTerraformResourceType() string {
	return "google_datastream_private_connection"
}

// GetConnectionDetailsMapping for this PrivateConnection
func (tr *PrivateConnection) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this PrivateConnection
func (tr *PrivateConnection) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PrivateConnection
func (tr *PrivateConnection) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PrivateConnection
func (tr *PrivateConnection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PrivateConnection
func (tr *PrivateConnection) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PrivateConnection
func (tr *PrivateConnection) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this PrivateConnection
func (tr *PrivateConnection) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this PrivateConnection
func (tr *PrivateConnection) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this PrivateConnection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PrivateConnection) LateInitialize(attrs []byte) (bool, error) {
	params := &PrivateConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PrivateConnection) GetTerraformSchemaVersion() int {
	return 0
}
