package noderoommember

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
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
