package validators

import (
	"errors"
	"github.com/dintzeler/platform-go-challenge/models"
	"fmt"
)

type Error 

type AddFavoriteRequest struct {
	AssetID   *int              `json:"asset_id"`
	AssetType *models.AssetType `json:"asset_type"`
}

func ValidateAddFavoriteRequest(addFavoriteRequest AddFavoriteRequest) error {
	err := validateAssetID(addFavoriteRequest.AssetID)
	if err != nil {
		return err
	}
	err = validateAssetType(addFavoriteRequest.AssetType)
	if err != nil {
		return err
	}
	return nil
}
	
func validateAssetType(assetType *models.AssetType) error {
    if assetType == nil {
        return errors.New("asset_type is a required field")
    }

	switch *assetType {
	case models.AssetTypeChart, models.AssetTypeInsight, models.AssetTypeAudience:
		return nil
	default:
		return errors.New("Invalid asset_type (type must be 'chart', 'insight', or 'audience')")
	}
}

func validateAssetID(assetID *int) error {
	fmt.Println("Validating asset ID:", assetID)
	if assetID == nil {
		return errors.New("asset_id is a required field")
	}else if *assetID <= 0 {
		return errors.New("Invalid asset_id (must be a positive integer)")
	}
	return nil
}