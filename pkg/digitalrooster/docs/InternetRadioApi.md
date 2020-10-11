# \InternetRadioApi

All URIs are relative to *http://digitalrooster:6666/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IradioReadAll**](InternetRadioApi.md#IradioReadAll) | **Get** /radios | Read all Internet radio streams



## IradioReadAll

> []Station IradioReadAll(ctx).Length(length).Offset(offset).Execute()

Read all Internet radio streams



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
    length := 987 // int32 | number of items if omitted or invalid all remaining elements will be assumed (optional)
    offset := 987 // int32 | offset from start of list if omitted or invalid 0 will be assumed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternetRadioApi.IradioReadAll(context.Background()).Length(length).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternetRadioApi.IradioReadAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IradioReadAll`: []Station
    fmt.Fprintf(os.Stdout, "Response from `InternetRadioApi.IradioReadAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIradioReadAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | number of items if omitted or invalid all remaining elements will be assumed | 
 **offset** | **int32** | offset from start of list if omitted or invalid 0 will be assumed | 

### Return type

[**[]Station**](Station.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

