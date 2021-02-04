package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaswilliameufrasio/imersao/codepix-go/app/usecase"
	"github.com/lucaswilliameufrasio/imersao/codepix-go/infra/repository"
)

func MakeTransactionUseCase(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDB{DB: database}
	transactionRepository := repository.TransactionRepositoryDB{DB: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
