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

package resourcemanager

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newFoldersClientHook clientHook

// FoldersCallOptions contains the retry settings for each method of FoldersClient.
type FoldersCallOptions struct {
	ListFolders        []gax.CallOption
	SearchFolders      []gax.CallOption
	GetFolder          []gax.CallOption
	CreateFolder       []gax.CallOption
	UpdateFolder       []gax.CallOption
	MoveFolder         []gax.CallOption
	DeleteFolder       []gax.CallOption
	UndeleteFolder     []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultFoldersGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudresourcemanager.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudresourcemanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFoldersCallOptions() *FoldersCallOptions {
	return &FoldersCallOptions{
		ListFolders:        []gax.CallOption{},
		SearchFolders:      []gax.CallOption{},
		GetFolder:          []gax.CallOption{},
		CreateFolder:       []gax.CallOption{},
		UpdateFolder:       []gax.CallOption{},
		MoveFolder:         []gax.CallOption{},
		DeleteFolder:       []gax.CallOption{},
		UndeleteFolder:     []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
	}
}

// internalFoldersClient is an interface that defines the methods availaible from Cloud Resource Manager API.
type internalFoldersClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListFolders(context.Context, *resourcemanagerpb.ListFoldersRequest, ...gax.CallOption) *FolderIterator
	SearchFolders(context.Context, *resourcemanagerpb.SearchFoldersRequest, ...gax.CallOption) *FolderIterator
	GetFolder(context.Context, *resourcemanagerpb.GetFolderRequest, ...gax.CallOption) (*resourcemanagerpb.Folder, error)
	CreateFolder(context.Context, *resourcemanagerpb.CreateFolderRequest, ...gax.CallOption) (*CreateFolderOperation, error)
	CreateFolderOperation(name string) *CreateFolderOperation
	UpdateFolder(context.Context, *resourcemanagerpb.UpdateFolderRequest, ...gax.CallOption) (*resourcemanagerpb.Folder, error)
	MoveFolder(context.Context, *resourcemanagerpb.MoveFolderRequest, ...gax.CallOption) (*MoveFolderOperation, error)
	MoveFolderOperation(name string) *MoveFolderOperation
	DeleteFolder(context.Context, *resourcemanagerpb.DeleteFolderRequest, ...gax.CallOption) (*resourcemanagerpb.Folder, error)
	UndeleteFolder(context.Context, *resourcemanagerpb.UndeleteFolderRequest, ...gax.CallOption) (*resourcemanagerpb.Folder, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

// FoldersClient is a client for interacting with Cloud Resource Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages Cloud Resource Folders.
// Cloud Resource Folders can be used to organize the resources under an
// organization and to control the IAM policies applied to groups of resources.
type FoldersClient struct {
	// The internal transport-dependent client.
	internalClient internalFoldersClient

	// The call options for this service.
	CallOptions *FoldersCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FoldersClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FoldersClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FoldersClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListFolders lists the Folders that are direct descendants of supplied parent resource.
// List provides a strongly consistent view of the Folders underneath
// the specified parent resource.
// List returns Folders sorted based upon the (ascending) lexical ordering
// of their display_name.
// The caller must have resourcemanager.folders.list permission on the
// identified parent.
func (c *FoldersClient) ListFolders(ctx context.Context, req *resourcemanagerpb.ListFoldersRequest, opts ...gax.CallOption) *FolderIterator {
	return c.internalClient.ListFolders(ctx, req, opts...)
}

// SearchFolders search for folders that match specific filter criteria.
// Search provides an eventually consistent view of the folders a user has
// access to which meet the specified filter criteria.
//
// This will only return folders on which the caller has the
// permission resourcemanager.folders.get.
func (c *FoldersClient) SearchFolders(ctx context.Context, req *resourcemanagerpb.SearchFoldersRequest, opts ...gax.CallOption) *FolderIterator {
	return c.internalClient.SearchFolders(ctx, req, opts...)
}

// GetFolder retrieves a Folder identified by the supplied resource name.
// Valid Folder resource names have the format folders/{folder_id}
// (for example, folders/1234).
// The caller must have resourcemanager.folders.get permission on the
// identified folder.
func (c *FoldersClient) GetFolder(ctx context.Context, req *resourcemanagerpb.GetFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	return c.internalClient.GetFolder(ctx, req, opts...)
}

// CreateFolder creates a Folder in the resource hierarchy.
// Returns an Operation which can be used to track the progress of the
// folder creation workflow.
// Upon success the Operation.response field will be populated with the
// created Folder.
//
// In order to succeed, the addition of this new Folder must not violate
// the Folder naming, height or fanout constraints.
//
//   The Folder’s display_name must be distinct from all other Folder’s that
//   share its parent.
//
//   The addition of the Folder must not cause the active Folder hierarchy
//   to exceed a height of 4. Note, the full active + deleted Folder hierarchy
//   is allowed to reach a height of 8; this provides additional headroom when
//   moving folders that contain deleted folders.
//
//   The addition of the Folder must not cause the total number of Folders
//   under its parent to exceed 100.
//
// If the operation fails due to a folder constraint violation, some errors
// may be returned by the CreateFolder request, with status code
// FAILED_PRECONDITION and an error description. Other folder constraint
// violations will be communicated in the Operation, with the specific
// PreconditionFailure returned via the details list in the Operation.error
// field.
//
// The caller must have resourcemanager.folders.create permission on the
// identified parent.
func (c *FoldersClient) CreateFolder(ctx context.Context, req *resourcemanagerpb.CreateFolderRequest, opts ...gax.CallOption) (*CreateFolderOperation, error) {
	return c.internalClient.CreateFolder(ctx, req, opts...)
}

// CreateFolderOperation returns a new CreateFolderOperation from a given name.
// The name must be that of a previously created CreateFolderOperation, possibly from a different process.
func (c *FoldersClient) CreateFolderOperation(name string) *CreateFolderOperation {
	return c.internalClient.CreateFolderOperation(name)
}

// UpdateFolder updates a Folder, changing its display_name.
// Changes to the folder display_name will be rejected if they violate either
// the display_name formatting rules or naming constraints described in
// the CreateFolder documentation.
//
// The Folder’s display name must start and end with a letter or digit,
// may contain letters, digits, spaces, hyphens and underscores and can be
// no longer than 30 characters. This is captured by the regular expression:
// [\p{L}\p{N}]([\p{L}\p{N}_- ]{0,28}[\p{L}\p{N}])?.
// The caller must have resourcemanager.folders.update permission on the
// identified folder.
//
// If the update fails due to the unique name constraint then a
// PreconditionFailure explaining this violation will be returned
// in the Status.details field.
func (c *FoldersClient) UpdateFolder(ctx context.Context, req *resourcemanagerpb.UpdateFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	return c.internalClient.UpdateFolder(ctx, req, opts...)
}

// MoveFolder moves a Folder under a new resource parent.
// Returns an Operation which can be used to track the progress of the
// folder move workflow.
// Upon success the Operation.response field will be populated with the
// moved Folder.
// Upon failure, a FolderOperationError categorizing the failure cause will
// be returned - if the failure occurs synchronously then the
// FolderOperationError will be returned via the Status.details field
// and if it occurs asynchronously then the FolderOperation will be returned
// via the Operation.error field.
// In addition, the Operation.metadata field will be populated with a
// FolderOperation message as an aid to stateless clients.
// Folder moves will be rejected if they violate either the naming, height
// or fanout constraints described in the
// CreateFolder documentation.
// The caller must have resourcemanager.folders.move permission on the
// folder’s current and proposed new parent.
func (c *FoldersClient) MoveFolder(ctx context.Context, req *resourcemanagerpb.MoveFolderRequest, opts ...gax.CallOption) (*MoveFolderOperation, error) {
	return c.internalClient.MoveFolder(ctx, req, opts...)
}

// MoveFolderOperation returns a new MoveFolderOperation from a given name.
// The name must be that of a previously created MoveFolderOperation, possibly from a different process.
func (c *FoldersClient) MoveFolderOperation(name string) *MoveFolderOperation {
	return c.internalClient.MoveFolderOperation(name)
}

// DeleteFolder requests deletion of a Folder. The Folder is moved into the
// DELETE_REQUESTED state
// immediately, and is deleted approximately 30 days later. This method may
// only be called on an empty Folder in the
// [ACTIVE][google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE (at http://google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE)] state, where a Folder is empty if
// it doesn’t contain any Folders or Projects in the
// [ACTIVE][google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE (at http://google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE)] state.
// The caller must have resourcemanager.folders.delete permission on the
// identified folder.
func (c *FoldersClient) DeleteFolder(ctx context.Context, req *resourcemanagerpb.DeleteFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	return c.internalClient.DeleteFolder(ctx, req, opts...)
}

// UndeleteFolder cancels the deletion request for a Folder. This method may only be
// called on a Folder in the
// DELETE_REQUESTED state.
// In order to succeed, the Folder’s parent must be in the
// [ACTIVE][google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE (at http://google.cloud.resourcemanager.v2.Folder.LifecycleState.ACTIVE)] state.
// In addition, reintroducing the folder into the tree must not violate
// folder naming, height and fanout constraints described in the
// CreateFolder documentation.
// The caller must have resourcemanager.folders.undelete permission on the
// identified folder.
func (c *FoldersClient) UndeleteFolder(ctx context.Context, req *resourcemanagerpb.UndeleteFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	return c.internalClient.UndeleteFolder(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a Folder. The returned policy may be
// empty if no such policy or resource exists. The resource field should
// be the Folder’s resource name, e.g. “folders/1234”.
// The caller must have resourcemanager.folders.getIamPolicy permission
// on the identified folder.
func (c *FoldersClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on a Folder, replacing any existing policy.
// The resource field should be the Folder’s resource name, e.g.
// “folders/1234”.
// The caller must have resourcemanager.folders.setIamPolicy permission
// on the identified folder.
func (c *FoldersClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified Folder.
// The resource field should be the Folder’s resource name,
// e.g. “folders/1234”.
//
// There are no permissions required for making this API call.
func (c *FoldersClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// foldersGRPCClient is a client for interacting with Cloud Resource Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type foldersGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FoldersClient
	CallOptions **FoldersCallOptions

	// The gRPC API client.
	foldersClient resourcemanagerpb.FoldersClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFoldersClient creates a new folders client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages Cloud Resource Folders.
// Cloud Resource Folders can be used to organize the resources under an
// organization and to control the IAM policies applied to groups of resources.
func NewFoldersClient(ctx context.Context, opts ...option.ClientOption) (*FoldersClient, error) {
	clientOpts := defaultFoldersGRPCClientOptions()
	if newFoldersClientHook != nil {
		hookOpts, err := newFoldersClientHook(ctx, clientHookParams{})
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
	client := FoldersClient{CallOptions: defaultFoldersCallOptions()}

	c := &foldersGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		foldersClient:    resourcemanagerpb.NewFoldersClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *foldersGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *foldersGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *foldersGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *foldersGRPCClient) ListFolders(ctx context.Context, req *resourcemanagerpb.ListFoldersRequest, opts ...gax.CallOption) *FolderIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListFolders[0:len((*c.CallOptions).ListFolders):len((*c.CallOptions).ListFolders)], opts...)
	it := &FolderIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListFoldersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.Folder, string, error) {
		var resp *resourcemanagerpb.ListFoldersResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.foldersClient.ListFolders(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetFolders(), resp.GetNextPageToken(), nil
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

func (c *foldersGRPCClient) SearchFolders(ctx context.Context, req *resourcemanagerpb.SearchFoldersRequest, opts ...gax.CallOption) *FolderIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchFolders[0:len((*c.CallOptions).SearchFolders):len((*c.CallOptions).SearchFolders)], opts...)
	it := &FolderIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.SearchFoldersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.Folder, string, error) {
		var resp *resourcemanagerpb.SearchFoldersResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.foldersClient.SearchFolders(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetFolders(), resp.GetNextPageToken(), nil
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

func (c *foldersGRPCClient) GetFolder(ctx context.Context, req *resourcemanagerpb.GetFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFolder[0:len((*c.CallOptions).GetFolder):len((*c.CallOptions).GetFolder)], opts...)
	var resp *resourcemanagerpb.Folder
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.GetFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) CreateFolder(ctx context.Context, req *resourcemanagerpb.CreateFolderRequest, opts ...gax.CallOption) (*CreateFolderOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateFolder[0:len((*c.CallOptions).CreateFolder):len((*c.CallOptions).CreateFolder)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.CreateFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateFolderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *foldersGRPCClient) UpdateFolder(ctx context.Context, req *resourcemanagerpb.UpdateFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "folder.name", url.QueryEscape(req.GetFolder().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateFolder[0:len((*c.CallOptions).UpdateFolder):len((*c.CallOptions).UpdateFolder)], opts...)
	var resp *resourcemanagerpb.Folder
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.UpdateFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) MoveFolder(ctx context.Context, req *resourcemanagerpb.MoveFolderRequest, opts ...gax.CallOption) (*MoveFolderOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MoveFolder[0:len((*c.CallOptions).MoveFolder):len((*c.CallOptions).MoveFolder)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.MoveFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &MoveFolderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *foldersGRPCClient) DeleteFolder(ctx context.Context, req *resourcemanagerpb.DeleteFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteFolder[0:len((*c.CallOptions).DeleteFolder):len((*c.CallOptions).DeleteFolder)], opts...)
	var resp *resourcemanagerpb.Folder
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.DeleteFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) UndeleteFolder(ctx context.Context, req *resourcemanagerpb.UndeleteFolderRequest, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UndeleteFolder[0:len((*c.CallOptions).UndeleteFolder):len((*c.CallOptions).UndeleteFolder)], opts...)
	var resp *resourcemanagerpb.Folder
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.UndeleteFolder(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *foldersGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.foldersClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateFolderOperation manages a long-running operation from CreateFolder.
type CreateFolderOperation struct {
	lro *longrunning.Operation
}

// CreateFolderOperation returns a new CreateFolderOperation from a given name.
// The name must be that of a previously created CreateFolderOperation, possibly from a different process.
func (c *foldersGRPCClient) CreateFolderOperation(name string) *CreateFolderOperation {
	return &CreateFolderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateFolderOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	var resp resourcemanagerpb.Folder
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateFolderOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	var resp resourcemanagerpb.Folder
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateFolderOperation) Metadata() (*resourcemanagerpb.FolderOperation, error) {
	var meta resourcemanagerpb.FolderOperation
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateFolderOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateFolderOperation) Name() string {
	return op.lro.Name()
}

// MoveFolderOperation manages a long-running operation from MoveFolder.
type MoveFolderOperation struct {
	lro *longrunning.Operation
}

// MoveFolderOperation returns a new MoveFolderOperation from a given name.
// The name must be that of a previously created MoveFolderOperation, possibly from a different process.
func (c *foldersGRPCClient) MoveFolderOperation(name string) *MoveFolderOperation {
	return &MoveFolderOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *MoveFolderOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	var resp resourcemanagerpb.Folder
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *MoveFolderOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.Folder, error) {
	var resp resourcemanagerpb.Folder
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *MoveFolderOperation) Metadata() (*resourcemanagerpb.FolderOperation, error) {
	var meta resourcemanagerpb.FolderOperation
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *MoveFolderOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *MoveFolderOperation) Name() string {
	return op.lro.Name()
}

// FolderIterator manages a stream of *resourcemanagerpb.Folder.
type FolderIterator struct {
	items    []*resourcemanagerpb.Folder
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
	InternalFetch func(pageSize int, pageToken string) (results []*resourcemanagerpb.Folder, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *FolderIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FolderIterator) Next() (*resourcemanagerpb.Folder, error) {
	var item *resourcemanagerpb.Folder
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FolderIterator) bufLen() int {
	return len(it.items)
}

func (it *FolderIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
