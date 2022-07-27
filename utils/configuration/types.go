/*
Copyright 2022 The Kruise Authors.

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

package configuration

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// kruise configmap name
	KruiseConfigurationName = "kruise-configuration"

	SidecarSetPatchPodMetadataWhiteListKey = "SidecarSet_PatchPodMetadata_WhiteList"
)

type SidecarSetPatchMetadataWhiteList struct {
	Rules []SidecarSetPatchMetadataWhiteRule `json:"rules"`
}

type SidecarSetPatchMetadataWhiteRule struct {
	// selector sidecarSet against labels
	// If selector is nil, assume that the rules should apply for every sidecarSets
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// Support for regular expressions
	AllowedAnnotationKeyExprs []string `json:"allowedAnnotationKeyExprs"`
}
