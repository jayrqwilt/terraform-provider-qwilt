// Package qwiltcdn
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2024 Qwilt Inc.
package cdn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestSitesDataSource(t *testing.T) {

	t.Logf("Starting TestSitesDataResource test DEBUG")

	//os.Setenv("TF_CLI_CONFIG_FILE", "/Users/efrats/.terraformrc")

	tfBinaryPath := "terraform"

	// Create a temporary directory to hold the Terraform configuration
	tempDir, err := os.MkdirTemp("", "tf-exec-example")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %s", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the temporary directory after the test

	// Write the Terraform configuration to a file in the temporary directory
	tfFilePath := tempDir + "/main.tf"

	var curSiteName string

	terraformBuilder := NewTerraformConfigBuilder().SiteResource("test", generateSiteName(&curSiteName))
	terraformConfig := terraformBuilder.Build()
	err = os.WriteFile(tfFilePath, []byte(terraformConfig), 0644)
	assert.Equal(t, nil, err)

	// Initialize a new Terraform instance
	tf, err := tfexec.NewTerraform(tempDir, tfBinaryPath)
	assert.Equal(t, nil, err)

	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	err = tf.Apply(context.Background())
	assert.Equal(t, nil, err)

	//get the site_id to test it later with data-source
	state, err := tf.Show(context.Background())
	siteState := findStateResource(state, "qwilt_cdn_site", "test")
	siteId := fmt.Sprintf("%s", siteState.AttributeValues["site_id"])

	tempDir2, err := os.MkdirTemp("", "tf-exec-example-data-sources")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %s", err)
	}
	defer os.RemoveAll(tempDir2) // Clean up the temporary directory after the test

	terraformBuilder2 := NewTerraformConfigBuilder().SitesDataResource("detail", siteId)
	terraformConfig2 := terraformBuilder2.Build()

	t.Logf("ncjhdbcj: %s", terraformConfig2)
	tfFilePath2 := tempDir2 + "/main.tf"
	err = os.WriteFile(tfFilePath2, []byte(terraformConfig2), 0644)
	assert.Equal(t, nil, err)

	// Initialize a new Terraform instance
	tf2, err := tfexec.NewTerraform(tempDir2, tfBinaryPath)
	assert.Equal(t, nil, err)

	//varOption := tfexec.Var(fmt.Sprintf("site_id=%s", siteId))
	err = tf2.Apply(context.Background())
	assert.Equal(t, nil, err)

	// Read the output value
	output, err := tf2.Output(context.Background())
	assert.Equal(t, nil, err)

	//t.Logf("type: %s", output["site_detail"].Type)

	// Assert that the output matches the expected value
	var data map[string]interface{}
	err = json.Unmarshal(output["site_detail"].Value, &data)
	assert.Equal(t, nil, err)

	assert.Equal(t, siteId, data["site_id"])
	assert.Equal(t, curSiteName, data["site_name"])

}
