package email

import (
	"context"
	"fmt"
	sa "github.com/Nusantara-Muda/scholarship-api"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

// Html ...
var html = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
  <link href="https://fonts.googleapis.com/css2?family=Quicksand:wght@500;700&display=swap" rel="stylesheet" />
  <title>Activate account</title>
  <style>
    h1, h3, p {
      margin: 0;
    }
    h3 {
      margin-top: 32px;
      font-weight: 700;
      font-size: 18px;
      line-height: 24px;
    }
    p {
      font-size: 12px;
      line-height: 16px;
    }
    body {
      font-family: "Quicksand";
      color: #464952;
      background: #EAECF1;
      font-weight: 500;
      padding: 16px;
    }
    button {
      background: #B31E1A;
      color: #FEFFFF;
      padding: 8px 16px;
      font-size: 16px;
      line-height: 24px;
      margin-top: 24px;
      border: none;
      outline: none;
      border-radius: 4px;
	  cursor: pointer;
    }
	a {
	  cursor: pointer;
	}
  </style>
</head>
<body>
  <h1>Bangun Scholarship</h1>
  <h3>Please confirm your email address.</h3>
  <p>Click the button below to activate your account.</p>
  <a href="%s" target="_blank">
    <button type="button">Activate my account</button>
  </a>
</body>
</html>
`

type emailRepo struct {
	mailgunImpl      *mailgun.MailgunImpl
	sender           string
	pathActivateUser string
}

// SendActivateUser ...
func (e emailRepo) SendActivateUser(ctx context.Context, email string) error {
	subject := "Activate User Bangun"
	recipient := email

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", recipient)
	html = fmt.Sprintf(html, e.pathActivateUser)
	message.SetHtml(html)

	_, _, err := e.mailgunImpl.Send(ctx, message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// NewEmailRepository ...
func NewEmailRepository(mailgunImpl *mailgun.MailgunImpl, sender, pathActivateUser string) sa.EmailRepository {
	return emailRepo{
		mailgunImpl:      mailgunImpl,
		sender:           sender,
		pathActivateUser: pathActivateUser,
	}
}
