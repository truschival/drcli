# \RadiosApi

All URIs are relative to *http://digitalrooster:6666/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IradioCreate**](RadiosApi.md#IradioCreate) | **Post** /radios | Create a station info and add it to the list
[**IradioDelete**](RadiosApi.md#IradioDelete) | **Delete** /radios/{id} | Delete Internet radio station from list
[**IradioReadOne**](RadiosApi.md#IradioReadOne) | **Get** /radios/{id} | Read one Internet radio station



## IradioCreate

> ResourceUuid IradioCreate(ctx).Station(station).Execute()

Create a station info and add it to the list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    station := openapiclient.Station{Id: "Id_example", Name: "Name_example", Url: "Url_example"} // Station | Internet radio station to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RadiosApi.IradioCreate(context.Background()).Station(station).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadiosApi.IradioCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IradioCreate`: ResourceUuid
    fmt.Fprintf(os.Stdout, "Response from `RadiosApi.IradioCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIradioCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **station** | [**Station**](Station.md) | Internet radio station to create | 

### Return type

[**ResourceUuid**](ResourceUuid.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IradioDelete

> IradioDelete(ctx, id).Execute()

Delete Internet radio station from list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | unique id to identify item

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RadiosApi.IradioDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadiosApi.IradioDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | unique id to identify item | 

### Other Parameters

Other parameters are passed through a pointer to a apiIradioDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IradioReadOne

> Station IradioReadOne(ctx, id).Execute()

Read one Internet radio station

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | unique id to identify item

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RadiosApi.IradioReadOne(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadiosApi.IradioReadOne``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IradioReadOne`: Station
    fmt.Fprintf(os.Stdout, "Response from `RadiosApi.IradioReadOne`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | unique id to identify item | 

### Other Parameters

Other parameters are passed through a pointer to a apiIradioReadOneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Station**](Station.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

