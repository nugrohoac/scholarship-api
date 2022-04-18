package email

import (
	"context"
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
	"strconv"
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
	<style>
		/* latin */
		@font-face {
			font-family: 'Quicksand';
			font-style: normal;
			font-weight: 500;
			font-display: swap;
			src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
		}
		/* latin */
		@font-face {
			font-family: 'Quicksand';
			font-style: normal;
			font-weight: 700;
			font-display: swap;
			src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
		}  
	</style>
</head>
<body style="font-family: 'Quicksand', 'open sans', 'helvetica neue', sans-serif; color: #464952;background: #EAECF1;font-weight: 500;padding: 16px;">
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
	
	<title>Forgot Password</title>
	<style>
		/* latin */
		@font-face {
			font-family: 'Quicksand';
			font-style: normal;
			font-weight: 500;
			font-display: swap;
			src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
		}
		/* latin */
		@font-face {
			font-family: 'Quicksand';
			font-style: normal;
			font-weight: 700;
			font-display: swap;
			src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
		}  
	</style>
</head>
<body style="font-family: 'Quicksand', 'open sans', 'helvetica neue', sans-serif; color: #464952;background: #EAECF1;font-weight: 500;padding: 16px;">
	<h1 style="margin: 0;">Bangun Scholarship</h1>
	<h3 style="margin: 0;margin-top: 32px;font-weight: 700;font-size: 18px;line-height: 24px;">Forgot password</h3>
	<p style="margin: 0;font-size: 12px;line-height: 16px;">You have been asked to reset your password account. To proceed please click the button below to reset your password.</p><br>
	<p style="margin: 0;font-size: 12px;line-height: 16px;">If you never ask this request, we hope you ignore this email.</p>
	<a href="%s" target="_blank" style="cursor: pointer;">
		<button type="button" style="background: #B31E1A;color: #FEFFFF;padding: 8px 16px;font-size: 16px;line-height: 24px;margin-top: 24px;border: none;outline: none;border-radius: 4px;cursor: pointer;">Reset Password</button>
	</a>
</body>
</html>
`

	htmlNotifyFundingConfirmation = `
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">

  <title>Activate account</title>
  <style>
	/* latin */
	@font-face {
	  font-family: 'Quicksand';
	  font-style: normal;
	  font-weight: 500;
	  font-display: swap;
	  src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
	}
	/* latin */
	@font-face {
	  font-family: 'Quicksand';
	  font-style: normal;
	  font-weight: 700;
	  font-display: swap;
	  src: url(https://fonts.gstatic.com/s/quicksand/v24/6xKtdSZaM9iE8KbpRA_hK1QN.woff2) format('woff2');
	}  
  </style>
</head>
<body style="font-family: 'Quicksand', 'open sans', 'helvetica neue', sans-serif; color: #464952;background: #EAECF1;font-weight: 500;padding: 16px;">
  <img src="https://s3.ap-southeast-3.amazonaws.com/cdn.stading.bangun.app/documents//1649429019215" width="138" height="40">
  <h3 style="margin: 0;margin-top: 16px;font-weight: 700;font-size: 18px;line-height: 24px;">Please confirm your awardee</h3>
  <ol style="padding-left: 18px;font-size: 12px;line-height: 16px;">
	%s
  </ol>
  <p style="margin: 0;font-size: 12px;line-height: 16px;color: #747793;font-weight: 400;">By confirming this email, we will sent notify to ask for confirmation on the awardees dashboard to transfer the money.</p>
  <a href="%s" target="_blank">
	<button type="button" style="background: #B31E1A;color: #FEFFFF;padding: 8px 16px;font-size: 16px;line-height: 24px;margin-top: 24px;border: none;outline: none;border-radius: 4px;">Yes, I confirm</button>
  </a>
</body>
</html>
	`
)

type emailRepo struct {
	mailgunImpl                   *mailgun.MailgunImpl
	sender                        string
	pathActivateUser              string
	pathForgotPassword            string
	pathNotifyFundingConfirmation string
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

// NotifyFundingConformation .
func (e emailRepo) NotifyFundingConformation(ctx context.Context, email, token string, scholarshipID int64, data string) error {
	subject := "Confirm Your Awardee"

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", email)
	// it can be replaced with fmt.sprintf
	path := e.pathNotifyFundingConfirmation + "?token=" + token + "&scholarship_id=" + strconv.Itoa(int(scholarshipID))
	// html copy, if sending to send email more, it will more extra string
	_html := htmlNotifyFundingConfirmation
	_html = fmt.Sprintf(_html, data, path)
	message.SetHtml(_html)

	_, _, err := e.mailgunImpl.Send(ctx, message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// NewEmailRepository ...
func NewEmailRepository(
	mailgunImpl *mailgun.MailgunImpl,
	sender,
	pathActivateUser,
	pathForgotPassword,
	pathNotifyFundingConfirmation string) business.EmailRepository {
	return emailRepo{
		mailgunImpl:                   mailgunImpl,
		sender:                        sender,
		pathActivateUser:              pathActivateUser,
		pathForgotPassword:            pathForgotPassword,
		pathNotifyFundingConfirmation: pathNotifyFundingConfirmation,
	}
}
