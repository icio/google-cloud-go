// Copyright 2022 Google LLC
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

package migration

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	migrationpb "google.golang.org/genproto/googleapis/cloud/bigquery/migration/v2alpha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateMigrationWorkflow []gax.CallOption
	GetMigrationWorkflow    []gax.CallOption
	ListMigrationWorkflows  []gax.CallOption
	DeleteMigrationWorkflow []gax.CallOption
	StartMigrationWorkflow  []gax.CallOption
	GetMigrationSubtask     []gax.CallOption
	ListMigrationSubtasks   []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerymigration.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerymigration.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://bigquerymigration.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateMigrationWorkflow: []gax.CallOption{},
		GetMigrationWorkflow: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListMigrationWorkflows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteMigrationWorkflow: []gax.CallOption{},
		StartMigrationWorkflow: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetMigrationSubtask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListMigrationSubtasks: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from BigQuery Migration API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateMigrationWorkflow(context.Context, *migrationpb.CreateMigrationWorkflowRequest, ...gax.CallOption) (*migrationpb.MigrationWorkflow, error)
	GetMigrationWorkflow(context.Context, *migrationpb.GetMigrationWorkflowRequest, ...gax.CallOption) (*migrationpb.MigrationWorkflow, error)
	ListMigrationWorkflows(context.Context, *migrationpb.ListMigrationWorkflowsRequest, ...gax.CallOption) *MigrationWorkflowIterator
	DeleteMigrationWorkflow(context.Context, *migrationpb.DeleteMigrationWorkflowRequest, ...gax.CallOption) error
	StartMigrationWorkflow(context.Context, *migrationpb.StartMigrationWorkflowRequest, ...gax.CallOption) error
	GetMigrationSubtask(context.Context, *migrationpb.GetMigrationSubtaskRequest, ...gax.CallOption) (*migrationpb.MigrationSubtask, error)
	ListMigrationSubtasks(context.Context, *migrationpb.ListMigrationSubtasksRequest, ...gax.CallOption) *MigrationSubtaskIterator
}

// Client is a client for interacting with BigQuery Migration API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to handle EDW migrations.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateMigrationWorkflow creates a migration workflow.
func (c *Client) CreateMigrationWorkflow(ctx context.Context, req *migrationpb.CreateMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	return c.internalClient.CreateMigrationWorkflow(ctx, req, opts...)
}

// GetMigrationWorkflow gets a previously created migration workflow.
func (c *Client) GetMigrationWorkflow(ctx context.Context, req *migrationpb.GetMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	return c.internalClient.GetMigrationWorkflow(ctx, req, opts...)
}

// ListMigrationWorkflows lists previously created migration workflow.
func (c *Client) ListMigrationWorkflows(ctx context.Context, req *migrationpb.ListMigrationWorkflowsRequest, opts ...gax.CallOption) *MigrationWorkflowIterator {
	return c.internalClient.ListMigrationWorkflows(ctx, req, opts...)
}

// DeleteMigrationWorkflow deletes a migration workflow by name.
func (c *Client) DeleteMigrationWorkflow(ctx context.Context, req *migrationpb.DeleteMigrationWorkflowRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteMigrationWorkflow(ctx, req, opts...)
}

// StartMigrationWorkflow starts a previously created migration workflow. I.e., the state transitions
// from DRAFT to RUNNING. This is a no-op if the state is already RUNNING.
// An error will be signaled if the state is anything other than DRAFT or
// RUNNING.
func (c *Client) StartMigrationWorkflow(ctx context.Context, req *migrationpb.StartMigrationWorkflowRequest, opts ...gax.CallOption) error {
	return c.internalClient.StartMigrationWorkflow(ctx, req, opts...)
}

// GetMigrationSubtask gets a previously created migration subtask.
func (c *Client) GetMigrationSubtask(ctx context.Context, req *migrationpb.GetMigrationSubtaskRequest, opts ...gax.CallOption) (*migrationpb.MigrationSubtask, error) {
	return c.internalClient.GetMigrationSubtask(ctx, req, opts...)
}

// ListMigrationSubtasks lists previously created migration subtasks.
func (c *Client) ListMigrationSubtasks(ctx context.Context, req *migrationpb.ListMigrationSubtasksRequest, opts ...gax.CallOption) *MigrationSubtaskIterator {
	return c.internalClient.ListMigrationSubtasks(ctx, req, opts...)
}

// gRPCClient is a client for interacting with BigQuery Migration API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client migrationpb.MigrationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new migration service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to handle EDW migrations.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           migrationpb.NewMigrationServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateMigrationWorkflow(ctx context.Context, req *migrationpb.CreateMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateMigrationWorkflow[0:len((*c.CallOptions).CreateMigrationWorkflow):len((*c.CallOptions).CreateMigrationWorkflow)], opts...)
	var resp *migrationpb.MigrationWorkflow
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetMigrationWorkflow(ctx context.Context, req *migrationpb.GetMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 120000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMigrationWorkflow[0:len((*c.CallOptions).GetMigrationWorkflow):len((*c.CallOptions).GetMigrationWorkflow)], opts...)
	var resp *migrationpb.MigrationWorkflow
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListMigrationWorkflows(ctx context.Context, req *migrationpb.ListMigrationWorkflowsRequest, opts ...gax.CallOption) *MigrationWorkflowIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMigrationWorkflows[0:len((*c.CallOptions).ListMigrationWorkflows):len((*c.CallOptions).ListMigrationWorkflows)], opts...)
	it := &MigrationWorkflowIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationWorkflowsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationWorkflow, string, error) {
		resp := &migrationpb.ListMigrationWorkflowsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListMigrationWorkflows(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMigrationWorkflows(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) DeleteMigrationWorkflow(ctx context.Context, req *migrationpb.DeleteMigrationWorkflowRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteMigrationWorkflow[0:len((*c.CallOptions).DeleteMigrationWorkflow):len((*c.CallOptions).DeleteMigrationWorkflow)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) StartMigrationWorkflow(ctx context.Context, req *migrationpb.StartMigrationWorkflowRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 120000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).StartMigrationWorkflow[0:len((*c.CallOptions).StartMigrationWorkflow):len((*c.CallOptions).StartMigrationWorkflow)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.StartMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetMigrationSubtask(ctx context.Context, req *migrationpb.GetMigrationSubtaskRequest, opts ...gax.CallOption) (*migrationpb.MigrationSubtask, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 120000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMigrationSubtask[0:len((*c.CallOptions).GetMigrationSubtask):len((*c.CallOptions).GetMigrationSubtask)], opts...)
	var resp *migrationpb.MigrationSubtask
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetMigrationSubtask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListMigrationSubtasks(ctx context.Context, req *migrationpb.ListMigrationSubtasksRequest, opts ...gax.CallOption) *MigrationSubtaskIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMigrationSubtasks[0:len((*c.CallOptions).ListMigrationSubtasks):len((*c.CallOptions).ListMigrationSubtasks)], opts...)
	it := &MigrationSubtaskIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationSubtasksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationSubtask, string, error) {
		resp := &migrationpb.ListMigrationSubtasksResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListMigrationSubtasks(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMigrationSubtasks(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// MigrationSubtaskIterator manages a stream of *migrationpb.MigrationSubtask.
type MigrationSubtaskIterator struct {
	items    []*migrationpb.MigrationSubtask
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*migrationpb.MigrationSubtask, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MigrationSubtaskIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MigrationSubtaskIterator) Next() (*migrationpb.MigrationSubtask, error) {
	var item *migrationpb.MigrationSubtask
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MigrationSubtaskIterator) bufLen() int {
	return len(it.items)
}

func (it *MigrationSubtaskIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// MigrationWorkflowIterator manages a stream of *migrationpb.MigrationWorkflow.
type MigrationWorkflowIterator struct {
	items    []*migrationpb.MigrationWorkflow
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*migrationpb.MigrationWorkflow, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MigrationWorkflowIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MigrationWorkflowIterator) Next() (*migrationpb.MigrationWorkflow, error) {
	var item *migrationpb.MigrationWorkflow
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MigrationWorkflowIterator) bufLen() int {
	return len(it.items)
}

func (it *MigrationWorkflowIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
