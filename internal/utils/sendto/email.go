package sendto

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/pkg/loggers"

	"go.uber.org/zap"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var logger = loggers.GetLogger()

func LoadSMTPConfigFromEnv() SMTPConfig {
	cfg := global.Config.SMTP

	host := cfg.Host
	if host == "" {
		host = "smtp.mandrillapp.com"
	}

	port := cfg.Port
	if port == 0 {
		port = 587
	}

	from := cfg.From
	if from == "" {
		from = "no-reply@example.com"
	}

	logger.Info("Loaded SMTP config",
		zap.String("host", host),
		zap.Int("port", port),
		zap.String("username", cfg.Username),
		zap.String("password", cfg.Password),
		zap.String("from", from),
	)

	return SMTPConfig{
		Host:     host,
		Port:     port,
		Username: cfg.Username,
		Password: cfg.Password,
		From:     from,
	}
}

type EmailSender interface {
	SendOTP(ctx context.Context, toEmail, otp string) error
}

type MandrillEmailSender struct {
	cfg SMTPConfig
}

func NewMandrillEmailSender(cfg SMTPConfig) *MandrillEmailSender {
	return &MandrillEmailSender{cfg: cfg}
}

func (m *MandrillEmailSender) SendOTP(ctx context.Context, toEmail, otp string) error {
	subject := "Your OTP Code"
	htmlBody := buildOTPTemplate(otp)

	msg := buildEmailMessage(m.cfg.From, toEmail, subject, htmlBody)

	addr := fmt.Sprintf("%s:%d", m.cfg.Host, m.cfg.Port)
	var auth smtp.Auth
	if m.cfg.Username != "" || m.cfg.Password != "" {
		auth = smtp.PlainAuth("", m.cfg.Username, m.cfg.Password, m.cfg.Host)
	}

	if err := smtp.SendMail(addr, auth, m.cfg.From, []string{toEmail}, msg); err != nil {
		return err
	}

	return nil
}

func buildEmailMessage(from, to, subject, htmlBody string) []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("From: %s\r\n", from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	buf.WriteString("\r\n")
	buf.WriteString(htmlBody)
	return buf.Bytes()
}

// buildOTPTemplate tạo template HTML đơn giản hiển thị OTP.
func buildOTPTemplate(otp string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>Your OTP Code</title>
  </head>
  <body style="font-family: Arial, sans-serif; background-color: #f5f5f5; padding: 24px;">
    <div style="max-width: 480px; margin: 0 auto; background: #ffffff; border-radius: 8px; padding: 24px;">
      <h2 style="color:#333; margin-top:0;">Xác thực đăng ký</h2>
      <p>Xin chào,</p>
      <p>Mã OTP của bạn là:</p>
      <div style="font-size: 28px; font-weight: bold; letter-spacing: 4px; margin: 16px 0;">
        %s
      </div>
      <p>Mã này có hiệu lực trong 5 phút. Nếu bạn không thực hiện yêu cầu này, vui lòng bỏ qua email.</p>
      <p>Trân trọng,<br/>Team Support</p>
    </div>
  </body>
</html>
`, otp)
}
