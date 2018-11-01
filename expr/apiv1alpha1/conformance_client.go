// Copyright 2018 Google LLC
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

// AUTO-GENERATED CODE. DO NOT EDIT.

package expr

import (
	"cloud.google.com/go/internal/version"
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ConformanceCallOptions contains the retry settings for each method of ConformanceClient.
type ConformanceCallOptions struct {
	Parse []gax.CallOption
	Check []gax.CallOption
	Eval  []gax.CallOption
}

func defaultConformanceClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("cel.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultConformanceCallOptions() *ConformanceCallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &ConformanceCallOptions{
		Parse: retry[[2]string{"default", "non_idempotent"}],
		Check: retry[[2]string{"default", "non_idempotent"}],
		Eval:  retry[[2]string{"default", "non_idempotent"}],
	}
}

// ConformanceClient is a client for interacting with Common Expression Language.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ConformanceClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	conformanceClient exprpb.ConformanceServiceClient

	// The call options for this service.
	CallOptions *ConformanceCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConformanceClient creates a new conformance service client.
//
// Access a CEL implementation from another process or machine.
// A CEL implementation is decomposed as a parser, a static checker,
// and an evaluator.  Every CEL implementation is expected to provide
// a server for this API.  The API will be used for conformance testing
// and other utilities.
func NewConformanceClient(ctx context.Context, opts ...option.ClientOption) (*ConformanceClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultConformanceClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ConformanceClient{
		conn:        conn,
		CallOptions: defaultConformanceCallOptions(),

		conformanceClient: exprpb.NewConformanceServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *ConformanceClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConformanceClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConformanceClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", version.Go()}, keyval...)
	kv = append(kv, "gapic", version.Repo, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Parse transforms CEL source text into a parsed representation.
func (c *ConformanceClient) Parse(ctx context.Context, req *exprpb.ParseRequest, opts ...gax.CallOption) (*exprpb.ParseResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.Parse[0:len(c.CallOptions.Parse):len(c.CallOptions.Parse)], opts...)
	var resp *exprpb.ParseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conformanceClient.Parse(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Check runs static checks on a parsed CEL representation and return
// an annotated representation, or a set of issues.
func (c *ConformanceClient) Check(ctx context.Context, req *exprpb.CheckRequest, opts ...gax.CallOption) (*exprpb.CheckResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.Check[0:len(c.CallOptions.Check):len(c.CallOptions.Check)], opts...)
	var resp *exprpb.CheckResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conformanceClient.Check(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Eval evaluates a parsed or annotation CEL representation given
// values of external bindings.
func (c *ConformanceClient) Eval(ctx context.Context, req *exprpb.EvalRequest, opts ...gax.CallOption) (*exprpb.EvalResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.Eval[0:len(c.CallOptions.Eval):len(c.CallOptions.Eval)], opts...)
	var resp *exprpb.EvalResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conformanceClient.Eval(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}