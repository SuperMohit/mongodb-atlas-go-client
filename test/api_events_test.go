/*
MongoDB Atlas Administration API

Testing EventsApiService

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

func Test_openapi_EventsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test EventsApiService ReturnAllEventsFromOneOrganization", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string

        resp, httpRes, err := apiClient.EventsApi.ReturnAllEventsFromOneOrganization(context.Background(), orgId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EventsApiService ReturnAllEventsFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.EventsApi.ReturnAllEventsFromOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EventsApiService ReturnOneEventFromOneOrganization", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var eventId string

        resp, httpRes, err := apiClient.EventsApi.ReturnOneEventFromOneOrganization(context.Background(), orgId, eventId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test EventsApiService ReturnOneEventFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var eventId string

        resp, httpRes, err := apiClient.EventsApi.ReturnOneEventFromOneProject(context.Background(), groupId, eventId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}