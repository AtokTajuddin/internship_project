package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"project_virtual_internship_evermos/internal/package/entity"
)

type RegionRepository interface {
	GetProvinces() ([]entity.Province, error)
	GetRegencies(provinceID string) ([]entity.Regency, error)
	GetDistricts(regencyID string) ([]entity.District, error)
	GetVillages(districtID string) ([]entity.Village, error)
}

type regionRepository struct {
	baseURL string
}

func NewRegionRepository() RegionRepository {
	return &regionRepository{
		baseURL: "https://www.emsifa.com/api-wilayah-indonesia/api",
	}
}

func (r *regionRepository) get(endpoint string, target interface{}) error {
	resp, err := http.Get(r.baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

func (r *regionRepository) GetProvinces() ([]entity.Province, error) {
	var provinces []entity.Province
	err := r.get("/provinces.json", &provinces)
	return provinces, err
}

func (r *regionRepository) GetRegencies(provinceID string) ([]entity.Regency, error) {
	var regencies []entity.Regency
	err := r.get(fmt.Sprintf("/regencies/%s.json", provinceID), &regencies)
	return regencies, err
}

func (r *regionRepository) GetDistricts(regencyID string) ([]entity.District, error) {
	var districts []entity.District
	err := r.get(fmt.Sprintf("/districts/%s.json", regencyID), &districts)
	return districts, err
}

func (r *regionRepository) GetVillages(districtID string) ([]entity.Village, error) {
	var villages []entity.Village
	err := r.get(fmt.Sprintf("/villages/%s.json", districtID), &villages)
	return villages, err
}
