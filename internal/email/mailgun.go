package email

import (
	"context"
	"fmt"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

// Html ...
var (
	htmlActivateUser = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">
  
  <title>Activate account</title>
  
</head>
<body style="font-family: &quot;Quicksand&quot;;color: #464952;background: #EAECF1;font-weight: 500;padding: 16px;">
  <h1 style="margin: 0;">Bangun Scholarship</h1>
  <h3 style="margin: 0;margin-top: 32px;font-weight: 700;font-size: 18px;line-height: 24px;">Please confirm your email address.</h3>
  <p style="margin: 0;font-size: 12px;line-height: 16px;">Click the button below to activate your account.</p>
  <a href="%s" target="_blank" style="cursor: pointer;">
    <button type="button" style="background: #B31E1A;color: #FEFFFF;padding: 8px 16px;font-size: 16px;line-height: 24px;margin-top: 24px;border: none;outline: none;border-radius: 4px;cursor: pointer;">Activate my account</button>
  </a>
</body>
</html>
`

	htmlForgotPassword = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">
  
  <title>Reset Password</title>
  
</head>
<body style="font-family: &quot;Quicksand&quot;;color: #464952;background: #EAECF1;font-weight: 500;padding: 16px;">
  <h1 style="margin: 0;">Bangun Scholarship</h1>
  <h3 style="margin: 0;margin-top: 32px;font-weight: 700;font-size: 18px;line-height: 24px;">Please confirm your email address.</h3>
  <p style="margin: 0;font-size: 12px;line-height: 16px;">Click the button below to reset password account.</p>
  <a href="%s" target="_blank" style="cursor: pointer;">
    <button type="button" style="background: #B31E1A;color: #FEFFFF;padding: 8px 16px;font-size: 16px;line-height: 24px;margin-top: 24px;border: none;outline: none;border-radius: 4px;cursor: pointer;">Reset password my account</button>
  </a>
</body>
</html>
`
)

type emailRepo struct {
	mailgunImpl        *mailgun.MailgunImpl
	sender             string
	pathActivateUser   string
	pathForgotPassword string
}

// SendActivateUser ...
func (e emailRepo) SendActivateUser(ctx context.Context, email, token string) error {
	subject := "Activate User Bangun"
	recipient := email

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", recipient)
	path := e.pathActivateUser + "?token=" + token
	// html copy, if sending to send email more, it will more extra string
	_html := htmlActivateUser
	_html = fmt.Sprintf(_html, path)
	message.SetHtml(_html)

	_, _, err := e.mailgunImpl.Send(ctx, message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// SendForgotPassword ...
func (e emailRepo) SendForgotPassword(ctx context.Context, email, token string) error {
	subject := "Reset Password"
	recipient := email

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", recipient)
	path := e.pathForgotPassword + "?token=" + token
	// html copy, if sending to send email more, it will more extra string
	_html := htmlForgotPassword
	_html = fmt.Sprintf(_html, path)
	message.SetHtml(_html)

	_, _, err := e.mailgunImpl.Send(ctx, message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// NewEmailRepository ...
func NewEmailRepository(mailgunImpl *mailgun.MailgunImpl, sender, pathActivateUser, pathForgotPassword string) sa.EmailRepository {
	return emailRepo{
		mailgunImpl:        mailgunImpl,
		sender:             sender,
		pathActivateUser:   pathActivateUser,
		pathForgotPassword: pathForgotPassword,
	}
}
