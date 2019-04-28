package link

import (
	"os"
	"reflect"
	"testing"
)

func TestLinkParse(t *testing.T) {

	testCases := make(map[string][]Link)

	testCases["ex1.html"] = []Link{
		Link{
			Href: "/other-page",
			Text: "A link to another page",
		},
	}

	testCases["ex2.html"] = []Link{
		Link{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		Link{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}

	testCases["ex3.html"] = []Link{
		Link{
			Href: "#",
			Text: "Login",
		},
		Link{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		Link{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}

	testCases["ex4.html"] = []Link{
		Link{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}

	for file, expected := range testCases {

		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}

		linkList, err := Parse(f)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(linkList, expected) {
			t.Fatalf("expected: %+v, got : %+v", expected, linkList)
		}
	}
}
