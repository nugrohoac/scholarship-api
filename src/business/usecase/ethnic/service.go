package ethnic

import (
	"context"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
)

type ethnicService struct {
	ethnicRepository business.EthnicRepository
}

// Fetch .
func (e ethnicService) Fetch(ctx context.Context) ([]entity.Ethnic, error) {
	return e.ethnicRepository.Fetch(ctx)
}

// NewEthnicService .
func NewEthnicService(ethnicRepository business.EthnicRepository) business.EthnicService {
	return ethnicService{ethnicRepository: ethnicRepository}
}
