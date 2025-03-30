package usecase

import (
	"project_virtual_internship_evermos/internal/package/entity"
	"project_virtual_internship_evermos/internal/package/repository"
)

type RegionUsecase interface {
	GetProvinces() ([]entity.Province, error)
	GetRegencies(provinceID string) ([]entity.Regency, error)
	GetDistricts(regencyID string) ([]entity.District, error)
	GetVillages(districtID string) ([]entity.Village, error)
}

type regionUsecase struct {
	regionRepository repository.RegionRepository
}

func NewRegionUsecase(regionRepository repository.RegionRepository) RegionUsecase {
	return &regionUsecase{
		regionRepository: regionRepository,
	}
}

func (u *regionUsecase) GetProvinces() ([]entity.Province, error) {
	return u.regionRepository.GetProvinces()
}

func (u *regionUsecase) GetRegencies(provinceID string) ([]entity.Regency, error) {
	return u.regionRepository.GetRegencies(provinceID)
}

func (u *regionUsecase) GetDistricts(regencyID string) ([]entity.District, error) {
	return u.regionRepository.GetDistricts(regencyID)
}

func (u *regionUsecase) GetVillages(districtID string) ([]entity.Village, error) {
	return u.regionRepository.GetVillages(districtID)
}
