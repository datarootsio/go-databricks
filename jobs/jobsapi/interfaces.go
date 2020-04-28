package jobsapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/innovationnorway/go-databricks/jobs"
    "github.com/Azure/go-autorest/autorest"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            Create(ctx context.Context, body jobs.CreateAttributes) (result jobs.Attributes, err error)
            Delete(ctx context.Context, body jobs.Attributes) (result autorest.Response, err error)
            Get(ctx context.Context, body jobs.Attributes) (result jobs.GetResult, err error)
            List(ctx context.Context) (result jobs.ListResult, err error)
            Post(ctx context.Context, body jobs.ResetAttributes) (result autorest.Response, err error)
        }

        var _ BaseClientAPI = (*jobs.BaseClient)(nil)
