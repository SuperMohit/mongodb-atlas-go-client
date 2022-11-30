/*
MongoDB Atlas Administration API

Testing SharedTierRestoreJobsApiService

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

func Test_openapi_SharedTierRestoreJobsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SharedTierRestoreJobsApiService CreateOneRestoreJobFromOneM2OrM5Cluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clusterName string
        var groupId string

        resp, httpRes, err := apiClient.SharedTierRestoreJobsApi.CreateOneRestoreJobFromOneM2OrM5Cluster(context.Background(), clusterName, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SharedTierRestoreJobsApiService ReturnAllRestoreJobsForOneM2OrM5Cluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clusterName string
        var groupId string

        resp, httpRes, err := apiClient.SharedTierRestoreJobsApi.ReturnAllRestoreJobsForOneM2OrM5Cluster(context.Background(), clusterName, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SharedTierRestoreJobsApiService ReturnOneRestoreJobForOneM2OrM5Cluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clusterName string
        var groupId string
        var restoreId string

        resp, httpRes, err := apiClient.SharedTierRestoreJobsApi.ReturnOneRestoreJobForOneM2OrM5Cluster(context.Background(), clusterName, groupId, restoreId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
