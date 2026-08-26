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

package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getSchemasClient(t *testing.T) *schemas.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "Failed to load AWS config")
	return schemas.NewFromConfig(cfg)
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	TestComposableCompleteReadonly(t, ctx)
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	t.Run("TestTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		arn := terraform.Output(t, opts, "arn")
		assert.NotEmpty(t, id)
		assert.Contains(t, arn, "arn:aws:schemas:")
		assert.Contains(t, arn, "discoverer")
	})

	t.Run("TestDiscovererViaAPI", func(t *testing.T) {
		client := getSchemasClient(t)
		id := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")
		out, err := client.DescribeDiscoverer(context.Background(), &schemas.DescribeDiscovererInput{
			DiscovererId: &id,
		})
		require.NoError(t, err, "DescribeDiscoverer should succeed")
		require.NotNil(t, out.DiscovererId)
		assert.Equal(t, id, *out.DiscovererId)
		require.NotNil(t, out.SourceArn)
		assert.Equal(t, terraform.Output(t, ctx.TerratestTerraformOptions(), "event_bus_arn"), *out.SourceArn)
	})
}
