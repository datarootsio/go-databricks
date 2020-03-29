package groups

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/innovationnorway/go-databricks/groups"

            // Attributes ...
            type Attributes struct {
            GroupName *string `json:"group_name,omitempty"`
            }

            // CreateResult ...
            type CreateResult struct {
            autorest.Response `json:"-"`
            GroupName *string `json:"group_name,omitempty"`
            }

            // DeleteAttributes ...
            type DeleteAttributes struct {
            GroupName *string `json:"group_name,omitempty"`
            }

            // ListMembersResult ...
            type ListMembersResult struct {
            autorest.Response `json:"-"`
            Members *[]PrincipalName `json:"members,omitempty"`
            }

            // ListResult ...
            type ListResult struct {
            autorest.Response `json:"-"`
            GroupNames *[]string `json:"group_names,omitempty"`
            }

            // MemberAttributes ...
            type MemberAttributes struct {
            UserName *string `json:"user_name,omitempty"`
            GroupName *string `json:"group_name,omitempty"`
            ParentName *string `json:"parent_name,omitempty"`
            }

            // PrincipalName ...
            type PrincipalName struct {
            UserName *string `json:"user_name,omitempty"`
            GroupName *string `json:"group_name,omitempty"`
            }

