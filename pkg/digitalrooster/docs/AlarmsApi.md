# \AlarmsApi

All URIs are relative to *http://digitalrooster:6666/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlarmsCreate**](AlarmsApi.md#AlarmsCreate) | **Post** /alarms | Create a new alarm entry and add it to the list
[**AlarmsDelete**](AlarmsApi.md#AlarmsDelete) | **Delete** /alarms/{id} | Delete alarm from list
[**AlarmsReadAll**](AlarmsApi.md#AlarmsReadAll) | **Get** /alarms | Read all Alarms
[**AlarmsReadOne**](AlarmsApi.md#AlarmsReadOne) | **Get** /alarms/{id} | Read one Alarm



## AlarmsCreate

> ResourceUuid AlarmsCreate(ctx).Alarm(alarm).Execute()

Create a new alarm entry and add it to the list

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
    alarm := openapiclient.Alarm{Time: "Time_example", Period: "Period_example", Url: "Url_example", Enabled: false, Volume: 123, Id: "Id_example"} // Alarm | Alarm object to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlarmsApi.AlarmsCreate(context.Background()).Alarm(alarm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmsApi.AlarmsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsCreate`: ResourceUuid
    fmt.Fprintf(os.Stdout, "Response from `AlarmsApi.AlarmsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarm** | [**Alarm**](Alarm.md) | Alarm object to create | 

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


## AlarmsDelete

> AlarmsDelete(ctx, id).Execute()

Delete alarm from list

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
    resp, r, err := api_client.AlarmsApi.AlarmsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmsApi.AlarmsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlarmsDeleteRequest struct via the builder pattern


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


## AlarmsReadAll

> []Alarm AlarmsReadAll(ctx).Length(length).Offset(offset).Execute()

Read all Alarms

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
    resp, r, err := api_client.AlarmsApi.AlarmsReadAll(context.Background()).Length(length).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmsApi.AlarmsReadAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsReadAll`: []Alarm
    fmt.Fprintf(os.Stdout, "Response from `AlarmsApi.AlarmsReadAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsReadAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | number of items if omitted or invalid all remaining elements will be assumed | 
 **offset** | **int32** | offset from start of list if omitted or invalid 0 will be assumed | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlarmsReadOne

> Alarm AlarmsReadOne(ctx, id).Execute()

Read one Alarm

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
    resp, r, err := api_client.AlarmsApi.AlarmsReadOne(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlarmsApi.AlarmsReadOne``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsReadOne`: Alarm
    fmt.Fprintf(os.Stdout, "Response from `AlarmsApi.AlarmsReadOne`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | unique id to identify item | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsReadOneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Alarm**](Alarm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

