package handlers

import (
	"encoding/json"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"myapi/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListItens - Lista todos os itens
func ListItems(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewItemRepository()
	items, err := repository.ListAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Erro ao listar os itens")
		return
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao codificar a resposta")
		return
	}
}

// GetItem - Busca um item por ID (via rota: /item/{id})
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ID não fornecido")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Item não encontrado")
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao codificar a resposta")
		return
	}
}

// GetItemByCode - Busca um item pelo campo "codigo" (via rota: /item/codigo/{codigo})
func GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["codigo"]

	if code == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Código não fornecido")
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByCode(code)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Item não encontrado")
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao codificar a resposta")
		return
	}
}

// CreateItem - Cria um novo item (envie JSON via POST)
func CreateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	repository := repositories.NewItemRepository()
	createdItem, err := repository.Create(item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao criar o item")
		return
	}
	err = json.NewEncoder(w).Encode(createdItem)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao codificar a resposta")
		return
	}
}

// UpdateItem - Atualiza um item existente (envie JSON via PUT, com o campo id preenchido)
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Update(item); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao atualizar o item")
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao codificar a resposta")
		return
	}
}

// DeleteItem - Deleta um item por ID (via rota: /item/{id})
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "ID não fornecido")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Delete(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Erro ao deletar o item")
		return
	}
	if _, err := w.Write([]byte("Item deletado com sucesso")); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
}
