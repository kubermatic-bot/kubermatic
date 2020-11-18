// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// ListClusterRoleNamesV2Reader is a Reader for the ListClusterRoleNamesV2 structure.
type ListClusterRoleNamesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleNamesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleNamesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleNamesV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleNamesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleNamesV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleNamesV2OK creates a ListClusterRoleNamesV2OK with default headers values
func NewListClusterRoleNamesV2OK() *ListClusterRoleNamesV2OK {
	return &ListClusterRoleNamesV2OK{}
}

/*ListClusterRoleNamesV2OK handles this case with default header values.

ClusterRoleName
*/
type ListClusterRoleNamesV2OK struct {
	Payload []*models.ClusterRoleName
}

func (o *ListClusterRoleNamesV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames][%d] listClusterRoleNamesV2OK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleNamesV2OK) GetPayload() []*models.ClusterRoleName {
	return o.Payload
}

func (o *ListClusterRoleNamesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleNamesV2Unauthorized creates a ListClusterRoleNamesV2Unauthorized with default headers values
func NewListClusterRoleNamesV2Unauthorized() *ListClusterRoleNamesV2Unauthorized {
	return &ListClusterRoleNamesV2Unauthorized{}
}

/*ListClusterRoleNamesV2Unauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleNamesV2Unauthorized struct {
}

func (o *ListClusterRoleNamesV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames][%d] listClusterRoleNamesV2Unauthorized ", 401)
}

func (o *ListClusterRoleNamesV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleNamesV2Forbidden creates a ListClusterRoleNamesV2Forbidden with default headers values
func NewListClusterRoleNamesV2Forbidden() *ListClusterRoleNamesV2Forbidden {
	return &ListClusterRoleNamesV2Forbidden{}
}

/*ListClusterRoleNamesV2Forbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleNamesV2Forbidden struct {
}

func (o *ListClusterRoleNamesV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames][%d] listClusterRoleNamesV2Forbidden ", 403)
}

func (o *ListClusterRoleNamesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleNamesV2Default creates a ListClusterRoleNamesV2Default with default headers values
func NewListClusterRoleNamesV2Default(code int) *ListClusterRoleNamesV2Default {
	return &ListClusterRoleNamesV2Default{
		_statusCode: code,
	}
}

/*ListClusterRoleNamesV2Default handles this case with default header values.

errorResponse
*/
type ListClusterRoleNamesV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster role names v2 default response
func (o *ListClusterRoleNamesV2Default) Code() int {
	return o._statusCode
}

func (o *ListClusterRoleNamesV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames][%d] listClusterRoleNamesV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleNamesV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleNamesV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
