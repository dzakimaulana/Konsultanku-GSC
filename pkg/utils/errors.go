package utils

func ErrorOutput(errInput map[string]interface{}) string {
	message := errInput["error"].(map[string]interface{})["message"].(string)
	switch message {
	case "EMAIL_EXISTS":
		message = "The email you entered already in use."
	case "EMAIL_NOT_FOUND":
		message = "Email not found!"
	case "INVALID_PASSWORD":
		message = "Check your email and password!"
	case "INVALID_LOGIN_CREDENTIALS":
		message = "Check your email and password!"
	case "PHONE_NUMBER_EXISTS":
		message = "The phone number is already in use."
	default:
		message = "Something wrong in server"
	}
	return message
}
