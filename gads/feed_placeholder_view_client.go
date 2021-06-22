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

package gads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "google.golang.org/genproto/googleapis/ads/googleads/v8/resources"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v8/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newFeedPlaceholderViewClientHook clientHook

// FeedPlaceholderViewCallOptions contains the retry settings for each method of FeedPlaceholderViewClient.
type FeedPlaceholderViewCallOptions struct {
	GetFeedPlaceholderView []gax.CallOption
}

func defaultFeedPlaceholderViewGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFeedPlaceholderViewCallOptions() *FeedPlaceholderViewCallOptions {
	return &FeedPlaceholderViewCallOptions{
		GetFeedPlaceholderView: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalFeedPlaceholderViewClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedPlaceholderViewClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedPlaceholderView(context.Context, *servicespb.GetFeedPlaceholderViewRequest, ...gax.CallOption) (*resourcespb.FeedPlaceholderView, error)
}

// FeedPlaceholderViewClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch feed placeholder views.
type FeedPlaceholderViewClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedPlaceholderViewClient

	// The call options for this service.
	CallOptions *FeedPlaceholderViewCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedPlaceholderViewClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedPlaceholderViewClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedPlaceholderViewClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedPlaceholderView returns the requested feed placeholder view in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedPlaceholderViewClient) GetFeedPlaceholderView(ctx context.Context, req *servicespb.GetFeedPlaceholderViewRequest, opts ...gax.CallOption) (*resourcespb.FeedPlaceholderView, error) {
	return c.internalClient.GetFeedPlaceholderView(ctx, req, opts...)
}

// feedPlaceholderViewGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedPlaceholderViewGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedPlaceholderViewClient
	CallOptions **FeedPlaceholderViewCallOptions

	// The gRPC API client.
	feedPlaceholderViewClient servicespb.FeedPlaceholderViewServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedPlaceholderViewClient creates a new feed placeholder view service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch feed placeholder views.
func NewFeedPlaceholderViewClient(ctx context.Context, opts ...option.ClientOption) (*FeedPlaceholderViewClient, error) {
	clientOpts := defaultFeedPlaceholderViewGRPCClientOptions()
	if newFeedPlaceholderViewClientHook != nil {
		hookOpts, err := newFeedPlaceholderViewClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FeedPlaceholderViewClient{CallOptions: defaultFeedPlaceholderViewCallOptions()}

	c := &feedPlaceholderViewGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		feedPlaceholderViewClient: servicespb.NewFeedPlaceholderViewServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedPlaceholderViewGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedPlaceholderViewGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedPlaceholderViewGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedPlaceholderViewGRPCClient) GetFeedPlaceholderView(ctx context.Context, req *servicespb.GetFeedPlaceholderViewRequest, opts ...gax.CallOption) (*resourcespb.FeedPlaceholderView, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedPlaceholderView[0:len((*c.CallOptions).GetFeedPlaceholderView):len((*c.CallOptions).GetFeedPlaceholderView)], opts...)
	var resp *resourcespb.FeedPlaceholderView
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedPlaceholderViewClient.GetFeedPlaceholderView(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
