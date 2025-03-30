package repository

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"project_virtual_internship_evermos/internal/package/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ApiRepoTestSuite struct {
	suite.Suite
	server *httptest.Server
	repo   RegionRepository
}

func (s *ApiRepoTestSuite) SetupSuite() {
	// Create a mock HTTP server that returns predefined responses
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Match request path and return appropriate mock data
		switch r.URL.Path {
		case "/provinces.json":
			// Mock provinces data
			provinces := []entity.Province{
				{ID: "11", Name: "ACEH"},
				{ID: "12", Name: "SUMATERA UTARA"},
			}
			json.NewEncoder(w).Encode(provinces)

		case "/regencies/11.json":
			// Mock regencies data for province 11 (ACEH)
			regencies := []entity.Regency{
				{ID: "1101", ProvinceID: "11", Name: "KABUPATEN SIMEULUE"},
				{ID: "1102", ProvinceID: "11", Name: "KABUPATEN ACEH SINGKIL"},
			}
			json.NewEncoder(w).Encode(regencies)

		case "/districts/1101.json":
			// Mock districts data for regency 1101 (KABUPATEN SIMEULUE)
			districts := []entity.District{
				{ID: "110101", RegencyID: "1101", Name: "TEUPAH SELATAN"},
				{ID: "110102", RegencyID: "1101", Name: "SIMEULUE TIMUR"},
			}
			json.NewEncoder(w).Encode(districts)

		case "/villages/110101.json":
			// Mock villages data for district 110101 (TEUPAH SELATAN)
			villages := []entity.Village{
				{ID: "1101012001", DistrictID: "110101", Name: "LATIUNG"},
				{ID: "1101012002", DistrictID: "110101", Name: "LABUHAN BAJAU"},
			}
			json.NewEncoder(w).Encode(villages)

		case "/regencies/invalid.json":
			// Test error handling for invalid province ID
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Province not found"}`))

		default:
			// Handle unexpected paths
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	// Create a new repository instance that uses our mock server
	s.repo = &regionRepository{
		baseURL: s.server.URL,
	}
}

func (s *ApiRepoTestSuite) TearDownSuite() {
	// Close the test server when done
	s.server.Close()
}

func (s *ApiRepoTestSuite) TestGetProvinces() {
	// Test getting provinces
	provinces, err := s.repo.GetProvinces()

	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(provinces))
	assert.Equal(s.T(), "11", provinces[0].ID)
	assert.Equal(s.T(), "ACEH", provinces[0].Name)
	assert.Equal(s.T(), "12", provinces[1].ID)
	assert.Equal(s.T(), "SUMATERA UTARA", provinces[1].Name)
}

func (s *ApiRepoTestSuite) TestGetRegencies() {
	// Test getting regencies for a valid province ID
	regencies, err := s.repo.GetRegencies("11")

	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(regencies))
	assert.Equal(s.T(), "1101", regencies[0].ID)
	assert.Equal(s.T(), "11", regencies[0].ProvinceID)
	assert.Equal(s.T(), "KABUPATEN SIMEULUE", regencies[0].Name)
}

func (s *ApiRepoTestSuite) TestGetDistricts() {
	// Test getting districts for a valid regency ID
	districts, err := s.repo.GetDistricts("1101")

	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(districts))
	assert.Equal(s.T(), "110101", districts[0].ID)
	assert.Equal(s.T(), "1101", districts[0].RegencyID)
	assert.Equal(s.T(), "TEUPAH SELATAN", districts[0].Name)
}

func (s *ApiRepoTestSuite) TestGetVillages() {
	// Test getting villages for a valid district ID
	villages, err := s.repo.GetVillages("110101")

	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(villages))
	assert.Equal(s.T(), "1101012001", villages[0].ID)
	assert.Equal(s.T(), "110101", villages[0].DistrictID)
	assert.Equal(s.T(), "LATIUNG", villages[0].Name)
}

func (s *ApiRepoTestSuite) TestErrorHandling() {
	// Test error handling for invalid province ID
	regencies, err := s.repo.GetRegencies("invalid")

	// Assertions
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "API returned status: 404")
	assert.Empty(s.T(), regencies)
}

// Run the test suite
func TestApiRepoSuite(t *testing.T) {
	suite.Run(t, new(ApiRepoTestSuite))
}
