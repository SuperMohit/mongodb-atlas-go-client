/*
MongoDB Atlas Administration API

Testing AccessTrackingApiService

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

func Test_openapi_AccessTrackingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AccessTrackingApiService ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AccessTrackingApiService ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var hostname string

        resp, httpRes, err := apiClient.AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname(context.Background(), groupId, hostname).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}