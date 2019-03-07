/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */

package config

//go:generate deepcopy-gen -i . --go-header-file ../../../hack/boilerplate/boilerplate.go.txt

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubevirtTestsConfiguration contains the configuration for KubeVirt tests
type KubeVirtTestsConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// storageClassLocal is the StorageClass to use to create local PVCs
	// Default: ""
	// +optional
	StorageClassLocal string `json:"storageClassLocal,omitempty"`
	// storageClassHostPath is the StorageClass to use to create host-path PVCs
	// Default: "host-path"
	// +optional
	StorageClassHostPath string `json:"storageClassHostPath,omitempty"`
	// storageClassBlockVolume is the StorageClass to use to create block-volume PVCs
	// Default: "block-volume"
	// +optional
	StorageClassBlockVolume string `json:"storageClassBlockVolume,omitempty"`
	// storageClassRhel is the StorageClass to use to create rhel PVCs
	// Default: "rhel"
	// +optional
	StorageClassRhel string `json:"storageClassRhel,omitempty"`
	// storageClassWindows is the StorageClass to use to create windows PVCs
	// Default: "windows"
	// +optional
	StorageClassWindows string `json:"storageClassWindows,omitempty"`
}
