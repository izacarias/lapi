package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/mock"
	"github.com/izacarias/lapi/routes"
)

// setupTestRouter creates and configures a router for testing
func setupTestRouter() *gin.Engine {
	// Set to test mode to avoid unnecessary logs
	gin.SetMode(gin.TestMode)

	// Connect to MongoDB
	client := configs.ConnectDB()

	// Insert mock data
	mock.InsertMockData(client)

	// Create and configure router
	router := gin.Default()
	routes.PingRoute(router)
	routes.ZoneRoute(router)
	routes.ApRoute(router)
	routes.UserRoute(router)
	routes.DistanceRoute(router)
	routes.SwaggerRoute(router)

	return router
}

// TestIntegration tests all the endpoints in the system using the routes from Requests.rest
func TestIntegration(t *testing.T) {
	// Setup test router
	router := setupTestRouter()

	// Create test cases based on Requests.rest
	testCases := []struct {
		name           string
		method         string
		url            string
		expectStatus   int
		expectFailure  bool
		skipStatusTest bool
	}{
		{
			name:         "Ping Endpoint",
			method:       "GET",
			url:          "/ping",
			expectStatus: http.StatusOK,
		},
		{
			name:         "List All Zones",
			method:       "GET",
			url:          "/location/v3/queries/zones",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Zone by ID",
			method:       "GET",
			url:          "/location/v3/queries/zones/zone1",
			expectStatus: http.StatusOK,
		},
		{
			name:         "List APs in Zone1",
			method:       "GET",
			url:          "/location/v3/queries/zones/zone1/accessPoints",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get AP1 in Zone1",
			method:       "GET",
			url:          "/location/v3/queries/zones/zone1/accessPoints/ap1",
			expectStatus: http.StatusOK,
		},
		{
			name:          "Get AP2 in Zone1 (Should return 404)",
			method:        "GET",
			url:           "/location/v3/queries/zones/zone1/accessPoints/ap2",
			expectStatus:  http.StatusNotFound,
			expectFailure: true, // The comment in Requests.rest mentions this should return 404
		},
		{
			name:         "Get AP2 in Zone2",
			method:       "GET",
			url:          "/location/v3/queries/zones/zone2/accessPoints/ap2",
			expectStatus: http.StatusOK,
		},
		{
			name:         "List Users in Zone1",
			method:       "GET",
			url:          "/location/v3/queries/users?zoneId=zone1",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get User by IP Address",
			method:       "GET",
			url:          "/location/v3/queries/users?address=192.168.1.1",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Users by Multiple IP Addresses",
			method:       "GET",
			url:          "/location/v3/queries/users?address=192.168.1.1&address=192.168.1.2",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get All Users",
			method:       "GET",
			url:          "/location/v3/queries/users",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Users Connected to AP1",
			method:       "GET",
			url:          "/location/v3/queries/users?accessPointId=ap1",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Users Connected to AP1 and AP2",
			method:       "GET",
			url:          "/location/v3/queries/users?accessPointId=ap1&accessPointId=ap2",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Users Connected to Non-existent AP",
			method:       "GET",
			url:          "/location/v3/queries/users?accessPointId=ap999",
			expectStatus: http.StatusOK, // Even though it returns empty, it should be 200 OK
		},
		{
			name:         "Get Distance Between Users",
			method:       "GET",
			url:          "/location/v3/queries/distance?address=192.168.1.1&address=192.168.1.5",
			expectStatus: http.StatusOK,
		},
		{
			name:         "Get Distance Between User and Point",
			method:       "GET",
			url:          "/location/v3/queries/distance?address=192.168.1.1&latitude=10&longitude=10",
			expectStatus: http.StatusOK,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test request
			req, err := http.NewRequest(tc.method, tc.url, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Check the status code
			if !tc.skipStatusTest {
				if tc.expectFailure {
					// For this specific case, we expect a failure (404)
					if w.Code != tc.expectStatus {
						t.Errorf("Expected status %d, got %d", tc.expectStatus, w.Code)
					}
				} else {
					// For normal cases, we expect a 2XX status code, or at least not 4XX or 5XX
					if w.Code < 200 || w.Code >= 300 {
						t.Errorf("Expected successful status code (2XX), got %d", w.Code)
					}
				}
			}

			// Additional check for 4XX and 5XX errors
			if (w.Code >= 400 && w.Code <= 599) && !tc.expectFailure {
				t.Errorf("Request failed with status code %d", w.Code)
			}

			// Log the response for debugging
			t.Logf("Response for %s: Status: %d", tc.name, w.Code)
		})
	}
}

// TestMain runs integration tests in a controlled environment
func TestMain(m *testing.M) {
	// Print test header
	fmt.Println("=== Running Integration Tests ===")

	// Run tests
	m.Run()

	// Print test footer
	fmt.Println("=== Integration Tests Complete ===")
}
