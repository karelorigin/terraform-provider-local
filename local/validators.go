package local

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// validateFilePermissions satisfies the tfsdk.AttributeValidator interface
type validateFilePermissions struct {
	skipNull bool
}

// Description returns the validator's description
func (v validateFilePermissions) Description(_ context.Context) string {
	return "file permissions must be in numeric notation"
}

// MarkdownDescription returns the validator's description in markdown format
func (v validateFilePermissions) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate validates the attribute
func (v validateFilePermissions) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	wrapAttrValidate(resp, func() interface{} {
		var s types.String

		derrs := tfsdk.ValueAs(ctx, req.AttributeConfig, &s)
		if derrs.HasError() {
			return derrs
		}

		// Skip under certain conditions
		if s.Unknown || (s.Null && v.skipNull) {
			return nil
		}

		// Null values cannot exist
		if s.Null {
			return diag.NewAttributeErrorDiagnostic(req.AttributePath, "invalid value", "file permissions cannot be null")
		}

		// Value must be 3 or 4 characters long
		if len(s.Value) < 3 || len(s.Value) > 4 {
			return diag.NewAttributeErrorDiagnostic(req.AttributePath, "invalid value", "string length should be 3 or 4 digits")
		}

		mode, err := strconv.ParseInt(s.Value, 8, 64)
		if err != nil || mode > 0777 || mode < 0 {
			return diag.NewAttributeErrorDiagnostic(req.AttributePath, "invalid value", "string must be expressed in octal numeric notation")
		}

		return nil
	})
}
