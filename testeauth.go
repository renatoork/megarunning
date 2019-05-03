package main

/*
import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
)
*/
func mainouth() {
	/*
		// Your credentials should be obtained from the Google
		// Developer Console (https://console.developers.google.com).

			opts, err := oauth2.New(
				oauth2.Client("1044380210135-h5v0ntqle19au35ief58g885tcj0k2rp.apps.googleusercontent.com", "3vKWghltQMsu_j0-nY89MeQo"),
				oauth2.RedirectURL("http://localhost:8080/"),
				oauth2.Scope(
					"https://www.googleapis.com/auth/userinfo.email",
				),
				google.Endpoint(),
			)
			if err != nil {
				fmt.Println("1")
				log.Fatal(err)
			}

		// Redirect user to Google's consent page to ask for permission
		// for the scopes specified above.
		url := opts.AuthCodeURL("state", "offline", "auto")
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		// Handle the exchange code to initiate a transport
		t, err := opts.NewTransportFromCode("4/1OGztZ8nIFUXO8ztNnCrIbjSR0Hw_JhAg2FGfIlsDqc.0rXU__1evQAaJvIeHux6iLZ3dusjlAI")
		if err != nil {
			fmt.Println("2")
			log.Fatal(err)
		}
		fmt.Println(t.Token())
		client := http.Client{Transport: t}
		result, err := client.Get("https://www.googleapis.com/plus/v1/people/me")
		if err != nil {
			fmt.Println("3")
			log.Fatal(err)
		} else {
			io.Copy(os.Stdout, result.Body)
		}
	*/
}
