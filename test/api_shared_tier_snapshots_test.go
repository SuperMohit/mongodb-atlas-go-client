/*
MongoDB Atlas Administration API

Testing SharedTierSnapshotsApiService

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

func Test_openapi_SharedTierSnapshotsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SharedTierSnapshotsApiService DownloadOneM2OrM5ClusterSnapshot", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clusterName string
        var groupId string

        resp, httpRes, err := apiClient.SharedTierSnapshotsApi.DownloadOneM2OrM5ClusterSnapshot(context.Background(), clusterName, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SharedTierSnapshotsApiService ReturnAllSnapshotsForOneM2OrM5Cluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.SharedTierSnapshotsApi.ReturnAllSnapshotsForOneM2OrM5Cluster(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SharedTierSnapshotsApiService ReturnOneSnapshotForOneM2OrM5Cluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.SharedTierSnapshotsApi.ReturnOneSnapshotForOneM2OrM5Cluster(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
