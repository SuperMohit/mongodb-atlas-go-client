/*
MongoDB Atlas Administration API

Testing AlertConfigurationsApiService

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

func Test_openapi_AlertConfigurationsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AlertConfigurationsApiService CreateOneAlertConfigurationInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.CreateOneAlertConfigurationInOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService RemoveOneAlertConfigurationFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var alertConfigId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.RemoveOneAlertConfigurationFromOneProject(context.Background(), groupId, alertConfigId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService ReturnAlertConfigMatchersFieldNames", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AlertConfigurationsApi.ReturnAlertConfigMatchersFieldNames(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService ReturnAllAlertConfigurationsForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.ReturnAllAlertConfigurationsForOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService ReturnAllOpenAlertsForAlertConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var alertConfigId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.ReturnAllOpenAlertsForAlertConfiguration(context.Background(), groupId, alertConfigId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService ReturnOneAlertConfigurationFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var alertConfigId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.ReturnOneAlertConfigurationFromOneProject(context.Background(), groupId, alertConfigId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService ToggleOneStateOfOneAlertConfigurationInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var alertConfigId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.ToggleOneStateOfOneAlertConfigurationInOneProject(context.Background(), groupId, alertConfigId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AlertConfigurationsApiService UpdateOneAlertConfigurationForOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var alertConfigId string

        resp, httpRes, err := apiClient.AlertConfigurationsApi.UpdateOneAlertConfigurationForOneProject(context.Background(), groupId, alertConfigId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
