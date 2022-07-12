package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfirdavs/api_gateway/genproto/position_service"
)

// attribute godoc
// @ID create-attribute
// @Router /v1/attribute [POST]
// @Summary create attribute
// @Description create attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.CreateAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateAttribute(c *gin.Context) {
	var attribute position_service.CreateAttributeRequest

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().Create(c.Request.Context(), &attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAll-attribute godoc
// @ID GetAll-attribute
// @Router /v1/attribute [GET]
// @Summary GetAll attribute
// @Description GetAll attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllAttribute(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.AttributeService().GetAll(
		c.Request.Context(),
		&position_service.GetAllAttributeRequest{
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

// GetID-attribute godoc
// @ID GetID-attribute
// @Router /v1/attribute/:id [get]
// @Summary GetID attribute
// @Description GetID attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} models.ResponseModel{data=position_service.GetByIdAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdattribute(c *gin.Context) {

	resp, err := h.services.AttributeService().GetById(
		c.Request.Context(),
		&position_service.GetByIdAttributeRequest{
			Id: c.Query("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting  attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// update attribute godoc
// @ID update-attribute
// @Router /v1/attribute [PUT]
// @Summary update attribute
// @Description update attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.UpdateAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.UpdateAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Updateattribute(c *gin.Context) {
	var attribute position_service.UpdateAttributeRequest

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().Update(c.Request.Context(), &attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// attribute godoc
// @ID delete-attribute
// @Router /v1/attribute [DELETE]
// @Summary delete attribute
// @Description delete attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.DeleteAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.DeleteAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Deleteattribute(c *gin.Context) {
	var attribute position_service.DeleteAttributeRequest

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().Delete(c.Request.Context(), &attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}
