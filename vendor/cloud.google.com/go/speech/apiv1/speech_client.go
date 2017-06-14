// Copyright 2017, Google Inc. All rights reserved.
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

// AUTO-GENERATED CODE. DO NOT EDIT.

package speech

import (
	"time"

	"cloud.google.com/go/internal/version"
	"cloud.google.com/go/longrunning"
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	Recognize            []gax.CallOption
	LongRunningRecognize []gax.CallOption
	StreamingRecognize   []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("speech.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
		{"default", "non_idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &CallOptions{
		Recognize:            retry[[2]string{"default", "idempotent"}],
		LongRunningRecognize: retry[[2]string{"default", "non_idempotent"}],
		StreamingRecognize:   retry[[2]string{"default", "idempotent"}],
	}
}

// Client is a client for interacting with Google Cloud Speech API.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client speechpb.SpeechClient

	// The call options for this service.
	CallOptions *CallOptions

	// The metadata to be sent with each request.
	xGoogHeader []string
}

// NewClient creates a new speech client.
//
// Service that implements Google Cloud Speech API.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		CallOptions: defaultCallOptions(),

		client: speechpb.NewSpeechClient(conn),
	}
	c.SetGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", version.Go()}, keyval...)
	kv = append(kv, "gapic", version.Repo, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeader = []string{gax.XGoogHeader(kv...)}
}

// Recognize performs synchronous speech recognition: receive results after all audio
// has been sent and processed.
func (c *Client) Recognize(ctx context.Context, req *speechpb.RecognizeRequest, opts ...gax.CallOption) (*speechpb.RecognizeResponse, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.Recognize[0:len(c.CallOptions.Recognize):len(c.CallOptions.Recognize)], opts...)
	var resp *speechpb.RecognizeResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.Recognize(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LongRunningRecognize performs asynchronous speech recognition: receive results via the
// google.longrunning.Operations interface. Returns either an
// `Operation.error` or an `Operation.response` which contains
// a `LongRunningRecognizeResponse` message.
func (c *Client) LongRunningRecognize(ctx context.Context, req *speechpb.LongRunningRecognizeRequest, opts ...gax.CallOption) (*LongRunningRecognizeOperation, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.LongRunningRecognize[0:len(c.CallOptions.LongRunningRecognize):len(c.CallOptions.LongRunningRecognize)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LongRunningRecognize(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &LongRunningRecognizeOperation{
		lro:         longrunning.InternalNewOperation(c.Connection(), resp),
		xGoogHeader: c.xGoogHeader,
	}, nil
}

// StreamingRecognize performs bidirectional streaming speech recognition: receive results while
// sending audio. This method is only available via the gRPC API (not REST).
func (c *Client) StreamingRecognize(ctx context.Context, opts ...gax.CallOption) (speechpb.Speech_StreamingRecognizeClient, error) {
	ctx = insertXGoog(ctx, c.xGoogHeader)
	opts = append(c.CallOptions.StreamingRecognize[0:len(c.CallOptions.StreamingRecognize):len(c.CallOptions.StreamingRecognize)], opts...)
	var resp speechpb.Speech_StreamingRecognizeClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StreamingRecognize(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LongRunningRecognizeOperation manages a long-running operation from LongRunningRecognize.
type LongRunningRecognizeOperation struct {
	lro *longrunning.Operation

	// The metadata to be sent with each request.
	xGoogHeader []string
}

// LongRunningRecognizeOperation returns a new LongRunningRecognizeOperation from a given name.
// The name must be that of a previously created LongRunningRecognizeOperation, possibly from a different process.
func (c *Client) LongRunningRecognizeOperation(name string) *LongRunningRecognizeOperation {
	return &LongRunningRecognizeOperation{
		lro:         longrunning.InternalNewOperation(c.Connection(), &longrunningpb.Operation{Name: name}),
		xGoogHeader: c.xGoogHeader,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *LongRunningRecognizeOperation) Wait(ctx context.Context) (*speechpb.LongRunningRecognizeResponse, error) {
	var resp speechpb.LongRunningRecognizeResponse
	ctx = insertXGoog(ctx, op.xGoogHeader)
	if err := op.lro.Wait(ctx, &resp); err != nil {
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
func (op *LongRunningRecognizeOperation) Poll(ctx context.Context) (*speechpb.LongRunningRecognizeResponse, error) {
	var resp speechpb.LongRunningRecognizeResponse
	ctx = insertXGoog(ctx, op.xGoogHeader)
	if err := op.lro.Poll(ctx, &resp); err != nil {
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
func (op *LongRunningRecognizeOperation) Metadata() (*speechpb.LongRunningRecognizeMetadata, error) {
	var meta speechpb.LongRunningRecognizeMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *LongRunningRecognizeOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *LongRunningRecognizeOperation) Name() string {
	return op.lro.Name()
}
