package user

import (
	"app/pkg/kernel/email"
	"app/pkg/user/internal/action"
	"app/pkg/user/internal/content"
	"app/pkg/user/internal/resource"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Config struct {
	Tracer trace.Tracer
	DB     *gorm.DB
	Mailer email.Mailer
}

type User struct {
	actionChangePassword      *action.ChangePassword
	actionForgotPassword      *action.ForgotPassword
	actionResetPassword       *action.ResetPassword
	actionSignInByEmail       *action.SignInByEmail
	actionSignOut             *action.SignOut
	actionSignUpByEmail       *action.SignUpByEmail
	actionVerifySignUpByEmail *action.VerifySignUpByEmail
	contentCurrentSession     *content.CurrentSession
	resourceList              *resource.List
	resourceOne               *resource.One
}

func (e *User) Fiber() *fiber.App {
	app := fiber.New()
	app.Get("/list", e.resourceList.Get)
	app.Get("/one", e.resourceOne.Get)
	app.Patch("/one", e.resourceOne.Patch)
	app.Delete("/one", e.resourceOne.Delete)

	action := app.Group("/action")
	action.Post("/change-password", e.actionChangePassword.Do)
	action.Post("/forgot-password", e.actionForgotPassword.Do)
	action.Post("/reset-password", e.actionResetPassword.Do)
	action.Post("/sign-in-by-email", e.actionSignInByEmail.Do)
	action.Post("/sign-out", e.actionSignOut.Do)
	action.Post("/sign-up-by-email", e.actionSignUpByEmail.Do)
	action.Post("/verify-sign-up-by-email", e.actionVerifySignUpByEmail.Do)

	content := app.Group("/content")
	content.Get("/current-session", e.contentCurrentSession.Fetch)

	return app
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(User), "*"),
	wire.Struct(new(action.ChangePassword), "*"),
	wire.Struct(new(action.ForgotPassword), "*"),
	wire.Struct(new(action.ResetPassword), "*"),
	wire.Struct(new(action.SignInByEmail), "*"),
	wire.Struct(new(action.SignOut), "*"),
	wire.Struct(new(action.SignUpByEmail), "*"),
	wire.Struct(new(action.VerifySignUpByEmail), "*"),
	wire.Struct(new(content.CurrentSession), "*"),
	wire.Struct(new(resource.List), "*"),
	wire.Struct(new(resource.One), "*"),
	wire.FieldsOf(new(Config), "Tracer", "DB", "Mailer"),
)
