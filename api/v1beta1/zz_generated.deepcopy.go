//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Licensed to Alexandre VILAIN under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Alexandre VILAIN licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/gocql/gocql"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraConsistencySpec) DeepCopyInto(out *CassandraConsistencySpec) {
	*out = *in
	if in.Consistency != nil {
		in, out := &in.Consistency, &out.Consistency
		*out = new(gocql.Consistency)
		**out = **in
	}
	if in.SerialConsistency != nil {
		in, out := &in.SerialConsistency, &out.SerialConsistency
		*out = new(gocql.SerialConsistency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraConsistencySpec.
func (in *CassandraConsistencySpec) DeepCopy() *CassandraConsistencySpec {
	if in == nil {
		return nil
	}
	out := new(CassandraConsistencySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraSpec) DeepCopyInto(out *CassandraSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Consistency != nil {
		in, out := &in.Consistency, &out.Consistency
		*out = new(CassandraConsistencySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraSpec.
func (in *CassandraSpec) DeepCopy() *CassandraSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesDurationSpec) DeepCopyInto(out *CertificatesDurationSpec) {
	*out = *in
	if in.RootCACertificate != nil {
		in, out := &in.RootCACertificate, &out.RootCACertificate
		*out = new(v1.Duration)
		**out = **in
	}
	if in.IntermediateCAsCertificates != nil {
		in, out := &in.IntermediateCAsCertificates, &out.IntermediateCAsCertificates
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ClientCertificates != nil {
		in, out := &in.ClientCertificates, &out.ClientCertificates
		*out = new(v1.Duration)
		**out = **in
	}
	if in.FrontendCertificate != nil {
		in, out := &in.FrontendCertificate, &out.FrontendCertificate
		*out = new(v1.Duration)
		**out = **in
	}
	if in.InternodeCertificate != nil {
		in, out := &in.InternodeCertificate, &out.InternodeCertificate
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesDurationSpec.
func (in *CertificatesDurationSpec) DeepCopy() *CertificatesDurationSpec {
	if in == nil {
		return nil
	}
	out := new(CertificatesDurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreSpec) DeepCopyInto(out *DatastoreSpec) {
	*out = *in
	if in.SQL != nil {
		in, out := &in.SQL, &out.SQL
		*out = new(SQLSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		*out = new(ElasticsearchSpec)
		**out = **in
	}
	if in.Cassandra != nil {
		in, out := &in.Cassandra, &out.Cassandra
		*out = new(CassandraSpec)
		(*in).DeepCopyInto(*out)
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(DatastoreTLSSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreSpec.
func (in *DatastoreSpec) DeepCopy() *DatastoreSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreTLSSpec) DeepCopyInto(out *DatastoreTLSSpec) {
	*out = *in
	if in.CertFileRef != nil {
		in, out := &in.CertFileRef, &out.CertFileRef
		*out = new(SecretKeyReference)
		**out = **in
	}
	if in.KeyFileRef != nil {
		in, out := &in.KeyFileRef, &out.KeyFileRef
		*out = new(SecretKeyReference)
		**out = **in
	}
	if in.CaFileRef != nil {
		in, out := &in.CaFileRef, &out.CaFileRef
		*out = new(SecretKeyReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreTLSSpec.
func (in *DatastoreTLSSpec) DeepCopy() *DatastoreTLSSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchIndices) DeepCopyInto(out *ElasticsearchIndices) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchIndices.
func (in *ElasticsearchIndices) DeepCopy() *ElasticsearchIndices {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchIndices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	out.Indices = in.Indices
	out.CloseIdleConnectionsInterval = in.CloseIdleConnectionsInterval
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendMTLSSpec) DeepCopyInto(out *FrontendMTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendMTLSSpec.
func (in *FrontendMTLSSpec) DeepCopy() *FrontendMTLSSpec {
	if in == nil {
		return nil
	}
	out := new(FrontendMTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternodeMTLSSpec) DeepCopyInto(out *InternodeMTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternodeMTLSSpec.
func (in *InternodeMTLSSpec) DeepCopy() *InternodeMTLSSpec {
	if in == nil {
		return nil
	}
	out := new(InternodeMTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MTLSSpec) DeepCopyInto(out *MTLSSpec) {
	*out = *in
	if in.Internode != nil {
		in, out := &in.Internode, &out.Internode
		*out = new(InternodeMTLSSpec)
		**out = **in
	}
	if in.Frontend != nil {
		in, out := &in.Frontend, &out.Frontend
		*out = new(FrontendMTLSSpec)
		**out = **in
	}
	if in.CertificatesDuration != nil {
		in, out := &in.CertificatesDuration, &out.CertificatesDuration
		*out = new(CertificatesDurationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MTLSSpec.
func (in *MTLSSpec) DeepCopy() *MTLSSpec {
	if in == nil {
		return nil
	}
	out := new(MTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsSpec) DeepCopyInto(out *MetricsSpec) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsSpec.
func (in *MetricsSpec) DeepCopy() *MetricsSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	*out = *in
	if in.ListenAddress != nil {
		in, out := &in.ListenAddress, &out.ListenAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSpec.
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSpec) DeepCopyInto(out *SQLSpec) {
	*out = *in
	if in.ConnectAttributes != nil {
		in, out := &in.ConnectAttributes, &out.ConnectAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSpec.
func (in *SQLSpec) DeepCopy() *SQLSpec {
	if in == nil {
		return nil
	}
	out := new(SQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.MembershipPort != nil {
		in, out := &in.MembershipPort, &out.MembershipPort
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesSpec) DeepCopyInto(out *ServicesSpec) {
	*out = *in
	if in.Frontend != nil {
		in, out := &in.Frontend, &out.Frontend
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Matching != nil {
		in, out := &in.Matching, &out.Matching
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesSpec.
func (in *ServicesSpec) DeepCopy() *ServicesSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalAdminToolsSpec) DeepCopyInto(out *TemporalAdminToolsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalAdminToolsSpec.
func (in *TemporalAdminToolsSpec) DeepCopy() *TemporalAdminToolsSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalAdminToolsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalCluster) DeepCopyInto(out *TemporalCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalCluster.
func (in *TemporalCluster) DeepCopy() *TemporalCluster {
	if in == nil {
		return nil
	}
	out := new(TemporalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterClient) DeepCopyInto(out *TemporalClusterClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterClient.
func (in *TemporalClusterClient) DeepCopy() *TemporalClusterClient {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalClusterClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterClientList) DeepCopyInto(out *TemporalClusterClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemporalClusterClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterClientList.
func (in *TemporalClusterClientList) DeepCopy() *TemporalClusterClientList {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalClusterClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterClientSpec) DeepCopyInto(out *TemporalClusterClientSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterClientSpec.
func (in *TemporalClusterClientSpec) DeepCopy() *TemporalClusterClientSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterClientStatus) DeepCopyInto(out *TemporalClusterClientStatus) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterClientStatus.
func (in *TemporalClusterClientStatus) DeepCopy() *TemporalClusterClientStatus {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterList) DeepCopyInto(out *TemporalClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemporalCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterList.
func (in *TemporalClusterList) DeepCopy() *TemporalClusterList {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterSpec) DeepCopyInto(out *TemporalClusterSpec) {
	*out = *in
	if in.JobTtlSecondsAfterFinished != nil {
		in, out := &in.JobTtlSecondsAfterFinished, &out.JobTtlSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(ServicesSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(TemporalUISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminTools != nil {
		in, out := &in.AdminTools, &out.AdminTools
		*out = new(TemporalAdminToolsSpec)
		**out = **in
	}
	if in.MTLS != nil {
		in, out := &in.MTLS, &out.MTLS
		*out = new(MTLSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterSpec.
func (in *TemporalClusterSpec) DeepCopy() *TemporalClusterSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalClusterStatus) DeepCopyInto(out *TemporalClusterStatus) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceStatus, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalClusterStatus.
func (in *TemporalClusterStatus) DeepCopy() *TemporalClusterStatus {
	if in == nil {
		return nil
	}
	out := new(TemporalClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalNamespace) DeepCopyInto(out *TemporalNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalNamespace.
func (in *TemporalNamespace) DeepCopy() *TemporalNamespace {
	if in == nil {
		return nil
	}
	out := new(TemporalNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalNamespaceList) DeepCopyInto(out *TemporalNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemporalNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalNamespaceList.
func (in *TemporalNamespaceList) DeepCopy() *TemporalNamespaceList {
	if in == nil {
		return nil
	}
	out := new(TemporalNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalNamespaceSpec) DeepCopyInto(out *TemporalNamespaceSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.RetentionPeriod != nil {
		in, out := &in.RetentionPeriod, &out.RetentionPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalNamespaceSpec.
func (in *TemporalNamespaceSpec) DeepCopy() *TemporalNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalNamespaceStatus) DeepCopyInto(out *TemporalNamespaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalNamespaceStatus.
func (in *TemporalNamespaceStatus) DeepCopy() *TemporalNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(TemporalNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalPersistenceSpec) DeepCopyInto(out *TemporalPersistenceSpec) {
	*out = *in
	if in.DefaultStore != nil {
		in, out := &in.DefaultStore, &out.DefaultStore
		*out = new(DatastoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VisibilityStore != nil {
		in, out := &in.VisibilityStore, &out.VisibilityStore
		*out = new(DatastoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdvancedVisibilityStore != nil {
		in, out := &in.AdvancedVisibilityStore, &out.AdvancedVisibilityStore
		*out = new(DatastoreSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalPersistenceSpec.
func (in *TemporalPersistenceSpec) DeepCopy() *TemporalPersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalPersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalUIIngressSpec) DeepCopyInto(out *TemporalUIIngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]networkingv1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalUIIngressSpec.
func (in *TemporalUIIngressSpec) DeepCopy() *TemporalUIIngressSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalUIIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalUISpec) DeepCopyInto(out *TemporalUISpec) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(TemporalUIIngressSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalUISpec.
func (in *TemporalUISpec) DeepCopy() *TemporalUISpec {
	if in == nil {
		return nil
	}
	out := new(TemporalUISpec)
	in.DeepCopyInto(out)
	return out
}
