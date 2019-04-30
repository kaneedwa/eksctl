package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2TransitGatewayRouteTableAssociation AWS CloudFormation Resource (AWS::EC2::TransitGatewayRouteTableAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html
type AWSEC2TransitGatewayRouteTableAssociation struct {

	// TransitGatewayAttachmentId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayattachmentid
	TransitGatewayAttachmentId *Value `json:"TransitGatewayAttachmentId,omitempty"`

	// TransitGatewayRouteTableId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html#cfn-ec2-transitgatewayroutetableassociation-transitgatewayroutetableid
	TransitGatewayRouteTableId *Value `json:"TransitGatewayRouteTableId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2TransitGatewayRouteTableAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::TransitGatewayRouteTableAssociation"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2TransitGatewayRouteTableAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2TransitGatewayRouteTableAssociation
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2TransitGatewayRouteTableAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2TransitGatewayRouteTableAssociation
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEC2TransitGatewayRouteTableAssociation(*res.Properties)
	}

	return nil
}

// GetAllAWSEC2TransitGatewayRouteTableAssociationResources retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2TransitGatewayRouteTableAssociationResources() map[string]AWSEC2TransitGatewayRouteTableAssociation {
	results := map[string]AWSEC2TransitGatewayRouteTableAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2TransitGatewayRouteTableAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::TransitGatewayRouteTableAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2TransitGatewayRouteTableAssociation{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2TransitGatewayRouteTableAssociationWithName retrieves all AWSEC2TransitGatewayRouteTableAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2TransitGatewayRouteTableAssociationWithName(name string) (AWSEC2TransitGatewayRouteTableAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2TransitGatewayRouteTableAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::TransitGatewayRouteTableAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2TransitGatewayRouteTableAssociation{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2TransitGatewayRouteTableAssociation{}, errors.New("resource not found")
}