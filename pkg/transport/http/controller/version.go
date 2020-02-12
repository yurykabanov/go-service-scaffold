package controller

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/yurykabanov/service-scaffold/pkg/version"
)

type VersionController struct {
}

func NewVersionController() *VersionController {
	return &VersionController{}
}

type versionResponse struct {
	Build     string `json:"build" example:"0123456789abcdef0123456789abcdef01234567"`
	Version   string `json:"version" example:"1.0.0"`
	BuildDate string `json:"build_date" example:"2020-02-03 12:34:56+00:00"`
}

// Version godoc
// @Router /version [get]
// @Summary Service version
// @Description Returns service build and version information
// @ID version.get
// @Tags about
// @Produce json
// @Success 200 {object} versionResponse
func (ctrl *VersionController) Version(ctx echo.Context) error {
	sp := jaegertracing.CreateChildSpan(ctx, "Child span for additional processing")
	defer sp.Finish()

	return ctx.JSON(200, versionResponse{
		Build:     version.Build,
		Version:   version.Version,
		BuildDate: version.BuildDate,
	})
}
