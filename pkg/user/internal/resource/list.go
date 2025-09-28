package resource

import (
	"app/pkg/user/internal/user"

	"github.com/Arsfiqball/talkback"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type List struct {
	Tracer trace.Tracer
	DB     *gorm.DB
}

func (r *List) Get(c *fiber.Ctx) error {
	ctx, span := r.Tracer.Start(c.UserContext(), "user/internal/resource.List.Get")
	defer span.End()

	qs := string(c.Request().URI().QueryString())

	query, err := talkback.FromQueryString(qs)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	users := user.Users{Query: query}

	if err := users.LoadFrom(r.DB.WithContext(ctx)); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}
