package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	emailFlag    = flag.String("email", "", "Email address")
	verboseFlag  = flag.Bool("verbose", false, "Verbose mode")
	settingsFlag = flag.String("settings", "", "Settings file")
)

func main() {
	flag.Parse()

	var simple, complex, smtp bool

	if *settingsFlag != "" {
		file, err := os.Open(*settingsFlag)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			switch strings.TrimSpace(line) {
			case "simple=1":
				simple = true
			case "complex=1":
				complex = true
			case "smtp=1":
				smtp = true
			case "verbose=0":
				*verboseFlag = false
			case "verbose=1":
				*verboseFlag = true
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	processEmail := func(email string) {
		if simple {
			fmt.Println("Performing simple function on", email)
		}
		if complex {
			fmt.Println("Performing complex function on", email)
		}
		if smtp {
			fmt.Println("Performing smtp function on", email)
		}
	}

	if *emailFlag != "" {
		processEmail(*emailFlag)
	} else {
		files, err := ioutil.ReadDir("./list")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				path := filepath.Join("./list", file.Name())
				f, err := os.Open(path)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					email := scanner.Text()
					if *verboseFlag {
						simple = false
						complex = false
						smtp = false
						fmt.Print("Enter the function to perform on", email, ": ")
						var choice string
						fmt.Scanln(&choice)
						switch choice {
						case "simple":
							simple = true
						case "complex":
							complex = true
						case "smtp":
							smtp = true
						default:
							fmt.Println("Unknown function")
							continue
						}
					}
					processEmail(email)
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

//package cli
//
//import (
//	"fmt"
//
//	emailverifier "github.com/AfterShip/email-verifier"
//)
//
//var (
//	verifier = emailverifier.NewVerifier()
//)
//
//func main() {
//	email := "example@exampledomain.org"
//
//	ret, err := verifier.Verify(email)
//	if err != nil {
//		fmt.Println("verify email address failed, error is: ", err)
//		return
//	}
//	if !ret.Syntax.Valid {
//		fmt.Println("email address syntax is invalid")
//		return
//	}
//
//	fmt.Println("email validation result", ret)
//	/*
//		result is:
//		{
//			"email":"example@exampledomain.org",
//			"disposable":false,
//			"reachable":"unknown",
//			"role_account":false,
//			"free":false,
//			"syntax":{
//			"username":"example",
//				"domain":"exampledomain.org",
//				"valid":true
//			},
//			"has_mx_records":true,
//			"smtp":null,
//			"gravatar":null
//		}
//	*/
//}

//// package main
//
//// import (
//// 	"fmt"
//// 	"bufio"
////     "flag"
////     "os"
////     "strings"
//// 	emailverifier "github.com/AfterShip/email-verifier"
//// )
//
//// func main() {
////     emailPtr := flag.String("email", "", "Email address")
////     filePtr := flag.String("file", "", "File path")
////     simplePtr := flag.Bool("simple", false, "Simple email validatioin")
////     complexPtr := flag.Bool("complex", false, "Email verification")
////     disposePtr := flag.Bool("dispose", false, "is it disposable")
////     verbosePtr := flag.Bool("verbose", false, "Verbose mode")
////     singlesPtr := flag.Bool("singles", false, "Individual processing")
//
////     flag.Parse()
//
//// 	if *verbosePtr {
////         fmt.Print("Enter the function to perform: ")
////         var choice string
////         fmt.Scanln(&choice)
////         switch choice {
////         case "simple":
////             *simplePtr = true
////         case "complex":
////             *complexPtr = true
////         case "dispose":
////             *disposePtr = true
////         default:
////             fmt.Println("Unknown function")
////             return
////         }
////     }
//
//// 	processEmail := func(email string) {
////         if *simplePtr {
////             fmt.Println("Performing simple function on", email)
////         } else if *complexPtr {
////             fmt.Println("Performing complex function on", email)
////         } else if *disposePtr {
////             fmt.Println("Performing dispose function on", email)
////         }
////     }
//
//// 	if *filePtr != "" {
////         file, err := os.Open(*filePtr)
////         if err != nil {
////             fmt.Println(err)
////             return
////         }
////         defer file.Close()
//
////         scanner := bufio.NewScanner(file)
////         for scanner.Scan() {
////             email := scanner.Text()
////             if *singlesPtr {
////                 *simplePtr = false
////                 *complexPtr = false
////                 *disposePtr = false
////                 fmt.Print("Enter the function to perform on", email, ": ")
////                 var choice string
////                 fmt.Scanln(&choice)
////                 switch choice {
////                 case "simple":
////                     *simplePtr = true
////                 case "complex":
////                     *complexPtr = true
////                 case "dispose":
////                     *disposePtr = true
////                 default:
////                     fmt.Println("Unknown function")
////                     continue
////                 }
////             }
////             processEmail(email)
//
//// 		}
//
////         if err := scanner.Err(); err != nil {
////             fmt.Println(err)
////         }
////     } else if *emailPtr != "" {
////         processEmail(*emailPtr)
////     } else {
////         fmt.Println("No email or file specified")
////     }
//// }
//
//package cli
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//
//	emailverifier "github.com/AfterShip/email-verifier"
//)
//
//type Options struct {
//	Simple   bool
//	Complex  bool
//	Dispose  bool
//	Verbose  bool
//	Singles  bool
//	Email    string
//	FilePath string
//}
//
//func ProcessEmails(options Options) {
//	processEmail := func(email string) {
//		if options.Simple {
//			fmt.Println("Performing simple function on", email)
//		} else if options.Complex {
//			fmt.Println("Performing complex function on", email)
//		} else if options.Dispose {
//			fmt.Println("Performing dispose function on", email)
//		}
//	}
//
//	if options.FilePath != "" {
//		file, err := os.Open(options.FilePath)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		for scanner.Scan() {
//			email := scanner.Text()
//			if options.Singles {
//				options.Simple = false
//				options.Complex = false
//				options.Dispose = false
//				fmt.Print("Enter the function to perform on", email, ": ")
//				var choice string
//				fmt.Scanln(&choice)
//				switch choice {
//				case "simple":
//					options.Simple = true
//				case "complex":
//					options.Complex = true
//				case "dispose":
//					options.Dispose = true
//				default:
//					fmt.Println("Unknown function")
//					continue
//				}
//			}
//			processEmail(email)
//		}
//
//		if err := scanner.Err(); err != nil {
//			fmt.Println(err)
//		}
//	} else if options.Email != "" {
//		processEmail(options.Email)
//	} else {
//		fmt.Println("No email or file specified")
//	}
//}
//
//////
//////var (
//////	verifier = emailverifier.NewVerifier()
//////)
//////
////////func VerifyEmail(email string) {
////////	ret, err := verifier.Verify(email)
////////	if err != nil {
////////		fmt.Println("verify email address failed, error is: ", err)
////////		return
////////	}
////////	if !ret.Syntax.Valid {
////////		fmt.Println("email address syntax is invalid")
////////		return
////////	}
////////
////////	fmt.Println("email validation result", ret)
////////}
//////
//////func main() {
//////	email := ""
//////
//////	ret, err := verifier.Verify(email)
//////	if err != nil {
//////		fmt.Println("verify email address failed, error is: ", err)
//////		return
//////	}
//////	if !ret.Syntax.Valid {
//////		fmt.Println("email address syntax is invalid")
//////		return
//////	}
//////
//////	fmt.Println("email validation result", ret)
//////	/*
//////		result is:
//////		{
//////			"email":"example@exampledomain.org",
//////			"disposable":false,
//////			"reachable":"unknown",
//////			"role_account":false,
//////			"free":false,
//////			"syntax":{
//////			"username":"example",
//////				"domain":"exampledomain.org",
//////				"valid":true
//////			},
//////			"has_mx_records":true,
//////			"smtp":null,
//////			"gravatar":null
//////		}
//////	*/
//////}
////
////var (
////	verifier = emailverifier.
////		NewVerifier().
////		EnableSMTPCheck().DisableCatchAllCheck()
////)
////
////func main() {
////
////	domain := "gmail.com"
////	username := ""
////	ret, err := verifier.CheckSMTP(domain, username)
////	if err != nil {
////		fmt.Println("check smtp failed: ", err)
////		return
////	}
////
////	fmt.Println("smtp validation result: ", ret)
////
////}
////
//
//var (
//	verifier = emailverifier.
//		NewVerifier().
//		EnableAutoUpdateDisposable()
//)
//
//func main() {
//	domain := "gmail.com"
//	if verifier.IsDisposable(domain) {
//		fmt.Printf("%s is a disposable domain\n", domain)
//		return
//	}
//	fmt.Printf("%s is not a disposable domain\n", domain)
//}
//
////
////
////
////
////package xrayd
////
////import (
////"fmt"
////emailverifier "github.com/AfterShip/email-verifier"
////)
////
////var (
////	verifier = emailverifier.NewVerifier()
////)
////
////func VerifyEmail(email string) {
////	ret, err := verifier.Verify(email)
////	if err != nil {
////		fmt.Println("verify email address failed, error is: ", err)
////		return
////	}
////	if !ret.Syntax.Valid {
////		fmt.Println("email address syntax is invalid")
////		return
////	}
////
////	fmt.Println("email validation result", ret)
////}
