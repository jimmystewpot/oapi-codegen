//go:build go1.22

// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version (devel) DO NOT EDIT.
package main

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
)

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Id Unique id of the pet
	Id int64 `json:"id"`

	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// Tags tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`

	// Limit maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = NewPet

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams)
	// Creates a new pet
	// (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(w http.ResponseWriter, r *http.Request, id int64)
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(w http.ResponseWriter, r *http.Request, id int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// FindPets operation middleware
func (siw *ServerInterfaceWrapper) FindPets(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams

	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tags", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPets(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPet(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeletePet operation middleware
func (siw *ServerInterfaceWrapper) DeletePet(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePet(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// FindPetByID operation middleware
func (siw *ServerInterfaceWrapper) FindPetByID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindPetByID(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/pets", wrapper.FindPets)
	m.HandleFunc("POST "+options.BaseURL+"/pets", wrapper.AddPet)
	m.HandleFunc("DELETE "+options.BaseURL+"/pets/{id}", wrapper.DeletePet)
	m.HandleFunc("GET "+options.BaseURL+"/pets/{id}", wrapper.FindPetByID)

	return m
}
