# \PodcastsApi

All URIs are relative to *http://digitalrooster:6666/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PodcastsCreate**](PodcastsApi.md#PodcastsCreate) | **Post** /podcasts | Add a podcast source with RSS feed and add it to the list
[**PodcastsDelete**](PodcastsApi.md#PodcastsDelete) | **Delete** /podcasts/{id} | Delete Podcast from list
[**PodcastsReadAll**](PodcastsApi.md#PodcastsReadAll) | **Get** /podcasts | Read all podcast rss sources
[**PodcastsReadOne**](PodcastsApi.md#PodcastsReadOne) | **Get** /podcasts/{id} | Read one Podcast identified by id



## PodcastsCreate

> ResourceUuid PodcastsCreate(ctx).Podcast(podcast).Execute()

Add a podcast source with RSS feed and add it to the list

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
    podcast := openapiclient.Podcast{Id: "Id_example", Title: "Title_example", Url: "Url_example", UpdateInterval: 123} // Podcast | Podcast source to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PodcastsApi.PodcastsCreate(context.Background()).Podcast(podcast).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodcastsApi.PodcastsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PodcastsCreate`: ResourceUuid
    fmt.Fprintf(os.Stdout, "Response from `PodcastsApi.PodcastsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPodcastsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **podcast** | [**Podcast**](Podcast.md) | Podcast source to create | 

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


## PodcastsDelete

> PodcastsDelete(ctx, id).Execute()

Delete Podcast from list

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
    resp, r, err := api_client.PodcastsApi.PodcastsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodcastsApi.PodcastsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPodcastsDeleteRequest struct via the builder pattern


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


## PodcastsReadAll

> []Podcast PodcastsReadAll(ctx).Length(length).Offset(offset).Execute()

Read all podcast rss sources

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
    resp, r, err := api_client.PodcastsApi.PodcastsReadAll(context.Background()).Length(length).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodcastsApi.PodcastsReadAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PodcastsReadAll`: []Podcast
    fmt.Fprintf(os.Stdout, "Response from `PodcastsApi.PodcastsReadAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPodcastsReadAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | number of items if omitted or invalid all remaining elements will be assumed | 
 **offset** | **int32** | offset from start of list if omitted or invalid 0 will be assumed | 

### Return type

[**[]Podcast**](Podcast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PodcastsReadOne

> Podcast PodcastsReadOne(ctx, id).Execute()

Read one Podcast identified by id

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
    resp, r, err := api_client.PodcastsApi.PodcastsReadOne(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodcastsApi.PodcastsReadOne``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PodcastsReadOne`: Podcast
    fmt.Fprintf(os.Stdout, "Response from `PodcastsApi.PodcastsReadOne`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | unique id to identify item | 

### Other Parameters

Other parameters are passed through a pointer to a apiPodcastsReadOneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Podcast**](Podcast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

