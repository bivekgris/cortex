// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alertgroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAlertGroupsParams creates a new GetAlertGroupsParams object
// with the default values initialized.
func NewGetAlertGroupsParams() GetAlertGroupsParams {

	var (
		// initialize parameters with default values

		activeDefault = bool(true)

		inhibitedDefault = bool(true)

		silencedDefault = bool(true)
	)

	return GetAlertGroupsParams{
		Active: &activeDefault,

		Inhibited: &inhibitedDefault,

		Silenced: &silencedDefault,
	}
}

// GetAlertGroupsParams contains all the bound params for the get alert groups operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAlertGroups
type GetAlertGroupsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Show active alerts
	  In: query
	  Default: true
	*/
	Active *bool
	/*A list of matchers to filter alerts by
	  In: query
	  Collection Format: multi
	*/
	Filter []string
	/*Show inhibited alerts
	  In: query
	  Default: true
	*/
	Inhibited *bool
	/*A regex matching receivers to filter alerts by
	  In: query
	*/
	Receiver *string
	/*Show silenced alerts
	  In: query
	  Default: true
	*/
	Silenced *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAlertGroupsParams() beforehand.
func (o *GetAlertGroupsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qActive, qhkActive, _ := qs.GetOK("active")
	if err := o.bindActive(qActive, qhkActive, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}

	qInhibited, qhkInhibited, _ := qs.GetOK("inhibited")
	if err := o.bindInhibited(qInhibited, qhkInhibited, route.Formats); err != nil {
		res = append(res, err)
	}

	qReceiver, qhkReceiver, _ := qs.GetOK("receiver")
	if err := o.bindReceiver(qReceiver, qhkReceiver, route.Formats); err != nil {
		res = append(res, err)
	}

	qSilenced, qhkSilenced, _ := qs.GetOK("silenced")
	if err := o.bindSilenced(qSilenced, qhkSilenced, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindActive binds and validates parameter Active from query.
func (o *GetAlertGroupsParams) bindActive(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAlertGroupsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("active", "query", "bool", raw)
	}
	o.Active = &value

	return nil
}

// bindFilter binds and validates array parameter Filter from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAlertGroupsParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	// CollectionFormat: multi
	filterIC := rawData
	if len(filterIC) == 0 {
		return nil
	}

	var filterIR []string
	for _, filterIV := range filterIC {
		filterI := filterIV

		filterIR = append(filterIR, filterI)
	}

	o.Filter = filterIR

	return nil
}

// bindInhibited binds and validates parameter Inhibited from query.
func (o *GetAlertGroupsParams) bindInhibited(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAlertGroupsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("inhibited", "query", "bool", raw)
	}
	o.Inhibited = &value

	return nil
}

// bindReceiver binds and validates parameter Receiver from query.
func (o *GetAlertGroupsParams) bindReceiver(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Receiver = &raw

	return nil
}

// bindSilenced binds and validates parameter Silenced from query.
func (o *GetAlertGroupsParams) bindSilenced(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAlertGroupsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("silenced", "query", "bool", raw)
	}
	o.Silenced = &value

	return nil
}