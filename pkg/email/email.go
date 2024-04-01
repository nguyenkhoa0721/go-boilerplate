package email

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go-boilerplate/config"
	"go-boilerplate/pkg/constant"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/otel"
	"go-boilerplate/pkg/uuid"
	"html/template"
	"strings"
)

type EmailService struct {
	config *config.Config
	uuid   *uuid.Uuid
	client *sendgrid.Client
}

func NewEmailService(config *config.Config, uuid *uuid.Uuid) *EmailService {
	return &EmailService{
		config: config,
		uuid:   uuid,
		client: sendgrid.NewSendClient(config.Email.Key),
	}
}

func (s *EmailService) SendEmail(ctx context.Context, toEmail, subject string, emailBody string, isGroup bool) (bool, error) {
	ctx, span := otel.Start(ctx, "email:domain:email:SendEmail", map[string]interface{}{
		"toEmail":   toEmail,
		"subject":   subject,
		"emailBody": emailBody,
		"isGroup":   isGroup,
	})
	defer span.End()

	fromEmail := s.config.Email.FromEmail
	if !isGroup {
		fromEmailArr := strings.Split(fromEmail, "@")
		randomID := fmt.Sprint(s.uuid.GenerateUuid(ctx, constant.UUID_UTILITY))
		fromEmail = fromEmailArr[0] + "+" + randomID + "@" + fromEmailArr[1]
	}

	from := mail.NewEmail(s.config.Email.FromName, fromEmail)
	to := mail.NewEmail("", toEmail)

	content := mail.NewContent("text/html", emailBody)
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)

	message := mail.NewV3MailInit(from, subject, to, content)
	message.AddPersonalizations(personalization)

	emailRes, err := s.client.Send(message)
	if err != nil {
		return false, err
	} else {
		logger.Info(ctx).Msgf("Email sent to %s with status code %s", toEmail, emailRes.Body)
	}
	return true, nil
}

func (s *EmailService) LoadTemplateByName(ctx context.Context, name string) (*template.Template, error) {
	ctx, span := otel.Start(ctx, "email:domain:email:LoadTemplateByName", map[string]interface{}{
		"name": name,
	})
	defer span.End()

	templatePath := strings.Join([]string{s.config.Email.TemplatePath, name}, "/")
	tmpl, err := template.ParseFiles(templatePath)
	return tmpl, err
}

func (s *EmailService) SendEmailTemplate(ctx context.Context, toEmail, subject, templateName string, data interface{}, isGroup bool) (bool, error) {
	ctx, span := otel.Start(ctx, "email:domain:email:SendEmailTemplate", map[string]interface{}{
		"toEmail":  toEmail,
		"subject":  subject,
		"template": templateName,
		"data":     data,
		"isGroup":  isGroup,
	})
	defer span.End()

	tmpl, err := s.LoadTemplateByName(ctx, templateName)
	if err != nil {
		return false, err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return false, err
	}
	return s.SendEmail(ctx, toEmail, subject, tpl.String(), isGroup)
}
