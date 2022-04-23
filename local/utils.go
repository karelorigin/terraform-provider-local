package local

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// wrapResourceCreate wraps `f` and automatically appends the returned diagnostics to the response
func wrapResourceCreate(resp *tfsdk.CreateResourceResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrapResourceRead wraps `f` and automatically appends the returned diagnostics to the response
func wrapResourceRead(resp *tfsdk.ReadResourceResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrapResourceUpdate wraps `f` and automatically appends the returned diagnostics to the response
func wrapResourceUpdate(resp *tfsdk.UpdateResourceResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrapResourceDelete wraps `f` and automatically appends the returned diagnostics to the response
func wrapResourceDelete(resp *tfsdk.DeleteResourceResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrapResourceImport wraps `f` and automatically appends the returned diagnostics to the response
func wrapResourceImport(resp *tfsdk.ImportResourceStateResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrapAttrValidate wraps `f` and automatically appends the returned diagnostics to the response
func wrapAttrValidate(resp *tfsdk.ValidateAttributeResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}

// wrappAttrModify wraps `f` and automatically appends the returned diagnostics to the response
func wrapAttrModify(resp *tfsdk.ModifyAttributePlanResponse, f func() interface{}) {
	err := f()
	if err == nil {
		return
	}

	switch v := err.(type) {
	case diag.Diagnostic:
		resp.Diagnostics = append(resp.Diagnostics, v)
	case diag.Diagnostics:
		resp.Diagnostics = append(resp.Diagnostics, v...)
	}
}
