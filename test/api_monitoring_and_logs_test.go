/*
MongoDB Atlas Administration API

Testing MonitoringAndLogsApiService

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

func Test_openapi_MonitoringAndLogsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MonitoringAndLogsApiService DownloadLogsForOneClusterHostInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var hostName string
        var logName string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.DownloadLogsForOneClusterHostInOneProject(context.Background(), groupId, hostName, logName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService GetFTSIndexMeasurementsForNamespaceAndIndexName", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var processId string
        var indexName string
        var databaseName string
        var collectionName string
        var groupId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.GetFTSIndexMeasurementsForNamespaceAndIndexName(context.Background(), processId, indexName, databaseName, collectionName, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService GetFTSMetricTypes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var processId string
        var groupId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.GetFTSMetricTypes(context.Background(), processId, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService GetFTSNamespaceIndexMeasurements", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var processId string
        var databaseName string
        var collectionName string
        var groupId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.GetFTSNamespaceIndexMeasurements(context.Background(), processId, databaseName, collectionName, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService GetFTSNonIndexMeasurements", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var processId string
        var groupId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.GetFTSNonIndexMeasurements(context.Background(), processId, groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService GetHostLogs", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var hostName string
        var logName string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.GetHostLogs(context.Background(), groupId, hostName, logName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnAllMongodbProcessesInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnAllMongodbProcessesInOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnAvailableDatabasesForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnAvailableDatabasesForOneMongodbProcess(context.Background(), groupId, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnAvailableDisksForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnAvailableDisksForOneMongodbProcess(context.Background(), groupId, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnDatabasesForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var databaseName string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnDatabasesForOneMongodbProcess(context.Background(), groupId, databaseName, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnMeasurementsForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsForOneMongodbProcess(context.Background(), groupId, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnMeasurementsOfOneDatabaseForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var databaseName string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDatabaseForOneMongodbProcess(context.Background(), groupId, databaseName, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnMeasurementsOfOneDisk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var partitionName string
        var groupId string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDisk(context.Background(), partitionName, groupId, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnMeasurementsOfOneDiskForOneMongodbProcess", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var partitionName string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDiskForOneMongodbProcess(context.Background(), groupId, partitionName, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringAndLogsApiService ReturnOneMongodbProcessById", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var processId string

        resp, httpRes, err := apiClient.MonitoringAndLogsApi.ReturnOneMongodbProcessById(context.Background(), groupId, processId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}