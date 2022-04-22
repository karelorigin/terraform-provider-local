package local

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// defaultStringModifier is a plan modifier that sets a configurable default value for optional string attributes
type defaultStringModifier string

// Description returns a plain text description of the modifier's behavior
func (d defaultStringModifier) Description(_ context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to: %s", d)
}

// MarkdownDescription returns a markdown description of the modifier's behavior
func (d defaultStringModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

// Modify modifies the plan according to the modifier's documented behavior
func (d defaultStringModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	wrapAttrModify(resp, func() interface{} {
		var s types.String

		derrs := tfsdk.ValueAs(ctx, req.AttributeConfig, &s)
		if derrs.HasError() {
			return derrs
		}

		// Skip if value is already set
		if !s.Null {
			return nil
		}

		resp.AttributePlan = types.String{
			Value: string(d),
		}

		return nil
	})
}
