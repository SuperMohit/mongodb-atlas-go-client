/*
MongoDB Atlas Administration API

Testing LegacyBackupSnapshotsApiService

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

func Test_openapi_LegacyBackupSnapshotsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LegacyBackupSnapshotsApiService ChangeOneLegacyBackupSnapshotExpiration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.LegacyBackupSnapshotsApi.ChangeOneLegacyBackupSnapshotExpiration(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LegacyBackupSnapshotsApiService RemoveOneLegacyBackupSnapshot", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.LegacyBackupSnapshotsApi.RemoveOneLegacyBackupSnapshot(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LegacyBackupSnapshotsApiService ReturnAllLegacyBackupSnapshots", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.LegacyBackupSnapshotsApi.ReturnAllLegacyBackupSnapshots(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LegacyBackupSnapshotsApiService ReturnOneLegacyBackupSnapshot", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var snapshotId string

        resp, httpRes, err := apiClient.LegacyBackupSnapshotsApi.ReturnOneLegacyBackupSnapshot(context.Background(), groupId, clusterName, snapshotId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
