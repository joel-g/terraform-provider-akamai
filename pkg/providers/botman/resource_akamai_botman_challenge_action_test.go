package botman

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v6/pkg/botman"
	"github.com/akamai/terraform-provider-akamai/v4/pkg/test"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestResourceChallengeAction(t *testing.T) {
	t.Run("ResourceChallengeAction", func(t *testing.T) {

		mockedBotmanClient := &botman.Mock{}
		createResponse := map[string]interface{}{"actionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey": "testValue3"}
		createRequest := test.FixtureBytes("testdata/JsonPayload/create.json")
		mockedBotmanClient.On("CreateChallengeAction",
			mock.Anything,
			botman.CreateChallengeActionRequest{
				ConfigID:    43253,
				Version:     15,
				JsonPayload: createRequest,
			},
		).Return(createResponse, nil).Once()

		mockedBotmanClient.On("GetChallengeAction",
			mock.Anything,
			botman.GetChallengeActionRequest{
				ConfigID: 43253,
				Version:  15,
				ActionID: "cc9c3f89-e179-4892-89cf-d5e623ba9dc7",
			},
		).Return(createResponse, nil).Times(3)
		expectedCreateJSON := `{"testKey":"testValue3"}`

		updateResponse := map[string]interface{}{"actionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey": "updated_testValue3"}
		updateRequest := `{"actionId":"cc9c3f89-e179-4892-89cf-d5e623ba9dc7","testKey":"updated_testValue3"}`
		mockedBotmanClient.On("UpdateChallengeAction",
			mock.Anything,
			botman.UpdateChallengeActionRequest{
				ConfigID:    43253,
				Version:     15,
				ActionID:    "cc9c3f89-e179-4892-89cf-d5e623ba9dc7",
				JsonPayload: json.RawMessage(updateRequest),
			},
		).Return(updateResponse, nil).Once()

		mockedBotmanClient.On("GetChallengeAction",
			mock.Anything,
			botman.GetChallengeActionRequest{
				ConfigID: 43253,
				Version:  15,
				ActionID: "cc9c3f89-e179-4892-89cf-d5e623ba9dc7",
			},
		).Return(updateResponse, nil).Times(2)
		expectedUpdateJSON := `{"testKey":"updated_testValue3"}`

		mockedBotmanClient.On("RemoveChallengeAction",
			mock.Anything,
			botman.RemoveChallengeActionRequest{
				ConfigID: 43253,
				Version:  15,
				ActionID: "cc9c3f89-e179-4892-89cf-d5e623ba9dc7",
			},
		).Return(nil).Once()

		useClient(mockedBotmanClient, func() {

			resource.Test(t, resource.TestCase{
				IsUnitTest:        true,
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: test.Fixture("testdata/TestResourceChallengeAction/create.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_botman_challenge_action.test", "id", "43253:cc9c3f89-e179-4892-89cf-d5e623ba9dc7"),
							resource.TestCheckResourceAttr("akamai_botman_challenge_action.test", "challenge_action", expectedCreateJSON)),
					},
					{
						Config: test.Fixture("testdata/TestResourceChallengeAction/update.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_botman_challenge_action.test", "id", "43253:cc9c3f89-e179-4892-89cf-d5e623ba9dc7"),
							resource.TestCheckResourceAttr("akamai_botman_challenge_action.test", "challenge_action", expectedUpdateJSON)),
					},
				},
			})
		})

		mockedBotmanClient.AssertExpectations(t)
	})
}