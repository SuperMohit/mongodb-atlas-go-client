/*
MongoDB Atlas Administration API

Testing CloudBackupsApiService

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

func Test_openapi_CloudBackupsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CloudBackupsApiService ChangeExpirationDateForOneCloudBackup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.ChangeExpirationDateForOneCloudBackup(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService RemoveOneReplicaSetCloudBackup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.RemoveOneReplicaSetCloudBackup(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService RemoveOneShardedClusterCloudBackup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.RemoveOneShardedClusterCloudBackup(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnAllReplicaSetCloudBackups", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnAllReplicaSetCloudBackups(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnAllShardedClusterCloudBackups", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnAllShardedClusterCloudBackups(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnAllSnapshotsOfOneServerlessInstance", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnAllSnapshotsOfOneServerlessInstance(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnOneReplicaSetCloudBackup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnOneReplicaSetCloudBackup(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnOneShardedClusterCloudBackup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnOneShardedClusterCloudBackup(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService ReturnOneSnapshotOfOneServerlessInstance", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.CloudBackupsApi.ReturnOneSnapshotOfOneServerlessInstance(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupsApiService TakeOneOnDemandSnapshot", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupsApi.TakeOneOnDemandSnapshot(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}