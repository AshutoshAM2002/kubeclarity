// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	ScannerRemoteGrypeServerAddress = "SCANNER_REMOTE_GRYPE_SERVER_ADDRESS"
	ScannerRemoteGrypeServerTimeout = "SCANNER_REMOTE_GRYPE_SERVER_TIMEOUT"
)

type RemoteGrypeConfig struct {
	GrypeServerAddress string        `yaml:"grype_server_address" mapstructure:"grype_server_address"`
	GrypeServerTimeout time.Duration `yaml:"grype_server_timeout" mapstructure:"grype_server_timeout"`
}

func loadRemoteGrypeConfig() RemoteGrypeConfig {
	setRemoteGrypeScannerConfigDefaults()
	return RemoteGrypeConfig{
		GrypeServerAddress: viper.GetString(ScannerRemoteGrypeServerAddress),
		GrypeServerTimeout: viper.GetDuration(ScannerRemoteGrypeServerTimeout),
	}
}

func setRemoteGrypeScannerConfigDefaults() {
	viper.SetDefault(ScannerRemoteGrypeServerAddress, "kubeclarity-grype-server.kubeclarity:9991")
	viper.SetDefault(ScannerRemoteGrypeServerTimeout, 2*time.Minute) // nolint:gomnd
}
