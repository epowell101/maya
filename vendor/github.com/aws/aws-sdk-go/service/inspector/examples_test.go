// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package inspector_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleInspector_AddAttributesToFindings() {
	svc := inspector.New(session.New())

	params := &inspector.AddAttributesToFindingsInput{
		Attributes: []*inspector.Attribute{ // Required
			{ // Required
				Key:   aws.String("AttributeKey"),
				Value: aws.String("AttributeValue"),
			},
			// More values...
		},
		FindingArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.AddAttributesToFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_AttachAssessmentAndRulesPackage() {
	svc := inspector.New(session.New())

	params := &inspector.AttachAssessmentAndRulesPackageInput{
		AssessmentArn:   aws.String("Arn"), // Required
		RulesPackageArn: aws.String("Arn"), // Required
	}
	resp, err := svc.AttachAssessmentAndRulesPackage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateApplication() {
	svc := inspector.New(session.New())

	params := &inspector.CreateApplicationInput{
		ApplicationName:  aws.String("Name"), // Required
		ResourceGroupArn: aws.String("Arn"),  // Required
	}
	resp, err := svc.CreateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateAssessment() {
	svc := inspector.New(session.New())

	params := &inspector.CreateAssessmentInput{
		ApplicationArn:    aws.String("Arn"),  // Required
		AssessmentName:    aws.String("Name"), // Required
		DurationInSeconds: aws.Int64(1),       // Required
		UserAttributesForFindings: []*inspector.Attribute{
			{ // Required
				Key:   aws.String("AttributeKey"),
				Value: aws.String("AttributeValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssessment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_CreateResourceGroup() {
	svc := inspector.New(session.New())

	params := &inspector.CreateResourceGroupInput{
		ResourceGroupTags: aws.String("ResourceGroupTags"), // Required
	}
	resp, err := svc.CreateResourceGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteApplication() {
	svc := inspector.New(session.New())

	params := &inspector.DeleteApplicationInput{
		ApplicationArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteAssessment() {
	svc := inspector.New(session.New())

	params := &inspector.DeleteAssessmentInput{
		AssessmentArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteAssessment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DeleteRun() {
	svc := inspector.New(session.New())

	params := &inspector.DeleteRunInput{
		RunArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeApplication() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeApplicationInput{
		ApplicationArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeAssessment() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeAssessmentInput{
		AssessmentArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeAssessment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeCrossAccountAccessRole() {
	svc := inspector.New(session.New())

	var params *inspector.DescribeCrossAccountAccessRoleInput
	resp, err := svc.DescribeCrossAccountAccessRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeFinding() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeFindingInput{
		FindingArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeFinding(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeResourceGroup() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeResourceGroupInput{
		ResourceGroupArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeResourceGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeRulesPackage() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeRulesPackageInput{
		RulesPackageArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeRulesPackage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DescribeRun() {
	svc := inspector.New(session.New())

	params := &inspector.DescribeRunInput{
		RunArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeRun(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_DetachAssessmentAndRulesPackage() {
	svc := inspector.New(session.New())

	params := &inspector.DetachAssessmentAndRulesPackageInput{
		AssessmentArn:   aws.String("Arn"), // Required
		RulesPackageArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DetachAssessmentAndRulesPackage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_GetAssessmentTelemetry() {
	svc := inspector.New(session.New())

	params := &inspector.GetAssessmentTelemetryInput{
		AssessmentArn: aws.String("Arn"), // Required
	}
	resp, err := svc.GetAssessmentTelemetry(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListApplications() {
	svc := inspector.New(session.New())

	params := &inspector.ListApplicationsInput{
		Filter: &inspector.ApplicationsFilter{
			ApplicationNamePatterns: []*string{
				aws.String("NamePattern"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListApplications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessmentAgents() {
	svc := inspector.New(session.New())

	params := &inspector.ListAssessmentAgentsInput{
		AssessmentArn: aws.String("Arn"), // Required
		Filter: &inspector.AgentsFilter{
			AgentHealthList: []*string{
				aws.String("AgentHealth"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessmentAgents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAssessments() {
	svc := inspector.New(session.New())

	params := &inspector.ListAssessmentsInput{
		ApplicationArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
		Filter: &inspector.AssessmentsFilter{
			AssessmentNamePatterns: []*string{
				aws.String("NamePattern"), // Required
				// More values...
			},
			AssessmentStates: []*string{
				aws.String("AssessmentState"), // Required
				// More values...
			},
			DataCollected: aws.Bool(true),
			DurationRange: &inspector.DurationRange{
				Maximum: aws.Int64(1),
				Minimum: aws.Int64(1),
			},
			EndTimeRange: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
			StartTimeRange: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAssessments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAttachedAssessments() {
	svc := inspector.New(session.New())

	params := &inspector.ListAttachedAssessmentsInput{
		RulesPackageArn: aws.String("Arn"), // Required
		Filter: &inspector.AssessmentsFilter{
			AssessmentNamePatterns: []*string{
				aws.String("NamePattern"), // Required
				// More values...
			},
			AssessmentStates: []*string{
				aws.String("AssessmentState"), // Required
				// More values...
			},
			DataCollected: aws.Bool(true),
			DurationRange: &inspector.DurationRange{
				Maximum: aws.Int64(1),
				Minimum: aws.Int64(1),
			},
			EndTimeRange: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
			StartTimeRange: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListAttachedAssessments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListAttachedRulesPackages() {
	svc := inspector.New(session.New())

	params := &inspector.ListAttachedRulesPackagesInput{
		AssessmentArn: aws.String("Arn"), // Required
		MaxResults:    aws.Int64(1),
		NextToken:     aws.String("PaginationToken"),
	}
	resp, err := svc.ListAttachedRulesPackages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListFindings() {
	svc := inspector.New(session.New())

	params := &inspector.ListFindingsInput{
		Filter: &inspector.FindingsFilter{
			Attributes: []*inspector.Attribute{
				{ // Required
					Key:   aws.String("AttributeKey"),
					Value: aws.String("AttributeValue"),
				},
				// More values...
			},
			RuleNames: []*string{
				aws.String("Name"), // Required
				// More values...
			},
			RulesPackageArns: []*string{
				aws.String("Arn"), // Required
				// More values...
			},
			Severities: []*string{
				aws.String("Severity"), // Required
				// More values...
			},
			UserAttributes: []*inspector.Attribute{
				{ // Required
					Key:   aws.String("AttributeKey"),
					Value: aws.String("AttributeValue"),
				},
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
		RunArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.ListFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListRulesPackages() {
	svc := inspector.New(session.New())

	params := &inspector.ListRulesPackagesInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListRulesPackages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListRuns() {
	svc := inspector.New(session.New())

	params := &inspector.ListRunsInput{
		AssessmentArns: []*string{
			aws.String("Arn"), // Required
			// More values...
		},
		Filter: &inspector.RunsFilter{
			CompletionTime: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
			CreationTime: &inspector.TimestampRange{
				Maximum: aws.Time(time.Now()),
				Minimum: aws.Time(time.Now()),
			},
			RulesPackages: []*string{
				aws.String("Arn"), // Required
				// More values...
			},
			RunNamePatterns: []*string{
				aws.String("NamePattern"), // Required
				// More values...
			},
			RunStates: []*string{
				aws.String("RunState"), // Required
				// More values...
			},
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListRuns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_ListTagsForResource() {
	svc := inspector.New(session.New())

	params := &inspector.ListTagsForResourceInput{
		ResourceArn: aws.String("Arn"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_LocalizeText() {
	svc := inspector.New(session.New())

	params := &inspector.LocalizeTextInput{
		Locale: aws.String("Locale"), // Required
		LocalizedTexts: []*inspector.LocalizedText{ // Required
			{ // Required
				Key: &inspector.LocalizedTextKey{
					Facility: aws.String("LocalizedFacility"),
					Id:       aws.String("LocalizedTextId"),
				},
				Parameters: []*inspector.Parameter{
					{ // Required
						Name:  aws.String("ParameterName"),
						Value: aws.String("ParameterValue"),
					},
					// More values...
				},
			},
			// More values...
		},
	}
	resp, err := svc.LocalizeText(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_PreviewAgentsForResourceGroup() {
	svc := inspector.New(session.New())

	params := &inspector.PreviewAgentsForResourceGroupInput{
		ResourceGroupArn: aws.String("Arn"), // Required
		MaxResults:       aws.Int64(1),
		NextToken:        aws.String("PaginationToken"),
	}
	resp, err := svc.PreviewAgentsForResourceGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_RegisterCrossAccountAccessRole() {
	svc := inspector.New(session.New())

	params := &inspector.RegisterCrossAccountAccessRoleInput{
		RoleArn: aws.String("Arn"), // Required
	}
	resp, err := svc.RegisterCrossAccountAccessRole(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_RemoveAttributesFromFindings() {
	svc := inspector.New(session.New())

	params := &inspector.RemoveAttributesFromFindingsInput{
		AttributeKeys: []*string{ // Required
			aws.String("AttributeKey"), // Required
			// More values...
		},
		FindingArns: []*string{ // Required
			aws.String("Arn"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveAttributesFromFindings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_RunAssessment() {
	svc := inspector.New(session.New())

	params := &inspector.RunAssessmentInput{
		AssessmentArn: aws.String("Arn"),  // Required
		RunName:       aws.String("Name"), // Required
	}
	resp, err := svc.RunAssessment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_SetTagsForResource() {
	svc := inspector.New(session.New())

	params := &inspector.SetTagsForResourceInput{
		ResourceArn: aws.String("Arn"), // Required
		Tags: []*inspector.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.SetTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_StartDataCollection() {
	svc := inspector.New(session.New())

	params := &inspector.StartDataCollectionInput{
		AssessmentArn: aws.String("Arn"), // Required
	}
	resp, err := svc.StartDataCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_StopDataCollection() {
	svc := inspector.New(session.New())

	params := &inspector.StopDataCollectionInput{
		AssessmentArn: aws.String("Arn"), // Required
	}
	resp, err := svc.StopDataCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_UpdateApplication() {
	svc := inspector.New(session.New())

	params := &inspector.UpdateApplicationInput{
		ApplicationArn:   aws.String("Arn"),  // Required
		ApplicationName:  aws.String("Name"), // Required
		ResourceGroupArn: aws.String("Arn"),  // Required
	}
	resp, err := svc.UpdateApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleInspector_UpdateAssessment() {
	svc := inspector.New(session.New())

	params := &inspector.UpdateAssessmentInput{
		AssessmentArn:     aws.String("Arn"),  // Required
		AssessmentName:    aws.String("Name"), // Required
		DurationInSeconds: aws.Int64(1),       // Required
	}
	resp, err := svc.UpdateAssessment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
