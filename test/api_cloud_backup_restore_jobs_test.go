/*
MongoDB Atlas Administration API

Testing CloudBackupRestoreJobsApiService

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

func Test_openapi_CloudBackupRestoreJobsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CloudBackupRestoreJobsApiService CancelOneRestoreJobOfOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var restoreJobId string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.CancelOneRestoreJobOfOneCluster(context.Background(), groupId, clusterName, restoreJobId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService RestoreOneSnapshotOfOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService RestoreOneSnapshotOfOneCluster1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster1(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService ReturnAllRestoreJobsForOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneCluster(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService ReturnAllRestoreJobsForOneServerlessInstance", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneServerlessInstance(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService ReturnOneRestoreJobForOneServerlessInstance", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var restoreJobId string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.ReturnOneRestoreJobForOneServerlessInstance(context.Background(), groupId, clusterName, restoreJobId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupRestoreJobsApiService ReturnOneRestoreJobOfOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var restoreJobId string

        resp, httpRes, err := apiClient.CloudBackupRestoreJobsApi.ReturnOneRestoreJobOfOneCluster(context.Background(), groupId, clusterName, restoreJobId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
