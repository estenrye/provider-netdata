package noderoommember

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const (
	ErrFmtNoAttribute    = "required attribute %s not found"
	ErrFmtUnexpectedType = "unexpected type for attribute %s, expected string"
)

func getNameFromFullyQualifiedID(tfstate map[string]any) (string, error) {
	spaceID, ok := tfstate["space_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "space_id")
	}
	spaceIDStr, ok := spaceID.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "space_id")
	}
	roomID, ok := tfstate["room_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "room_id")
	}
	roomIDStr, ok := roomID.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "room_id")
	}
	return fmt.Sprintf("%s,%s", spaceIDStr, roomIDStr), nil
}

func getFullyQualifiedIDfunc(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	spaceID, ok := parameters["space_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "space_id")
	}
	spaceIDStr, ok := spaceID.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "space_id")
	}

	roomID, ok := parameters["room_id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "room_id")
	}
	roomIDStr, ok := roomID.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "room_id")
	}

	return fmt.Sprintf("%s,%s", spaceIDStr, roomIDStr), nil
}

// NewExternalName returns a new ExternalName for the node room member resource.
func NewExternalName() config.ExternalName {
	externalName := config.NewExternalNameFrom(
		config.TemplatedStringAsIdentifier("", "{{ .parameters.space_id }},{{ .parameters.room_id }}"),
	)
	externalName.SetIdentifierArgumentFn = config.IdentifierFromProvider.SetIdentifierArgumentFn
	externalName.GetIDFn = getFullyQualifiedIDfunc
	externalName.GetExternalNameFn = getNameFromFullyQualifiedID

	return externalName
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_node_room_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "noderoommember"

		r.ExternalName = NewExternalName()

		r.TerraformResource.Schema["rule"].Type = schema.TypeList
		r.TerraformResource.Schema["rule"].Elem = &schema.Resource{

			Schema: map[string]*schema.Schema{
				"action": {Type: schema.TypeString, Required: true, ForceNew: false},
				"clause": {
					Type:     schema.TypeList,
					Required: true,
					ForceNew: false,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"label": {
								Type:     schema.TypeString,
								Required: true,
								ForceNew: false,
							},
							"negate": {
								Type:     schema.TypeBool,
								Optional: true,
								Default:  false,
								ForceNew: false,
							},
							"operator": {
								Type:     schema.TypeString,
								Required: true,
								ForceNew: false,
							},
							"value": {
								Type:     schema.TypeString,
								Required: true,
								ForceNew: false,
							},
						},
					},
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: false,
				},
				"id": {
					Type:     schema.TypeString,
					Computed: true,
					Required: true,
					ForceNew: false,
					// This is a computed field, so we don't need to set it in the schema
				},
			},
		}

		r.References["space_id"] = config.Reference{
			TerraformName: "netdata_space",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}

		r.References["room_id"] = config.Reference{
			TerraformName: "netdata_room",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
