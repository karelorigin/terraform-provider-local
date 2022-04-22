package local

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LocalFile is a local file model
type LocalFile struct {
	Path        types.String `tfsdk:"path"`
	Content     types.String `tfsdk:"content"`
	Permissions types.String `tfsdk:"permissions"`
}
