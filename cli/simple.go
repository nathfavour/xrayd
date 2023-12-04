package cli

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier()
)

func main() {
	email := "example@exampledomain.org"

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return
	}

	fmt.Println("email validation result", ret)
	/*
		result is:
		{
			"email":"example@exampledomain.org",
			"disposable":false,
			"reachable":"unknown",
			"role_account":false,
			"free":false,
			"syntax":{
			"username":"example",
				"domain":"exampledomain.org",
				"valid":true
			},
			"has_mx_records":true,
			"smtp":null,
			"gravatar":null
		}
	*/
}
