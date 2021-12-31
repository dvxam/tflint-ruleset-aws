// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretPolicyInvalidPolicyRule checks the pattern is valid
type AwsSecretsmanagerSecretPolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretPolicyInvalidPolicyRule() *AwsSecretsmanagerSecretPolicyInvalidPolicyRule {
	return &AwsSecretsmanagerSecretPolicyInvalidPolicyRule{
		resourceType:  "aws_secretsmanager_secret_policy",
		attributeName: "policy",
		max:           20480,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Name() string {
	return "aws_secretsmanager_secret_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 20480 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
