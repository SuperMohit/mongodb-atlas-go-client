/*
MongoDB Atlas Administration API

Testing CloudBackupExportApiService

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

func Test_openapi_CloudBackupExportApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CloudBackupExportApiService CreateOneCloudBackupSnapshotExportJob", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupExportApi.CreateOneCloudBackupSnapshotExportJob(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService GrantAccessToAwsS3BucketForCloudBackupSnapshotExports", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.CloudBackupExportApi.GrantAccessToAwsS3BucketForCloudBackupSnapshotExports(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.CloudBackupExportApi.ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService ReturnAllCloudBackupSnapshotExportJobs", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.CloudBackupExportApi.ReturnAllCloudBackupSnapshotExportJobs(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var exportBucketId string

        resp, httpRes, err := apiClient.CloudBackupExportApi.ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports(context.Background(), groupId, exportBucketId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService ReturnOneCloudBackupSnapshotExportJob", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var exportId string

        resp, httpRes, err := apiClient.CloudBackupExportApi.ReturnOneCloudBackupSnapshotExportJob(context.Background(), groupId, clusterName, exportId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudBackupExportApiService RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var exportBucketId string

        resp, httpRes, err := apiClient.CloudBackupExportApi.RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports(context.Background(), groupId, exportBucketId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}