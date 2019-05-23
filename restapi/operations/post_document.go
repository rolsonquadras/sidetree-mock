// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostDocumentHandlerFunc turns a function with the right signature into a post document handler
type PostDocumentHandlerFunc func(PostDocumentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostDocumentHandlerFunc) Handle(params PostDocumentParams) middleware.Responder {
	return fn(params)
}

// PostDocumentHandler interface for that can handle valid post document params
type PostDocumentHandler interface {
	Handle(PostDocumentParams) middleware.Responder
}

// NewPostDocument creates a new http.Handler for the post document operation
func NewPostDocument(ctx *middleware.Context, handler PostDocumentHandler) *PostDocument {
	return &PostDocument{Context: ctx, Handler: handler}
}

/*PostDocument swagger:route POST /document postDocument

Create side tree did document

*/
type PostDocument struct {
	Context *middleware.Context
	Handler PostDocumentHandler
}

func (o *PostDocument) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostDocumentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}