package promotionPolicyController

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/promotionPolicyClient"
)

func GetArtifactPromotionPolicyController(name string) (ArtifactPromotionPolicy.PayloadPolicyForCreate, error) {

	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return ArtifactPromotionPolicy.PayloadPolicyForCreate{}, nil
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return ArtifactPromotionPolicy.PayloadPolicyForCreate{}, nil

	}
	return promotionPolicyClient.GetArtifactPromotionPolicyDetails(name)
}
