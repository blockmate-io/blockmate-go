# Go API client for blockmate

Blockmate API OpenAPI documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import blockmate "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), blockmate.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), blockmate.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), blockmate.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), blockmate.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.blockmate.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountProviderInfoApi* | [**ConnectAccount**](docs/AccountProviderInfoApi.md#connectaccount) | **Post** /v1/{account_provider}/connect | Connect new account
*AccountProviderInfoApi* | [**DeleteAccount**](docs/AccountProviderInfoApi.md#deleteaccount) | **Delete** /v1/{account_provider}/account/{account_id} | Delete account
*AccountProviderInfoApi* | [**GetAccountHint**](docs/AccountProviderInfoApi.md#getaccounthint) | **Get** /v1/{account_provider}/connect | Get account hint
*AddressNameAndCategoryInfoApi* | [**GetAddressNameInfoMulti**](docs/AddressNameAndCategoryInfoApi.md#getaddressnameinfomulti) | **Post** /v1/addressname/multi | Get address name and category info for multiple addresses
*AddressNameAndCategoryInfoApi* | [**GetAddressNameInfoSingle**](docs/AddressNameAndCategoryInfoApi.md#getaddressnameinfosingle) | **Get** /v1/addressname/simple | Get address name and category info for single address
*AggregatedInfoApi* | [**AccountProviderHintsList**](docs/AggregatedInfoApi.md#accountproviderhintslist) | **Get** /v1/aggregate/account_provider_hints | Get list of account providers hints
*AggregatedInfoApi* | [**AccountProvidersList**](docs/AggregatedInfoApi.md#accountproviderslist) | **Get** /v1/aggregate/account_providers | Get list of account providers
*AggregatedInfoApi* | [**Accounts**](docs/AggregatedInfoApi.md#accounts) | **Get** /v1/aggregate/accounts | List accounts
*AggregatedInfoApi* | [**Balance**](docs/AggregatedInfoApi.md#balance) | **Get** /v1/aggregate/balance | Get balance
*AggregatedInfoApi* | [**NFTMetadata**](docs/AggregatedInfoApi.md#nftmetadata) | **Get** /v1/aggregate/nft_metadata | Get NFT metadata
*AggregatedInfoApi* | [**Transactions**](docs/AggregatedInfoApi.md#transactions) | **Get** /v1/aggregate/transactions | Get transactions
*AnalyticsApi* | [**GetAnalytics**](docs/AnalyticsApi.md#getanalytics) | **Get** /v1/analytics/{account_provider}/account/{account_id}/stats | Get analytics focused on gaming
*AuthenticationApi* | [**UserAPIAuthenticateDeveloper**](docs/AuthenticationApi.md#userapiauthenticatedeveloper) | **Get** /v1/auth/developer | Authenticate developer
*AuthenticationApi* | [**UserAPIAuthenticateProject**](docs/AuthenticationApi.md#userapiauthenticateproject) | **Get** /v1/auth | Authenticate project
*ENSApi* | [**GetAddressFromDomain**](docs/ENSApi.md#getaddressfromdomain) | **Get** /v1/ens/addressFromDomain | Get address for specified ENS domain
*ENSApi* | [**GetDomainFromAddress**](docs/ENSApi.md#getdomainfromaddress) | **Get** /v1/ens/domainFromAddress | Get domain and metadata for specified ENS address
*ExchangeRateInfoApi* | [**GetCurrentExchangeRate**](docs/ExchangeRateInfoApi.md#getcurrentexchangerate) | **Get** /v1/exchangerate/current | Get current exchange rate
*ExchangeRateInfoApi* | [**GetHistoricalExchangeRate**](docs/ExchangeRateInfoApi.md#gethistoricalexchangerate) | **Get** /v1/exchangerate/history | Get historical exchange rate
*RiskInfoApi* | [**GetAddressRiskScore**](docs/RiskInfoApi.md#getaddressriskscore) | **Get** /v1/risk/score | Get address risk score
*RiskInfoApi* | [**GetAddressRiskScoreCase**](docs/RiskInfoApi.md#getaddressriskscorecase) | **Get** /v1/risk/score/details/{case_id} | Get address risk score case
*RiskInfoApi* | [**GetAddressRiskScoreDetails**](docs/RiskInfoApi.md#getaddressriskscoredetails) | **Get** /v1/risk/score/details | Get address risk score details
*RiskInfoApi* | [**GetTransactionRiskScore**](docs/RiskInfoApi.md#gettransactionriskscore) | **Get** /v1/risk/transaction/score | Get transaction risk score
*RiskInfoApi* | [**GetTransactionRiskScoreCase**](docs/RiskInfoApi.md#gettransactionriskscorecase) | **Get** /v1/risk/transaction/score/details/{case_id} | Get transaction risk score case
*RiskInfoApi* | [**GetTransactionRiskScoreDetails**](docs/RiskInfoApi.md#gettransactionriskscoredetails) | **Get** /v1/risk/transaction/score/details | Get transaction risk score details
*UserManagementApi* | [**AuthUser**](docs/UserManagementApi.md#authuser) | **Get** /v1/users/{id}/auth | Authenticate user
*UserManagementApi* | [**CreateUser**](docs/UserManagementApi.md#createuser) | **Post** /v1/users | Create user
*UserManagementApi* | [**DeleteUser**](docs/UserManagementApi.md#deleteuser) | **Delete** /v1/users/{id} | Delete user
*UserManagementApi* | [**GetUser**](docs/UserManagementApi.md#getuser) | **Get** /v1/users/{id} | Get user
*UserManagementApi* | [**ListUsers**](docs/UserManagementApi.md#listusers) | **Get** /v1/users | List users
*UserManagementApi* | [**UpdateUser**](docs/UserManagementApi.md#updateuser) | **Post** /v1/users/{id} | Update user


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountProvider](docs/AccountProvider.md)
 - [AccountProviderHint](docs/AccountProviderHint.md)
 - [AccountProviderHintFields](docs/AccountProviderHintFields.md)
 - [AddressRiskReport](docs/AddressRiskReport.md)
 - [AddressRiskReportDetails](docs/AddressRiskReportDetails.md)
 - [Amount](docs/Amount.md)
 - [AuthUser200Response](docs/AuthUser200Response.md)
 - [Balance200Response](docs/Balance200Response.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [BalanceResponseAccountsInner](docs/BalanceResponseAccountsInner.md)
 - [BalanceResponseAccountsInnerState](docs/BalanceResponseAccountsInnerState.md)
 - [ConnectAccount200Response](docs/ConnectAccount200Response.md)
 - [ConnectAccount400Response](docs/ConnectAccount400Response.md)
 - [ConnectAccount405Response](docs/ConnectAccount405Response.md)
 - [ConnectAccountRequest](docs/ConnectAccountRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [DeleteAccount404Response](docs/DeleteAccount404Response.md)
 - [DeleteUser200Response](docs/DeleteUser200Response.md)
 - [DeleteUser404Response](docs/DeleteUser404Response.md)
 - [ExchangeRate](docs/ExchangeRate.md)
 - [GetAccountHint200Response](docs/GetAccountHint200Response.md)
 - [GetAccountHint200ResponseFields](docs/GetAccountHint200ResponseFields.md)
 - [GetAccountHint403Response](docs/GetAccountHint403Response.md)
 - [GetAccountHint404Response](docs/GetAccountHint404Response.md)
 - [GetAddressFromDomain200Response](docs/GetAddressFromDomain200Response.md)
 - [GetAddressNameInfoSingle200Response](docs/GetAddressNameInfoSingle200Response.md)
 - [GetAddressRiskScore200Response](docs/GetAddressRiskScore200Response.md)
 - [GetAnalytics200Response](docs/GetAnalytics200Response.md)
 - [GetAnalytics200ResponseValue](docs/GetAnalytics200ResponseValue.md)
 - [GetDomainFromAddress200Response](docs/GetDomainFromAddress200Response.md)
 - [GetDomainFromAddress200ResponseMetadata](docs/GetDomainFromAddress200ResponseMetadata.md)
 - [GetDomainFromAddress200ResponseMetadataAttributesInner](docs/GetDomainFromAddress200ResponseMetadataAttributesInner.md)
 - [GetDomainFromAddress200ResponseMetadataAttributesInnerValue](docs/GetDomainFromAddress200ResponseMetadataAttributesInnerValue.md)
 - [GetTransactionRiskScore200Response](docs/GetTransactionRiskScore200Response.md)
 - [GetUser404Response](docs/GetUser404Response.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataAttributesInner](docs/MetadataAttributesInner.md)
 - [Movement](docs/Movement.md)
 - [NFTMetadata200ResponseValue](docs/NFTMetadata200ResponseValue.md)
 - [NftContractMetadata](docs/NftContractMetadata.md)
 - [NftId](docs/NftId.md)
 - [NftIdTokenMetadata](docs/NftIdTokenMetadata.md)
 - [NftMedia](docs/NftMedia.md)
 - [NftSpamInfo](docs/NftSpamInfo.md)
 - [NftTokenUri](docs/NftTokenUri.md)
 - [OwnedNft](docs/OwnedNft.md)
 - [OwnedNftContract](docs/OwnedNftContract.md)
 - [OwnedNftMedia](docs/OwnedNftMedia.md)
 - [RiskReportCategory](docs/RiskReportCategory.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionRiskReport](docs/TransactionRiskReport.md)
 - [Transactions200Response](docs/Transactions200Response.md)
 - [Transactions200ResponseAccountsInner](docs/Transactions200ResponseAccountsInner.md)
 - [User](docs/User.md)
 - [UserAPIAuthenticateProject200Response](docs/UserAPIAuthenticateProject200Response.md)
 - [UserAPIAuthenticateProject400Response](docs/UserAPIAuthenticateProject400Response.md)
 - [UserAPIAuthenticateProject401Response](docs/UserAPIAuthenticateProject401Response.md)


## Documentation For Authorization



### ProjectJWT

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### ProjectToken

- **Type**: API key
- **API key parameter name**: X-API-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-KEY and passed in as the auth context for each request.


### UserJWT

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



