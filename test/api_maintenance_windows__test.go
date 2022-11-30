/*
MongoDB Atlas Administration API

Testing MaintenanceWindowsApiService

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

func Test_openapi_MaintenanceWindowsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MaintenanceWindowsApiService AutomateDeferralOfMaintenanceForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MaintenanceWindowsApi.AutomateDeferralOfMaintenanceForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MaintenanceWindowsApiService DeferOneMaintenanceWindowForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MaintenanceWindowsApi.DeferOneMaintenanceWindowForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MaintenanceWindowsApiService ResetOneMaintenanceWindowForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MaintenanceWindowsApi.ResetOneMaintenanceWindowForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MaintenanceWindowsApiService ReturnOneMaintenanceWindowForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MaintenanceWindowsApi.ReturnOneMaintenanceWindowForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MaintenanceWindowsApiService UpdateMaintenanceWindowForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.MaintenanceWindowsApi.UpdateMaintenanceWindowForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}