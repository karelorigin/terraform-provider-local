package local

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetSchema returns the resource type's schema
func (r resourceLocalStickyFileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Generates a local file with given content and permissions. Unlike hashicorp/local.local_file, this resource " +
			"makes sure that the file will always exist, even in environments such as Terraform Cloud, where disk state is not preserved.",
		Attributes: map[string]tfsdk.Attribute{
			"path": {
				Type:        types.StringType,
				Required:    true,
				Description: "The path where the file will be created.",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"content": {
				Type:        types.StringType,
				Required:    true,
				Description: "The content of the file.",
			},
			"permissions": {
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The file permissions in numerc notation.",
				Validators: []tfsdk.AttributeValidator{
					validateFilePermissions{
						skipNull: true,
					},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					defaultStringModifier("0666"),
				},
			},
		},
	}, nil
}

// NewResource returns a new local file resource instance
func (r resourceLocalStickyFileType) NewResource(_ context.Context, _ tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceLocalStickyFileType{}, nil
}

// Create creates a new resource
func (r resourceLocalStickyFileType) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	wrapResourceCreate(resp, func() interface{} {
		var file LocalFile

		derrs := req.Plan.Get(ctx, &file)
		if derrs != nil {
			return derrs
		}

		derr := r.write(file)
		if derr != nil {
			return derr
		}

		derrs = resp.State.Set(ctx, &file)
		if derrs != nil {
			return derrs
		}

		return nil
	})
}

// Read reads resource state
func (r resourceLocalStickyFileType) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	wrapResourceRead(resp, func() interface{} {
		var file LocalFile

		derrs := req.State.Get(ctx, &file)
		if derrs != nil {
			return derrs
		}

		_, err := os.Stat(file.Path.Value)
		if !(err == nil || errors.Is(err, os.ErrNotExist)) {
			return diag.NewErrorDiagnostic("could not verify whether path exists or not", err.Error())
		}

		// File already exists, skip
		if !errors.Is(err, os.ErrNotExist) {
			return nil
		}

		derr := r.write(file)
		if err != nil {
			return derr
		}

		return nil
	})
}

// Update updates the resource
func (r resourceLocalStickyFileType) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	wrapResourceUpdate(resp, func() interface{} {
		var file LocalFile

		derrs := req.Plan.Get(ctx, &file)
		if derrs != nil {
			return derrs
		}

		derr := r.write(file)
		if derr != nil {
			return derr
		}

		derrs = resp.State.Set(ctx, &file)
		if derrs != nil {
			return derrs
		}

		return nil
	})
}

// Delete deletes the resource
func (r resourceLocalStickyFileType) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	wrapResourceDelete(resp, func() interface{} {
		var file LocalFile

		derrs := req.State.Get(ctx, &file)
		if derrs != nil {
			return derrs
		}

		_, err := os.Stat(file.Path.Value)
		if err != nil {
			return diag.NewErrorDiagnostic("could not verify whether path exists or not", err.Error())
		}

		// Delete file if it exists
		if !errors.Is(err, os.ErrNotExist) {
			if err := os.Remove(file.Path.Value); err != nil {
				return diag.NewErrorDiagnostic("unable to remove file from disk", err.Error())
			}
		}

		resp.State.RemoveResource(ctx)

		return nil
	})
}

// Import imports the resource
func (r resourceLocalStickyFileType) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
}

// write writes the file object to disk
func (r resourceLocalStickyFileType) write(file LocalFile) diag.Diagnostic {
	f, err := os.Create(file.Path.Value)
	if err != nil {
		return diag.NewErrorDiagnostic("error while creating file", err.Error())
	}

	defer f.Close()

	perms, err := strconv.ParseUint(file.Permissions.Value, 8, 32)
	if err != nil {
		return diag.NewErrorDiagnostic("error while parsing file permissions", err.Error())
	}

	// Set file permissions
	f.Chmod(fs.FileMode(perms))

	// Write string to file
	_, err = f.WriteString(file.Content.Value)
	if err != nil {
		return diag.NewErrorDiagnostic("error while writing file content", err.Error())
	}

	return nil
}
