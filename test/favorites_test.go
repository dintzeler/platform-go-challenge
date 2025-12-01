package tests

import (
    "github.com/dintzeler/platform-go-challenge/controllers"
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "strings"
    "os"
    "path/filepath"
    "runtime"
)

func init() {
    // Get the directory of this test file
    _, filename, _, _ := runtime.Caller(0)
    testDir := filepath.Dir(filename)
    
    // Go up one level to the project root
    projectRoot := filepath.Dir(testDir)
    
    // Set the absolute path to test_data.json
    dataFile := filepath.Join(projectRoot, "test_data.json")
    os.Setenv("DATA_FILE", dataFile)
    
    // Verify the file exists
    if _, err := os.Stat(dataFile); os.IsNotExist(err) {
        panic("test_data.json not found at: " + dataFile)
    }
}

func TestAddFavorite(t *testing.T) {
	t.Run("Add favorite with invalid User-ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/favorites", nil)
		controllers.FavoritesHandler(w, req) 
		
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
        var response map[string]interface{}
        err := json.Unmarshal(w.Body.Bytes(), &response)
        if err != nil {
            t.Fatalf("Failed to parse JSON response: %v", err)
        }

        if message := response["message"]; message != "Invalid User-ID"  {
            t.Errorf("Response message: %v", message)
        }

		if errorCode := response["error_code"]; errorCode != "INVALID_USER_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	
	t.Run("Add favorite without providing request body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(""))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Invalid request body"  {
			t.Errorf("Response message: %v", message)
		}
		if errorCode := response["error_code"]; errorCode != "INVALID_REQUEST_BODY" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite without providing asset_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_type": "chart"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "asset_id is a required field"  {
			t.Errorf("Response message: %v", message)
		}
		if errorCode := response["error_code"]; errorCode != "MISSING_ASSET_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite without providing asset_type", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "asset_type is a required field"  {
			t.Errorf("Response message: %v", message)
		}
		if errorCode := response["error_code"]; errorCode != "MISSING_ASSET_TYPE" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite with negative asset_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": -5, "asset_type": "chart"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "asset_id must be a positive integer"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_ASSET_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite with invalid asset_type", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1, "asset_type": "invalid_type"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Invalid asset_type (type must be 'chart', 'insight', or 'audience')"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_ASSET_TYPE" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite for non-existing asset", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 9999, "asset_type": "chart"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Asset does not exist"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "ASSET_NOT_FOUND" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})

	t.Run("Add favorite that already exists", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1, "asset_type": "chart"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req) 

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if assetID := int(response["asset_id"].(float64)); assetID != 1 {
			t.Errorf("Response asset_id: %v", assetID)
		}
		if assetType := response["asset_type"]; assetType != "chart" {
			t.Errorf("Response asset_type: %v", assetType)
		}
		if userID := int(response["user_id"].(float64)); userID != 1 {
			t.Errorf("Response user_id: %v", userID)
		}
	})
	
	t.Run("Add favorite successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 2, "asset_type": "insight"}`
		req, _ := http.NewRequest("POST", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if assetID := int(response["asset_id"].(float64)); assetID != 2 {
			t.Errorf("Response asset_id: %v", assetID)
		}

		if assetType := response["asset_type"]; assetType != "insight" {
			t.Errorf("Response asset_type: %v", assetType)
		}

		if userID := int(response["user_id"].(float64)); userID != 1 {
			t.Errorf("Response user_id: %v", userID)
		}

		// Verify favorite is persisted in data file
		verifyFavoriteInDataFile(t, 1, 2, "insight")
	})
}

func TestUpdateFavorite(t *testing.T) {
	t.Run("Update favorite with invalid User-ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/favorites", nil)
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Invalid User-ID"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_USER_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite without providing request body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(""))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Invalid request body"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_REQUEST_BODY" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite without providing asset_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_type": "chart"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "asset_id is a required field"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "MISSING_ASSET_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite without providing asset_type", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)

		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "asset_type is a required field"  {
			t.Errorf("Response message: %v", message)
		}
		if errorCode := response["error_code"]; errorCode != "MISSING_ASSET_TYPE" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite with negative asset_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": -5, "asset_type": "chart"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}
		
		if message := response["message"]; message != "asset_id must be a positive integer"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_ASSET_ID" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite with missing description", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1, "asset_type": "chart"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "description is a required field"  {
			t.Errorf("Response message: %v", message)
		}
		if errorCode := response["error_code"]; errorCode != "MISSING_DESCRIPTION" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite with invalid asset_type", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1, "asset_type": "invalid_type", "description": "New description"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Invalid asset_type (type must be 'chart', 'insight', or 'audience')"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "INVALID_ASSET_TYPE" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite for non-existing asset", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 9999, "asset_type": "chart", "description": "New description"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "Asset does not exist"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "ASSET_NOT_FOUND" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update asset that is not favorited", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 2, "asset_type": "chart", "description": "Updated description"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusForbidden {
			t.Errorf("Expected status 403, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if message := response["message"]; message != "User does not have this asset as favorite"  {
			t.Errorf("Response message: %v", message)
		}

		if errorCode := response["error_code"]; errorCode != "ASSET_NOT_FAVORITE" {
			t.Errorf("Response error_code: %v", errorCode)
		}
	})
	t.Run("Update favorite successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := `{"asset_id": 1, "asset_type": "chart", "description": "Updated description"}`
		req, _ := http.NewRequest("PATCH", "/favorites", strings.NewReader(reqBody))
		req.Header.Set("User-ID", "1")
		controllers.FavoritesHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		// Check JSON response
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse JSON response: %v", err)
		}

		if assetID := int(response["asset_id"].(float64)); assetID != 1 {
			t.Errorf("Response asset_id: %v", assetID)
		}

		if assetType := response["asset_type"]; assetType != "chart" {
			t.Errorf("Response asset_type: %v", assetType)
		}

		if userID := int(response["user_id"].(float64)); userID != 1 {
			t.Errorf("Response user_id: %v", userID)
		}

		if description := response["updated_description"]; description != "Updated description" {
			t.Errorf("Response description: %v", description)
		}

		//verify description changed
		verifyDescriptionChanged(t, 1, 1, "chart", "Updated description")
	})
}

func verifyFavoriteInDataFile(t *testing.T, userID, assetID int, assetType string) {
    dataFile := os.Getenv("DATA_FILE")
    if dataFile == "" {
        t.Fatal("DATA_FILE environment variable not set")
    }

    content, err := os.ReadFile(dataFile)
    if err != nil {
        t.Fatalf("Failed to read test data file: %v", err)
    }

    var data map[string]interface{}
    err = json.Unmarshal(content, &data)
    if err != nil {
        t.Fatalf("Failed to parse test data JSON: %v", err)
    }

    favorites, ok := data["favorites"].([]interface{})
    if !ok {
        t.Fatal("Favorites not found or not an array in test data")
    }

    found := false
    for _, fav := range favorites {
        favorite := fav.(map[string]interface{})
        favUserID := int(favorite["user_id"].(float64))
        favAssetID := int(favorite["asset_id"].(float64))
        favAssetType := favorite["asset_type"].(string)

        if favUserID == userID && favAssetID == assetID && favAssetType == assetType {
            found = true
            break
        }
    }

    if !found {
        t.Errorf("Favorite not found in data file: user_id=%d, asset_id=%d, asset_type=%s", userID, assetID, assetType)
    }
}

func verifyDescriptionChanged(t *testing.T, userID, assetID int, assetType, expectedDescription string) {
	dataFile := os.Getenv("DATA_FILE")
	if dataFile == "" {
		t.Fatal("DATA_FILE environment variable not set")
	}

	content, err := os.ReadFile(dataFile)
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		t.Fatalf("Failed to parse test data JSON: %v", err)
	}

	charts, ok := data["charts"].([]interface{})
	if !ok {
		t.Fatal("Charts not found or not an array in test data")
	}

	found := false
	for _, ch := range charts {
		chart := ch.(map[string]interface{})
		chID := int(chart["id"].(float64))
		if chID == assetID {
			description := chart["description"].(string)
			if description != expectedDescription {
				t.Errorf("Description mismatch: expected '%s', got '%s'", expectedDescription, description)
			}
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Asset not found in data file: asset_id=%d, asset_type=%s", assetID, assetType)
	}
}




