// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetV1ClustersHandlerFunc turns a function with the right signature into a get v1 clusters handler
type GetV1ClustersHandlerFunc func(GetV1ClustersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1ClustersHandlerFunc) Handle(params GetV1ClustersParams) middleware.Responder {
	return fn(params)
}

// GetV1ClustersHandler interface for that can handle valid get v1 clusters params
type GetV1ClustersHandler interface {
	Handle(GetV1ClustersParams) middleware.Responder
}

// NewGetV1Clusters creates a new http.Handler for the get v1 clusters operation
func NewGetV1Clusters(ctx *middleware.Context, handler GetV1ClustersHandler) *GetV1Clusters {
	return &GetV1Clusters{Context: ctx, Handler: handler}
}

/*GetV1Clusters swagger:route GET /v1/clusters clusters getV1Clusters

Gets some clusters

Returns a list containing all clusters.

*/
type GetV1Clusters struct {
	Context *middleware.Context
	Handler GetV1ClustersHandler
}

func (o *GetV1Clusters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1ClustersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}