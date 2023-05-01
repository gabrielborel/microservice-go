package usecases

import "github.com/gabrielborel/microservice-go/internal/order/entity"

type GetTotalOrdersOutputDTO struct {
	Total int
}

type GetTotalOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalOrdersUseCase(repository entity.OrderRepositoryInterface) *GetTotalOrdersUseCase {
	return &GetTotalOrdersUseCase{OrderRepository: repository}
}

func (c *GetTotalOrdersUseCase) Execute() (*GetTotalOrdersOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOrdersOutputDTO{Total: total}, nil
}
