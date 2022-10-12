/*
Blockmate

Blockmate API OpenAPI documentation

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blockmate

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// AddressNameAndCategoryInfoApiService AddressNameAndCategoryInfoApi service
type AddressNameAndCategoryInfoApiService service

type ApiGetAddressNameInfoRequest struct {
	ctx context.Context
	ApiService *AddressNameAndCategoryInfoApiService
	address *string
	chain *string
}

// Address for wich name and category should be returned
func (r ApiGetAddressNameInfoRequest) Address(address string) ApiGetAddressNameInfoRequest {
	r.address = &address
	return r
}

// Blockchain identifier
func (r ApiGetAddressNameInfoRequest) Chain(chain string) ApiGetAddressNameInfoRequest {
	r.chain = &chain
	return r
}

func (r ApiGetAddressNameInfoRequest) Execute() (*GetAddressNameInfo200Response, *http.Response, error) {
	return r.ApiService.GetAddressNameInfoExecute(r)
}

/*
GetAddressNameInfo Get address name and category info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAddressNameInfoRequest
*/
func (a *AddressNameAndCategoryInfoApiService) GetAddressNameInfo(ctx context.Context) ApiGetAddressNameInfoRequest {
	return ApiGetAddressNameInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAddressNameInfo200Response
func (a *AddressNameAndCategoryInfoApiService) GetAddressNameInfoExecute(r ApiGetAddressNameInfoRequest) (*GetAddressNameInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAddressNameInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressNameAndCategoryInfoApiService.GetAddressNameInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/addressname/simple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.address == nil {
		return localVarReturnValue, nil, reportError("address is required and must be specified")
	}
	if r.chain == nil {
		return localVarReturnValue, nil, reportError("chain is required and must be specified")
	}

	localVarQueryParams.Add("address", parameterToString(*r.address, ""))
	localVarQueryParams.Add("chain", parameterToString(*r.chain, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UserAPIAuthenticateProject400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UserAPIAuthenticateProject401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
