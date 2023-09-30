package cerpen

import (
	"net/http"
	"prakerja11/configs"
	"prakerja11/models/base"
	cerpenmodel "prakerja11/models/cerpen"

	"github.com/labstack/echo/v4"
)

func CreateCerpen(c echo.Context) error {
	var cerpen cerpenmodel.Cerpen
	c.Bind(&cerpen)

	result := configs.DB.Create(&cerpen)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success register",
		Data:    cerpen,
	})
}
