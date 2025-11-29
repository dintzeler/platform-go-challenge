package validators

import (
	"github.com/dintzeler/platform-go-challenge/models"
)

type ValidationError struct {
	Message string `json:"message"`
	ErrorCode string `json:"error_code"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

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
        return &ValidationError{
			Message:   "asset_type is a required field",
			ErrorCode: "MISSING_ASSET_TYPE",
		}
    }

	switch *assetType {
	case models.AssetTypeChart, models.AssetTypeInsight, models.AssetTypeAudience:
		return nil
	default:
		return &ValidationError{
			Message:   "Invalid asset_type (type must be 'chart', 'insight', or 'audience')",
			ErrorCode: "INVALID_ASSET_TYPE",
		}
	}
}

func validateAssetID(assetID *int) error {
	if assetID == nil {
		return &ValidationError{
			Message:   "asset_id is a required field",
			ErrorCode: "MISSING_ASSET_ID",
		}
	}else if *assetID <= 0 {
		return &ValidationError{
			Message:   "asset_id must be a positive integer",
			ErrorCode: "INVALID_ASSET_ID",
		}
	}
	return nil
}