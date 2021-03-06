// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudwatchMetricAlarmInvalidUnitRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	unit = "GB"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchMetricAlarmInvalidUnitRule(),
					Message: `unit is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	unit = "Gigabytes"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchMetricAlarmInvalidUnitRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
