// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIotPolicyInvalidPolicyRule checks the pattern is valid
type AwsIotPolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsIotPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsIotPolicyInvalidPolicyRule() *AwsIotPolicyInvalidPolicyRule {
	return &AwsIotPolicyInvalidPolicyRule{
		resourceType:  "aws_iot_policy",
		attributeName: "policy",
		max:           404600,
		pattern:       regexp.MustCompile(`^[\s\S]*$`),
	}
}

// Name returns the rule name
func (r *AwsIotPolicyInvalidPolicyRule) Name() string {
	return "aws_iot_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIotPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIotPolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIotPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIotPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 404600 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\s\S]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}