package action

import (
	"app/pkg/kernel/email"
	"app/pkg/kernel/typex"
	"app/pkg/user/internal/user"
	"embed"
	"fmt"
	"html/template"

	"github.com/Arsfiqball/codec/flame"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type SignUpByEmail struct {
	Tracer        trace.Tracer
	DB            *gorm.DB
	Mailer        email.Mailer
	GomailChannel chan *gomail.Message
}

func (a *SignUpByEmail) Do(c *fiber.Ctx) error {
	ctx, span := a.Tracer.Start(c.UserContext(), "user/internal/action.SignUpByEmail.Do")
	defer span.End()

	// prepare new user
	u := user.NewUser()

	var payload struct {
		Email    typex.Email    `json:"email"`
		Password typex.Password `json:"password"`
	}

	// parsing payload
	if err := c.BodyParser(&payload); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	// validate each fields
	errd := flame.Data{
		"email":    payload.Email.Validate(),
		"password": payload.Password.Validate(),
	}.RemoveNil()

	// if there are invalid fields, response 422
	if !errd.IsEmpty() {
		return c.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"errors": errd})
	}

	// hash the password
	hash, err := payload.Password.Hash()
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// attach payloads & hashed password to new user
	u.Email = payload.Email
	u.PasswordHash = hash

	// save the new user to database
	if err := u.SaveTo(a.DB.WithContext(ctx)); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	m := email.Message{
		To:      u.Email.String(),
		Subject: "Please verify your email address",
		Body: fmt.Sprintf(`
		<p>Thank you for signing up with your email address: %s</p>
		<p>Please verify your email address by clicking the link below:</p>
		<a>Verify Email</a>
		<p>If you did not sign up for this account, please ignore this email.</p>
	`, u.Email.String()),
	}

	_, err = template.ParseFS(verifyEmailTemplate, "verify_email.html")
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	if err := a.Mailer.Send(ctx, m); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"info": "success created user using email address, please verify the email",
		})
}

//go:embed verify_email.html
var verifyEmailTemplate embed.FS
