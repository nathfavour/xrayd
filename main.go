package main

import (
	"github.com/zserge/lorca"
	"log"
)

func main() {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	quit := make(chan struct{})
	go func() {
		<-ui.Done()
		close(quit)
	}()

	for {
		select {
		case <-quit:
			return
		default:
			// Keep the application running
		}
	}
}

//package main
//
//import (
//	"log"
//
//	"github.com/zserge/lorca"
//)
//
//func main() {
//	ui, err := lorca.New("", "", 480, 320)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ui.Close()
//
//	// Wait until the UI window is closed
//	<-ui.Done()
//}

//// package main
//
//// func main() {
////     VerifyEmail("nathfavour02@proton.me")
//// }
//
//package main
//
//import (
//	"github.com/zserge/lorca"
//	"log"
//	"net/url"
//	"xrayd/cli"
//)
//
//func main() {
//	ui, err := lorca.New("", "", 800, 600)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ui.Close()
//
//	ui.Bind("processEmails", func(email string, simple bool, complex bool, dispose bool) {
//		options := cli.Options{
//			Simple:  simple,
//			Complex: complex,
//			Dispose: dispose,
//			Email:   email,
//		}
//		cli.ProcessEmails(options)
//	})
//
//	html := `
//    <!DOCTYPE html>
//    <html>
//    <head>
//        <title>CLI Wrapper</title>
//    </head>
//    <body>
//        <input id="email" type="text" placeholder="Email">
//        <input id="simple" type="checkbox"> Simple
//        <input id="complex" type="checkbox"> Complex
//        <input id="dispose" type="checkbox"> Dispose
//        <button onclick="process()">Process Email</button>
//        <script>
//            async function process() {
//                let email = document.getElementById('email').value;
//                let simple = document.getElementById('simple').checked;
//                let complex = document.getElementById('complex').checked;
//                let dispose = document.getElementById('dispose').checked;
//                await processEmails(email, simple, complex, dispose);
//            }
//        </script>
//    </body>
//    </html>
//    `
//
//	ui.Load("data:text/html," + url.PathEscape(html))
//
//	<-ui.Done()
//}
