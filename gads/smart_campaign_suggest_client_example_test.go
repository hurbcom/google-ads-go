// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gads_test

import (
	"context"

	gads "github.com/opteo/google-ads-go"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v8/services"
)

func ExampleNewSmartCampaignSuggestClient() {
	ctx := context.Background()
	c, err := gads.NewSmartCampaignSuggestClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleSmartCampaignSuggestClient_SuggestSmartCampaignBudgetOptions() {
	ctx := context.Background()
	c, err := gads.NewSmartCampaignSuggestClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.SuggestSmartCampaignBudgetOptionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SuggestSmartCampaignBudgetOptions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
