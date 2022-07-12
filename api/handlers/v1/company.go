package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfirdavs/api_gateway/genproto/company_service"
)

// company godoc
// @ID create-company
// @Router /v1/company [POST]
// @Summary create company
// @Description create company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.CreateCompanyRequest true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCompany(c *gin.Context) {
	var company company_service.CreateCompanyRequest

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Create(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-company godoc
// @ID get-company
// @Router /v1/company [get]
// @Summary get company
// @Description get company
// @Tags company
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=company_service.GetAllCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCompany(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CompanyService().GetAll(
		c.Request.Context(),
		&company_service.GetAllCompanyRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all companies", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// GetID-company godoc
// @ID GetID-company
// @Router /v1/company/:id [get]
// @Summary GetID company
// @Description GetID company
// @Tags company
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} models.ResponseModel{data=company_service.GetByIdCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdCompany(c *gin.Context) {

	resp, err := h.services.CompanyService().GetById(
		c.Request.Context(),
		&company_service.GetByIdCompanyRequest{
			Id: c.Query("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting  company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// company godoc
// @ID update-company
// @Router /v1/company [PUT]
// @Summary update company
// @Description update company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.UpdateCompanyRequest true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCompany(c *gin.Context) {
	var company company_service.UpdateCompanyRequest

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Update(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// company godoc
// @ID delete-company
// @Router /v1/company [DELETE]
// @Summary delete company
// @Description delete company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.DeleteCompanyRequest true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.DeleteCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCompany(c *gin.Context) {
	var company company_service.DeleteCompanyRequest

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Delete(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}
