// Package provider
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2024 Qwilt Inc.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func AddResponseSchema(resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Interact with Qwilt services./n" +
			"[See the Terraform User Guide for details about authentication.](https://docs.qwilt.com/docs/terraform-user-guide-1#authentication)",
		Attributes: map[string]schema.Attribute{
			// CDN
			"env_type": schema.StringAttribute{
				Description: "Qwilt CDN environment [prod,prestg,stage,dev]. May also be provided via QCDN_ENVTYPE environment variable.",
				MarkdownDescription: "FOR INTERNAL USE ONLY!! /n" +
					"Qwilt CDN environment [prod,prestg,stage,dev]. May also be provided via QCDN_ENVTYPE environment variable",
				Optional: true,
			},
			"username": schema.StringAttribute{
				Description:         "Username for Qwilt CDN Sites API. May also be provided via QCDN_USERNAME environment variable.",
				MarkdownDescription: "QC services username.  May also be provided via QCDN_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				Description:         "Password for Qwilt CDN Sites API. May also be provided via QCDN_PASSWORD environment variable.",
				MarkdownDescription: "QC services password. May also be provided via QCDN_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"api_key": schema.StringAttribute{
				Description:         "API key for Qwilt CDN Sites API. May also be provided via QCDN_API_KEY environment variable.",
				MarkdownDescription: "API key for Qwilt CDN Sites API. May also be provided via QCDN_API_KEY environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}
