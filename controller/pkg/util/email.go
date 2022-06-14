package util

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// SendProofRequestEmail sends an email.
func SendProofRequestEmail(recipientEmail, qrImgName string, qrCode []byte) error {
	img, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./" + qrImgName + ".png")
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		log.Println(err)
	}

	htmlBody :=
		`<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Iamza Cornerstone Verifier</title>
	</head>
	<body>
		<p>Hi</p>
		<p>Please scan the following <b>QR Code</b> with your wallet app to verify your <b>Cornerstone Credential</b>: <br/><br/><img src="cid:` + qrImgName + `.png" alt="QR Code" /></p>
		<p>Regards<br/>Iamza Cornerstone Verifier</p>
	</body>
</html>`

	server := mail.NewSMTPClient()

	//SMTP Server
	server.Host = "smtp.office365.com"
	server.Port = 587
	server.Username = "aws_iamzanet@bankservafrica.com"
	server.Password = "K5fEoXa8"
	server.Encryption = mail.EncryptionSTARTTLS
	server.Authentication = mail.AuthLogin

	//Variable to keep alive connection
	server.KeepAlive = true

	//Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	//Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	//SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	//New email simple html with inline and CC
	email := mail.NewMSG()

	email.SetFrom("aws_iamzanet@bankservafrica.com").
		AddTo(recipientEmail).
		// AddCc("otherto@example.com").
		SetSubject("Verify your Cornerstone Credential")

	email.SetBody(mail.TextHTML, htmlBody)

	email.Attach(&mail.File{FilePath: "./" + qrImgName + ".png", Inline: true})

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}

	return nil
}

// SendNotificationEmail sends an email.
func SendNotificationEmail(name, recipientEmail, status string) error {
	htmlBody :=
		`<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Iamza Cornerstone Verifier</title>
	</head>
	<body>
		<p>Hi ` + name + `</p>
		<p>Please note that your Cornerstone Credential has been verified and ` + status + `</p>
		<p>Regards<br/>Iamza Cornerstone Verifier</p>
	</body>
</html>`

	server := mail.NewSMTPClient()

	//SMTP Server
	server.Host = "smtp.office365.com"
	server.Port = 587
	server.Username = "aws_iamzanet@bankservafrica.com"
	server.Password = "K5fEoXa8"
	server.Encryption = mail.EncryptionSTARTTLS
	server.Authentication = mail.AuthLogin

	//Variable to keep alive connection
	server.KeepAlive = true

	//Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	//Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	//SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	//New email simple html with inline and CC
	email := mail.NewMSG()

	email.SetFrom("aws_iamzanet@bankservafrica.com").
		AddTo(recipientEmail).
		// AddCc("otherto@example.com").
		SetSubject("Your Cornerstone Credential is ready")

	email.SetBody(mail.TextHTML, htmlBody)

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}

	return nil
}
