// Code generated by zenrpc; DO NOT EDIT.

package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/semrush/zenrpc/v2"
	"github.com/semrush/zenrpc/v2/smd"
)

var RPC = struct {
	ArithService struct{ Sum, Multiply, Divide, Pow string }
	GreetService struct{ Greet string }
}{
	ArithService: struct{ Sum, Multiply, Divide, Pow string }{
		Sum:      "sum",
		Multiply: "multiply",
		Divide:   "divide",
		Pow:      "pow",
	},
	GreetService: struct{ Greet string }{
		Greet: "greet",
	},
}

func (ArithService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Sum": {
				Description: `Sum sums two digits and returns error with error code as result and IP from context.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Boolean,
				},
			},
			"Multiply": {
				Description: `Multiply multiples two digits and returns result.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Integer,
				},
			},
			"Divide": {
				Description: ``,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Optional:    false,
						Description: ``,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"Quo": {
							Description: ``,
							Type:        smd.Integer,
						},
						"Rem": {
							Description: ``,
							Type:        smd.Integer,
						},
					},
				},
			},
			"Pow": {
				Description: `Pow returns x**y, the base-x exponential of y. If Exp is not set then default value is 2.
`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "base",
						Optional:    false,
						Description: ``,
						Type:        smd.Float,
					},
					{
						Name:        "exp",
						Optional:    true,
						Description: ``,
						Type:        smd.Float,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.Float,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s ArithService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.ArithService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Sum(ctx, args.A, args.B))

	case RPC.ArithService.Multiply:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Multiply(args.A, args.B))

	case RPC.ArithService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Divide(args.A, args.B))

	case RPC.ArithService.Pow:
		var args = struct {
			Base float64  `json:"base"`
			Exp  *float64 `json:"exp"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"base", "exp"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:exp=2
		if args.Exp == nil {
			var v float64 = 2
			args.Exp = &v
		}

		resp.Set(s.Pow(args.Base, *args.Exp))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (GreetService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Greet": {
				Description: `Greet`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "name",
						Optional:    false,
						Description: ``,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: ``,
					Optional:    false,
					Type:        smd.String,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s GreetService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.GreetService.Greet:
		var args = struct {
			Name string `json:"name"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"name"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Greet(ctx, args.Name))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
