// Copyright 2019, OpenTelemetry Authors
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

package fileexporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/config/configcheck"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := Factory{}
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configcheck.ValidateConfig(cfg))
}

func TestCreateMetricsExporter(t *testing.T) {
	factory := &Factory{}
	cfg := factory.CreateDefaultConfig()

	_, err := factory.CreateMetricsExporter(zap.NewNop(), cfg)
	assert.Error(t, err)
}

func TestCreateTraceExporter(t *testing.T) {
	factory := &Factory{}
	cfg := factory.CreateDefaultConfig()

	_, err := factory.CreateTraceExporter(zap.NewNop(), cfg)
	assert.Error(t, err)
}
