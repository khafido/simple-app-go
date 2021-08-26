package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	_"path/filepath"

	"github.com/gorilla/mux"
	_"github.com/khafido/simple-app-go/api/auth"
	"github.com/khafido/simple-app-go/api/models"
	"github.com/khafido/simple-app-go/api/responses"
	"github.com/khafido/simple-app-go/api/utils/formaterror"
)

// API
func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	product := models.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product.Prepare()
	if errValidator := product.ValidateStruct(); errValidator != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, errValidator)
		return
	}

	// uid, err := auth.ExtractTokenID(r)
	// if err != nil || uid==0{
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	productCreated, err := product.SaveProduct(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, productCreated.ID_Produk))
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Created Success"})
	// responses.JSON(w, http.StatusCreated, productCreated)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	products, err := product.FindAllProducts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)
}

func (server *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id_produk"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	product := models.Product{}
	productGotten, err := product.FindProductByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, productGotten)
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id_produk"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// uid, err := auth.ExtractTokenID(r)
	// if err != nil || uid==0{
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	product := models.Product{}
	err = server.DB.Debug().Model(models.Product{}).Where("id_produk = ?", pid).Take(&product).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Product not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productUpdate := models.Product{}
	err = json.Unmarshal(body, &productUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productUpdate.Prepare()
	if errValidator := productUpdate.ValidateStruct(); errValidator != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, errValidator)
		return
	}

	productUpdated, err := productUpdate.UpdateProduct(server.DB, uint32(pid))
	if err != nil && productUpdated == nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Updated Success"})
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := models.Product{}

	pid, err := strconv.ParseUint(vars["id_produk"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// uid, err := auth.ExtractTokenID(r)
	// if err != nil || uid==0{
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	_, err = product.DeleteProduct(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusCreated, map[string]string{"status":"Deleted Success"})
}
