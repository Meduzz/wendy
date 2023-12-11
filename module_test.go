package wendy_test

import (
	"testing"

	"github.com/Meduzz/wendy"
)

func TestWendyModule(t *testing.T) {
	text := "test"

	echoRequest := wendy.Request{}
	echoRequest.App = text
	echoRequest.Module = text
	echoRequest.Method = text

	t.Run("handler addition and presence", func(t *testing.T) {
		subject := wendy.NewModule(text, text)

		if subject.App() != text {
			t.Errorf("app was not correct, was: %s", subject.App())
		}

		if subject.Name() != text {
			t.Errorf("name was not correct, was: %s", subject.Name())
		}

		subject.WithHandler(text, echo)

		if !subject.CanHandle(&echoRequest) {
			t.Errorf("could not handle echoRequest")
		}

		subject.WithHandler(text, echo)

		if !subject.CanHandle(&echoRequest) {
			t.Errorf("could not handle echoRequest")
		}

		subject.WithHandler("asdf", echo)

		if !subject.CanHandle(&echoRequest) {
			t.Errorf("could not handle echoRequest")
		}

		appLess := wendy.NewModuleNoApp(text)

		appLess.WithHandler(text, echo)

		if appLess.CanHandle(&echoRequest) {
			t.Errorf("appLess could handle echoRequest")
		}

		echoRequest.App = ""

		if !appLess.CanHandle(&echoRequest) {
			t.Errorf("appLess could not handle echoRequest")
		}

		appLess.SetApp(text)
		echoRequest.App = text

		if !appLess.CanHandle(&echoRequest) {
			t.Errorf("appLess could not handle echoRequest")
		}
	})
}

func echo(req *wendy.Request) *wendy.Response {
	return wendy.Ok(req.Body)
}
