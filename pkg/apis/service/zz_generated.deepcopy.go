//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package service

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEExternalAccountBinding) DeepCopyInto(out *ACMEExternalAccountBinding) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEExternalAccountBinding.
func (in *ACMEExternalAccountBinding) DeepCopy() *ACMEExternalAccountBinding {
	if in == nil {
		return nil
	}
	out := new(ACMEExternalAccountBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alerting) DeepCopyInto(out *Alerting) {
	*out = *in
	if in.CertExpirationAlertDays != nil {
		in, out := &in.CertExpirationAlertDays, &out.CertExpirationAlertDays
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alerting.
func (in *Alerting) DeepCopy() *Alerting {
	if in == nil {
		return nil
	}
	out := new(Alerting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertConfig) DeepCopyInto(out *CertConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Issuers != nil {
		in, out := &in.Issuers, &out.Issuers
		*out = make([]IssuerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSChallengeOnShoot != nil {
		in, out := &in.DNSChallengeOnShoot, &out.DNSChallengeOnShoot
		*out = new(DNSChallengeOnShoot)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootIssuers != nil {
		in, out := &in.ShootIssuers, &out.ShootIssuers
		*out = new(ShootIssuers)
		**out = **in
	}
	if in.PrecheckNameservers != nil {
		in, out := &in.PrecheckNameservers, &out.PrecheckNameservers
		*out = new(string)
		**out = **in
	}
	if in.Alerting != nil {
		in, out := &in.Alerting, &out.Alerting
		*out = new(Alerting)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertConfig.
func (in *CertConfig) DeepCopy() *CertConfig {
	if in == nil {
		return nil
	}
	out := new(CertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSChallengeOnShoot) DeepCopyInto(out *DNSChallengeOnShoot) {
	*out = *in
	if in.DNSClass != nil {
		in, out := &in.DNSClass, &out.DNSClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSChallengeOnShoot.
func (in *DNSChallengeOnShoot) DeepCopy() *DNSChallengeOnShoot {
	if in == nil {
		return nil
	}
	out := new(DNSChallengeOnShoot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSelection) DeepCopyInto(out *DNSSelection) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSelection.
func (in *DNSSelection) DeepCopy() *DNSSelection {
	if in == nil {
		return nil
	}
	out := new(DNSSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerConfig) DeepCopyInto(out *IssuerConfig) {
	*out = *in
	if in.RequestsPerDayQuota != nil {
		in, out := &in.RequestsPerDayQuota, &out.RequestsPerDayQuota
		*out = new(int)
		**out = **in
	}
	if in.PrivateKeySecretName != nil {
		in, out := &in.PrivateKeySecretName, &out.PrivateKeySecretName
		*out = new(string)
		**out = **in
	}
	if in.ExternalAccountBinding != nil {
		in, out := &in.ExternalAccountBinding, &out.ExternalAccountBinding
		*out = new(ACMEExternalAccountBinding)
		**out = **in
	}
	if in.SkipDNSChallengeValidation != nil {
		in, out := &in.SkipDNSChallengeValidation, &out.SkipDNSChallengeValidation
		*out = new(bool)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = new(DNSSelection)
		(*in).DeepCopyInto(*out)
	}
	if in.PrecheckNameservers != nil {
		in, out := &in.PrecheckNameservers, &out.PrecheckNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerConfig.
func (in *IssuerConfig) DeepCopy() *IssuerConfig {
	if in == nil {
		return nil
	}
	out := new(IssuerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootIssuers) DeepCopyInto(out *ShootIssuers) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootIssuers.
func (in *ShootIssuers) DeepCopy() *ShootIssuers {
	if in == nil {
		return nil
	}
	out := new(ShootIssuers)
	in.DeepCopyInto(out)
	return out
}
