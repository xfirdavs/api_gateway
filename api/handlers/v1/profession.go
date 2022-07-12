package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfirdavs/api_gateway/genproto/position_service"
)

// Profession godoc
// @ID create-profession
// @Router /v1/profession [POST]
// @Summary create profession
// @Description create profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body position_service.CreateProfessionRequest true "profession"
// @Success 200 {object} models.ResponseModel{data=position_service.Profession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfession(c *gin.Context) {
	var profession position_service.CreateProfessionRequest

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessionService().Create(c.Request.Context(), &profession)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-Profession godoc
// @ID get-profession
// @Router /v1/profession [get]
// @Summary get profession
// @Description get profession
// @Tags profession
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfession(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.ProfessionService().GetAll(
		c.Request.Context(),
		&position_service.GetAllProfessionRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all professions", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// GetID-profession godoc
// @ID GetID-profession
// @Router /v1/profession/:id [get]
// @Summary GetID profession
// @Description GetID profession
// @Tags profession
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} models.ResponseModel{data=position_service.GetByIdProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdProfession(c *gin.Context) {

	resp, err := h.services.ProfessionService().GetById(
		c.Request.Context(),
		&position_service.GetByIdProfessionRequest{
			Id: c.Query("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting  profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// profession godoc
// @ID update-profession
// @Router /v1/profession [PUT]
// @Summary update profession
// @Description update profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body position_service.UpdateProfessionRequest true "profession"
// @Success 200 {object} models.ResponseModel{data=position_service.UpdateProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateProfession(c *gin.Context) {
	var profession position_service.UpdateProfessionRequest

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessionService().Update(c.Request.Context(), &profession)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// profession godoc
// @ID delete-profession
// @Router /v1/profession [DELETE]
// @Summary delete profession
// @Description delete profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body position_service.DeleteProfessionRequest true "profession"
// @Success 200 {object} models.ResponseModel{data=position_service.DeleteProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteProfession(c *gin.Context) {
	var profession position_service.DeleteProfessionRequest

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessionService().Delete(c.Request.Context(), &profession)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}
