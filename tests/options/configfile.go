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

import (
	"io/ioutil"

	"k8s.io/apimachinery/pkg/runtime"

	kubevirttestsconfig "kubevirt.io/kubevirt/tests/apis/config"
	kubevirttestsscheme "kubevirt.io/kubevirt/tests/apis/config/scheme"
)

func loadConfigFromFile(file string) (*kubevirttestsconfig.KubeVirtTestsConfiguration, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return loadConfig(data)
}

func loadConfig(data []byte) (*kubevirttestsconfig.KubeVirtTestsConfiguration, error) {
	config := &kubevirttestsconfig.KubeVirtTestsConfiguration{}
	if err := runtime.DecodeInto(kubevirttestsscheme.Codecs.UniversalDecoder(), data, config); err != nil {
		return nil, err
	}

	return config, nil
}
