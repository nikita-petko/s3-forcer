// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the calling Amazon Web Services account's dedicated origination numbers
// and their metadata. For more information about origination numbers, see
// Origination numbers (https://docs.aws.amazon.com/sns/latest/dg/channels-sms-originating-identities-origination-numbers.html)
// in the Amazon SNS Developer Guide.
func (c *Client) ListOriginationNumbers(ctx context.Context, params *ListOriginationNumbersInput, optFns ...func(*Options)) (*ListOriginationNumbersOutput, error) {
	if params == nil {
		params = &ListOriginationNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOriginationNumbers", params, optFns, c.addOperationListOriginationNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOriginationNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOriginationNumbersInput struct {

	// The maximum number of origination numbers to return.
	MaxResults *int32

	// Token that the previous ListOriginationNumbers request returns.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOriginationNumbersOutput struct {

	// A NextToken string is returned when you call the ListOriginationNumbers
	// operation if additional pages of records are available.
	NextToken *string

	// A list of the calling account's verified and pending origination numbers.
	PhoneNumbers []types.PhoneNumberInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOriginationNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListOriginationNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListOriginationNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOriginationNumbers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

// ListOriginationNumbersAPIClient is a client that implements the
// ListOriginationNumbers operation.
type ListOriginationNumbersAPIClient interface {
	ListOriginationNumbers(context.Context, *ListOriginationNumbersInput, ...func(*Options)) (*ListOriginationNumbersOutput, error)
}

var _ ListOriginationNumbersAPIClient = (*Client)(nil)

// ListOriginationNumbersPaginatorOptions is the paginator options for
// ListOriginationNumbers
type ListOriginationNumbersPaginatorOptions struct {
	// The maximum number of origination numbers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOriginationNumbersPaginator is a paginator for ListOriginationNumbers
type ListOriginationNumbersPaginator struct {
	options   ListOriginationNumbersPaginatorOptions
	client    ListOriginationNumbersAPIClient
	params    *ListOriginationNumbersInput
	nextToken *string
	firstPage bool
}

// NewListOriginationNumbersPaginator returns a new ListOriginationNumbersPaginator
func NewListOriginationNumbersPaginator(client ListOriginationNumbersAPIClient, params *ListOriginationNumbersInput, optFns ...func(*ListOriginationNumbersPaginatorOptions)) *ListOriginationNumbersPaginator {
	if params == nil {
		params = &ListOriginationNumbersInput{}
	}

	options := ListOriginationNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOriginationNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOriginationNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOriginationNumbers page.
func (p *ListOriginationNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOriginationNumbersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListOriginationNumbers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListOriginationNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListOriginationNumbers",
	}
}
