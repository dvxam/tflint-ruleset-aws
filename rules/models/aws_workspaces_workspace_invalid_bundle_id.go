// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorkspacesWorkspaceInvalidBundleIDRule checks the pattern is valid
type AwsWorkspacesWorkspaceInvalidBundleIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsWorkspacesWorkspaceInvalidBundleIDRule returns new rule with default attributes
func NewAwsWorkspacesWorkspaceInvalidBundleIDRule() *AwsWorkspacesWorkspaceInvalidBundleIDRule {
	return &AwsWorkspacesWorkspaceInvalidBundleIDRule{
		resourceType:  "aws_workspaces_workspace",
		attributeName: "bundle_id",
		pattern:       regexp.MustCompile(`^wsb-[0-9a-z]{8,63}$`),
	}
}

// Name returns the rule name
func (r *AwsWorkspacesWorkspaceInvalidBundleIDRule) Name() string {
	return "aws_workspaces_workspace_invalid_bundle_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorkspacesWorkspaceInvalidBundleIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorkspacesWorkspaceInvalidBundleIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorkspacesWorkspaceInvalidBundleIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorkspacesWorkspaceInvalidBundleIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^wsb-[0-9a-z]{8,63}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}