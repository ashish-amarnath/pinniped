//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	identity "go.pinniped.dev/generated/1.26/apis/concierge/identity"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*KubernetesUserInfo)(nil), (*identity.KubernetesUserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo(a.(*KubernetesUserInfo), b.(*identity.KubernetesUserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.KubernetesUserInfo)(nil), (*KubernetesUserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo(a.(*identity.KubernetesUserInfo), b.(*KubernetesUserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserInfo)(nil), (*identity.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_UserInfo_To_identity_UserInfo(a.(*UserInfo), b.(*identity.UserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.UserInfo)(nil), (*UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_UserInfo_To_v1alpha1_UserInfo(a.(*identity.UserInfo), b.(*UserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WhoAmIRequest)(nil), (*identity.WhoAmIRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WhoAmIRequest_To_identity_WhoAmIRequest(a.(*WhoAmIRequest), b.(*identity.WhoAmIRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.WhoAmIRequest)(nil), (*WhoAmIRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_WhoAmIRequest_To_v1alpha1_WhoAmIRequest(a.(*identity.WhoAmIRequest), b.(*WhoAmIRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WhoAmIRequestList)(nil), (*identity.WhoAmIRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WhoAmIRequestList_To_identity_WhoAmIRequestList(a.(*WhoAmIRequestList), b.(*identity.WhoAmIRequestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.WhoAmIRequestList)(nil), (*WhoAmIRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_WhoAmIRequestList_To_v1alpha1_WhoAmIRequestList(a.(*identity.WhoAmIRequestList), b.(*WhoAmIRequestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WhoAmIRequestSpec)(nil), (*identity.WhoAmIRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec(a.(*WhoAmIRequestSpec), b.(*identity.WhoAmIRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.WhoAmIRequestSpec)(nil), (*WhoAmIRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec(a.(*identity.WhoAmIRequestSpec), b.(*WhoAmIRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WhoAmIRequestStatus)(nil), (*identity.WhoAmIRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus(a.(*WhoAmIRequestStatus), b.(*identity.WhoAmIRequestStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*identity.WhoAmIRequestStatus)(nil), (*WhoAmIRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus(a.(*identity.WhoAmIRequestStatus), b.(*WhoAmIRequestStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo(in *KubernetesUserInfo, out *identity.KubernetesUserInfo, s conversion.Scope) error {
	if err := Convert_v1alpha1_UserInfo_To_identity_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	return nil
}

// Convert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo is an autogenerated conversion function.
func Convert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo(in *KubernetesUserInfo, out *identity.KubernetesUserInfo, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo(in, out, s)
}

func autoConvert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo(in *identity.KubernetesUserInfo, out *KubernetesUserInfo, s conversion.Scope) error {
	if err := Convert_identity_UserInfo_To_v1alpha1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	return nil
}

// Convert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo is an autogenerated conversion function.
func Convert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo(in *identity.KubernetesUserInfo, out *KubernetesUserInfo, s conversion.Scope) error {
	return autoConvert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo(in, out, s)
}

func autoConvert_v1alpha1_UserInfo_To_identity_UserInfo(in *UserInfo, out *identity.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]identity.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1alpha1_UserInfo_To_identity_UserInfo is an autogenerated conversion function.
func Convert_v1alpha1_UserInfo_To_identity_UserInfo(in *UserInfo, out *identity.UserInfo, s conversion.Scope) error {
	return autoConvert_v1alpha1_UserInfo_To_identity_UserInfo(in, out, s)
}

func autoConvert_identity_UserInfo_To_v1alpha1_UserInfo(in *identity.UserInfo, out *UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_identity_UserInfo_To_v1alpha1_UserInfo is an autogenerated conversion function.
func Convert_identity_UserInfo_To_v1alpha1_UserInfo(in *identity.UserInfo, out *UserInfo, s conversion.Scope) error {
	return autoConvert_identity_UserInfo_To_v1alpha1_UserInfo(in, out, s)
}

func autoConvert_v1alpha1_WhoAmIRequest_To_identity_WhoAmIRequest(in *WhoAmIRequest, out *identity.WhoAmIRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_WhoAmIRequest_To_identity_WhoAmIRequest is an autogenerated conversion function.
func Convert_v1alpha1_WhoAmIRequest_To_identity_WhoAmIRequest(in *WhoAmIRequest, out *identity.WhoAmIRequest, s conversion.Scope) error {
	return autoConvert_v1alpha1_WhoAmIRequest_To_identity_WhoAmIRequest(in, out, s)
}

func autoConvert_identity_WhoAmIRequest_To_v1alpha1_WhoAmIRequest(in *identity.WhoAmIRequest, out *WhoAmIRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_identity_WhoAmIRequest_To_v1alpha1_WhoAmIRequest is an autogenerated conversion function.
func Convert_identity_WhoAmIRequest_To_v1alpha1_WhoAmIRequest(in *identity.WhoAmIRequest, out *WhoAmIRequest, s conversion.Scope) error {
	return autoConvert_identity_WhoAmIRequest_To_v1alpha1_WhoAmIRequest(in, out, s)
}

func autoConvert_v1alpha1_WhoAmIRequestList_To_identity_WhoAmIRequestList(in *WhoAmIRequestList, out *identity.WhoAmIRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]identity.WhoAmIRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_WhoAmIRequestList_To_identity_WhoAmIRequestList is an autogenerated conversion function.
func Convert_v1alpha1_WhoAmIRequestList_To_identity_WhoAmIRequestList(in *WhoAmIRequestList, out *identity.WhoAmIRequestList, s conversion.Scope) error {
	return autoConvert_v1alpha1_WhoAmIRequestList_To_identity_WhoAmIRequestList(in, out, s)
}

func autoConvert_identity_WhoAmIRequestList_To_v1alpha1_WhoAmIRequestList(in *identity.WhoAmIRequestList, out *WhoAmIRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]WhoAmIRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_identity_WhoAmIRequestList_To_v1alpha1_WhoAmIRequestList is an autogenerated conversion function.
func Convert_identity_WhoAmIRequestList_To_v1alpha1_WhoAmIRequestList(in *identity.WhoAmIRequestList, out *WhoAmIRequestList, s conversion.Scope) error {
	return autoConvert_identity_WhoAmIRequestList_To_v1alpha1_WhoAmIRequestList(in, out, s)
}

func autoConvert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec(in *WhoAmIRequestSpec, out *identity.WhoAmIRequestSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec is an autogenerated conversion function.
func Convert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec(in *WhoAmIRequestSpec, out *identity.WhoAmIRequestSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_WhoAmIRequestSpec_To_identity_WhoAmIRequestSpec(in, out, s)
}

func autoConvert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec(in *identity.WhoAmIRequestSpec, out *WhoAmIRequestSpec, s conversion.Scope) error {
	return nil
}

// Convert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec is an autogenerated conversion function.
func Convert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec(in *identity.WhoAmIRequestSpec, out *WhoAmIRequestSpec, s conversion.Scope) error {
	return autoConvert_identity_WhoAmIRequestSpec_To_v1alpha1_WhoAmIRequestSpec(in, out, s)
}

func autoConvert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus(in *WhoAmIRequestStatus, out *identity.WhoAmIRequestStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_KubernetesUserInfo_To_identity_KubernetesUserInfo(&in.KubernetesUserInfo, &out.KubernetesUserInfo, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus is an autogenerated conversion function.
func Convert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus(in *WhoAmIRequestStatus, out *identity.WhoAmIRequestStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WhoAmIRequestStatus_To_identity_WhoAmIRequestStatus(in, out, s)
}

func autoConvert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus(in *identity.WhoAmIRequestStatus, out *WhoAmIRequestStatus, s conversion.Scope) error {
	if err := Convert_identity_KubernetesUserInfo_To_v1alpha1_KubernetesUserInfo(&in.KubernetesUserInfo, &out.KubernetesUserInfo, s); err != nil {
		return err
	}
	return nil
}

// Convert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus is an autogenerated conversion function.
func Convert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus(in *identity.WhoAmIRequestStatus, out *WhoAmIRequestStatus, s conversion.Scope) error {
	return autoConvert_identity_WhoAmIRequestStatus_To_v1alpha1_WhoAmIRequestStatus(in, out, s)
}
