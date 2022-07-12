package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xfirdavs/api_gateway/genproto/position_service"
)

// position godoc
// @ID create-position
// @Router /v1/position [POST]
// @Summary create position
// @Description create position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.CreatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreatePosition(c *gin.Context) {
	var position position_service.CreatePositionRequest

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().Create(c.Request.Context(), &position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-position godoc
// @ID get-position
// @Router /v1/position [get]
// @Summary get position
// @Description get position
// @Tags position
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllPositionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllPosition(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionService().GetAll(
		c.Request.Context(),
		&position_service.GetAllPositionRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// GetID-position godoc
// @ID GetID-position
// @Router /v1/position/:id [get]
// @Summary GetID position
// @Description GetID position
// @Tags position
// @Accept json
// @Produce json
// @Param id query position_service.PositionId false "id"
// @Success 200 {object} models.ResponseModel{data=position_service.Position} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdposition(c *gin.Context) {

	resp, err := h.services.PositionService().GetById(
		c.Request.Context(),
		&position_service.PositionId{
			Id: c.Query("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting  position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// position godoc
// @ID update-position
// @Router /v1/position [PUT]
// @Summary update position
// @Description update position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.UpdatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Updateposition(c *gin.Context) {
	var position position_service.UpdatePositionRequest

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().Update(c.Request.Context(), &position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// position godoc
// @ID delete-position
// @Router /v1/position [DELETE]
// @Summary delete position
// @Description delete position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.PositionId true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Deleteposition(c *gin.Context) {
	var position position_service.PositionId

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().Delete(c.Request.Context(), &position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
