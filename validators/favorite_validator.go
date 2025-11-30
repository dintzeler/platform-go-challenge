package validators

import (
	"github.com/dintzeler/platform-go-challenge/models"
	"github.com/dintzeler/platform-go-challenge/customerrors"
	"net/http"
	"strconv"
)

type AddFavoriteRequest struct {
	AssetID   *int              `json:"asset_id"`
	AssetType *models.AssetType `json:"asset_type"`
}

type DeleteFavoriteRequest struct {
	AssetID *int              `json:"asset_id"`
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

func ValidateDeleteFavorite(r *http.Request) (*DeleteFavoriteRequest, error) {
	assetIDStr := r.URL.Query().Get("asset_id")
	if assetIDStr == "" {
		return nil, &customerrors.ValidationError{
			Message:   "asset_id is a required field",
			ErrorCode: "MISSING_ASSET_ID",
		}
	}
	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		return nil, &customerrors.ValidationError{
			Message:   "Invalid asset_id",
			ErrorCode: "INVALID_ASSET_ID",
		}
	}
	err = validateAssetID(&assetID)
	if err != nil {
		return nil, err
	}

	assetTypeStr := r.URL.Query().Get("asset_type")
	if assetTypeStr == "" {
		return nil, &customerrors.ValidationError{
			Message:   "asset_type is a required field",
			ErrorCode: "MISSING_ASSET_TYPE",
		}
	}

	assetType := models.AssetType(assetTypeStr)
	err = validateAssetType(&assetType)
	if err != nil {
		return nil, err
	}

	return &DeleteFavoriteRequest{
		AssetID: &assetID,
		AssetType: &assetType,
	}, nil
}

func validateAssetType(assetType *models.AssetType) error {
    if assetType == nil {
        return &customerrors.ValidationError{
			Message:   "asset_type is a required field",
			ErrorCode: "MISSING_ASSET_TYPE",
		}
    }

	switch *assetType {
	case models.AssetTypeChart, models.AssetTypeInsight, models.AssetTypeAudience:
		return nil
	default:
		return &customerrors.ValidationError{
			Message:   "Invalid asset_type (type must be 'chart', 'insight', or 'audience')",
			ErrorCode: "INVALID_ASSET_TYPE",
		}
	}
}

func validateAssetID(assetID *int) error {
	if assetID == nil {
		return &customerrors.ValidationError{
			Message:   "asset_id is a required field",
			ErrorCode: "MISSING_ASSET_ID",
		}
	}else if *assetID <= 0 {
		return &customerrors.ValidationError{
			Message:   "asset_id must be a positive integer",
			ErrorCode: "INVALID_ASSET_ID",
		}
	}
	return nil
}