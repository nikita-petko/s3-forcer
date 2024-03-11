// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a statement to a topic's access control policy, granting access for the
// specified Amazon Web Services accounts to the specified actions. To remove the
// ability to change topic permissions, you must deny permissions to the
// AddPermission , RemovePermission , and SetTopicAttributes actions in your IAM
// policy.
func (c *Client) AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) {
	if params == nil {
		params = &AddPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPermission", params, optFns, c.addOperationAddPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddPermissionInput struct {

	// The Amazon Web Services account IDs of the users (principals) who will be given
	// access to the specified actions. The users must have Amazon Web Services
	// account, but do not need to be signed up for this service.
	//
	// This member is required.
	AWSAccountId []string

	// The action you want to allow for the specified principal(s). Valid values: Any
	// Amazon SNS action name, for example Publish .
	//
	// This member is required.
	ActionName []string

	// A unique identifier for the new policy statement.
	//
	// This member is required.
	Label *string

	// The ARN of the topic whose access control policy you wish to modify.
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

type AddPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddPermission"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAddPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPermission(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAddPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddPermission",
	}
}
