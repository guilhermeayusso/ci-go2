package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateItem(item *models.Item) error {
	if item.Nome == "" {
		return errors.New("o nome do item é obrigatório")
	}
	if len(item.Codigo) != 6 {
		return errors.New("o código do item deve ter 6 caracteres")
	}
	if item.Preco <= 0 {
		return errors.New("o preço do item não pode ser zero ou negativo")
	}
	if item.Quantidade <= 0 {
		return errors.New("a quantidade do item não pode ser zero ou negativa")
	}
	return nil
}
