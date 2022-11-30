/*
MongoDB Atlas Administration API

Testing ProjectIPAccessListApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_ProjectIPAccessListApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProjectIPAccessListApiService AddEntriesToProjectIpAccessList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ProjectIPAccessListApi.AddEntriesToProjectIpAccessList(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProjectIPAccessListApiService GetAtlasNetworkPermissionEntryStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var entryValue string

        resp, httpRes, err := apiClient.ProjectIPAccessListApi.GetAtlasNetworkPermissionEntryStatus(context.Background(), groupId, entryValue).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProjectIPAccessListApiService RemoveOneEntryFromOneProjectIpAccessList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var entryValue string

        resp, httpRes, err := apiClient.ProjectIPAccessListApi.RemoveOneEntryFromOneProjectIpAccessList(context.Background(), groupId, entryValue).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProjectIPAccessListApiService ReturnOneProjectIpAccessListEntry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var entryValue string

        resp, httpRes, err := apiClient.ProjectIPAccessListApi.ReturnOneProjectIpAccessListEntry(context.Background(), groupId, entryValue).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProjectIPAccessListApiService ReturnProjectIpAccessList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ProjectIPAccessListApi.ReturnProjectIpAccessList(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
