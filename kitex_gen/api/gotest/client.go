// Code generated by Kitex v0.0.8. DO NOT EDIT.

package gotest

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"gotest/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Add(ctx context.Context, req *api.AddRequest, callOptions ...callopt.Option) (r *api.AddResponse, err error)
	Delete(ctx context.Context, req *api.DeleteRequest, callOptions ...callopt.Option) (r *api.DeleteResponse, err error)
	Select(ctx context.Context, req *api.SelectRequest, callOptions ...callopt.Option) (r *api.SelectResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kGotestClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kGotestClient struct {
	*kClient
}

func (p *kGotestClient) Add(ctx context.Context, req *api.AddRequest, callOptions ...callopt.Option) (r *api.AddResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Add(ctx, req)
}

func (p *kGotestClient) Delete(ctx context.Context, req *api.DeleteRequest, callOptions ...callopt.Option) (r *api.DeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Delete(ctx, req)
}

func (p *kGotestClient) Select(ctx context.Context, req *api.SelectRequest, callOptions ...callopt.Option) (r *api.SelectResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Select(ctx, req)
}
