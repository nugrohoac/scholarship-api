package email

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/Nusantara-Muda/scholarship-api/src/business/entity"
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

	htmlAwardeeConfirmation = `
	<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">

  <title>Awardee confirmation</title>
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
  <h2>Congratulation, you made it!</h2>
  <p style="margin: 0;font-size: 12px;line-height: 16px;font-weight: 500;">
    We are pleased to inform that you have been choosen as awardee for
    <span style="line-height: 20px;font-weight: bold;font-size: 16px;">%s</span>
    scholarship.
  </p>
  <br>
  <p class="alert" style="margin: 0;font-size: 12px;line-height: 16px;font-weight: 500;color: #878C96;">
    Please confrm your membership (by clicking the button below) as awardee within
    <span style="line-height: 20px;font-weight: bold;font-size: 12px;color: black;">
      72 hours
    </span>
    or your membership will be canceled automatically by the system.
  </p>
  <a href="%s" target="_blank">
    <button type="button" style="background: #B31E1A;color: #FEFFFF;padding: 8px 16px;font-size: 16px;line-height: 24px;margin-top: 24px;border: none;outline: none;border-radius: 4px;">Yes, I confirm</button>
  </a>
</body>
</html>
	`

	htmlSuccessConfirmationAwardee = `
	<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous">

  <title>Awardee confirmation notification</title>
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
  <h2>Confirmation by awardee</h2>
  <p style="margin: 0;font-size: 12px;line-height: 16px;font-weight: 500;">
    Kindly notify that your awardee
    <span style="line-height: 20px;font-weight: bold;font-size: 14px;">%s</span>
    for
    <span style="line-height: 20px;font-weight: bold;font-size: 16px;">%s</span>
    scholarship
    has confirmed their membership.
  </p>
  <br>
  <p class="alert" style="margin: 0;font-size: 12px;line-height: 16px;font-weight: 500;">
    Best Regards,<br>
    <span style="line-height: 20px;font-weight: bold;font-size: 12px;color: black;">
      Bangun Team
    </span>
  </p>
</body>
</html>
`
)

type emailRepo struct {
	mailgunImpl                   *mailgun.MailgunImpl
	sender                        string
	pathActivateUser              string
	pathForgotPassword            string
	pathForgotPasswordBackoffice  string
	pathNotifyFundingConfirmation string
	pathConfirmationByAwardee     string
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
func (e emailRepo) SendForgotPassword(ctx context.Context, email, token, userType string) error {
	subject := "Reset Password"
	recipient := email
	pathForgotPassword := e.pathForgotPassword

	if strings.ToLower(userType) == "admin" {
		pathForgotPassword = e.pathForgotPasswordBackoffice
	}

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", recipient)
	path := pathForgotPassword + "?token=" + token
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
	subject := "Bangun Awardee Funding Confirmation"

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

// BlazingToAwardee .
// rezabaguspermana.rbp@gmail.com .
// reza.bagus@tanihub.com
func (e emailRepo) BlazingToAwardee(ctx context.Context, mapEmailToken map[string]string, scholarship entity.Scholarship) error {
	subject := "Bangun Awardee Funding"

	if len(mapEmailToken) > 0 {
		for email, token := range mapEmailToken {
			message := e.mailgunImpl.NewMessage(e.sender, subject, "", email)
			// it can be replaced with fmt.sprintf
			path := e.pathConfirmationByAwardee + "?token=" + token + "&scholarship_id=" + strconv.Itoa(int(scholarship.ID))
			// html copy, if sending to send email more, it will more extra string
			_html := htmlAwardeeConfirmation
			_html = fmt.Sprintf(_html, scholarship.Name, path)
			message.SetHtml(_html)

			_, _, err := e.mailgunImpl.Send(ctx, message)
			if err != nil {
				logrus.Error("failed sending email to : ", email, err)
			}
		}
	}

	return nil
}

// ConfirmToSponsor .
func (e emailRepo) ConfirmToSponsor(ctx context.Context, emailSponsor, studentName string, scholarshipName string) error {
	subject := "Success Confirmation Awardee"
	recipient := emailSponsor

	message := e.mailgunImpl.NewMessage(e.sender, subject, "", recipient)
	// html copy, if sending to send email more, it will more extra string
	_html := htmlSuccessConfirmationAwardee
	_html = fmt.Sprintf(_html, studentName, scholarshipName)
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
	pathForgotPasswordBackoffice,
	pathNotifyFundingConfirmation,
	pathConfirmationByAwardee string) business.EmailRepository {
	return emailRepo{
		mailgunImpl:                   mailgunImpl,
		sender:                        sender,
		pathActivateUser:              pathActivateUser,
		pathForgotPassword:            pathForgotPassword,
		pathForgotPasswordBackoffice:  pathForgotPasswordBackoffice,
		pathNotifyFundingConfirmation: pathNotifyFundingConfirmation,
		pathConfirmationByAwardee:     pathConfirmationByAwardee,
	}
}
