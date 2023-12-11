package wendy_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Meduzz/wendy"
)

type (
	private struct {
		Message string `json:"message"`
	}
)

func TestWendyModel(t *testing.T) {
	t.Run("response helpers", func(t *testing.T) {

		t.Run("json body (nil)", func(t *testing.T) {
			subject := wendy.Json(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"application/json","data":null}` {
				t.Errorf("data was incorrect, was: %s", data)
			}
		})

		t.Run("json body (number)", func(t *testing.T) {
			subject := wendy.Json(42)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"application/json","data":42}` {
				t.Errorf("data was incorrect, was: %s", data)
			}
		})

		t.Run("json body (string)", func(t *testing.T) {
			subject := wendy.Json("test")
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"application/json","data":"test"}` {
				t.Errorf("data was incorrect, was: %s", data)
			}
		})

		t.Run("json body (struct)", func(t *testing.T) {
			in := &private{"test"}
			subject := wendy.Json(in)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"application/json","data":{"message":"test"}}` {
				t.Errorf("data was incorrect, was: %s", data)
			}

			if subject.Type != wendy.JSON {
				t.Errorf("type was incorrect, was: %s", subject.Type)
			}

			out := &private{}
			err = subject.AsJson(out)

			if err != nil {
				t.Errorf("there was an error, %v", err)
			}

			if in.Message != out.Message {
				t.Errorf("asJson was incorrect, was: %v, expected: %v", out, in)
			}
		})

		t.Run("text body", func(t *testing.T) {
			text := "test"
			subject := wendy.Text(text)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"text/plain","data":"test"}` {
				t.Errorf("data was incorrect, was: %s", data)
			}

			if subject.AsText() != text {
				t.Errorf("text was incorrect: was: %s", subject.AsText())
			}

			if subject.Type != wendy.TEXT {
				t.Errorf("type was incorrect, was: %s", subject.Type)
			}
		})

		t.Run("static body", func(t *testing.T) {
			text := "test"
			subject := wendy.Static("text/plain", []byte(text))
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"text/plain","data":"74657374"}` {
				t.Errorf("data was incorrect, was: %s", data)
			}

			if subject.Type != "text/plain" {
				t.Errorf("type was incorrect, was: %s", subject.Type)
			}

			if bytes.Equal(subject.AsStatic(), []byte(text)) {
				t.Errorf("asStatic was incorrect, was: %v", subject.AsStatic())
			}
		})

		t.Run("form body", func(t *testing.T) {
			in := make(map[string]interface{})
			in["test"] = "test"
			subject := wendy.Form(in)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if data != `{"type":"application/x-www-form-urlencoded","data":{"test":"test"}}` {
				t.Errorf("data was incorrect, was: %s", data)
			}

			if subject.Type != wendy.FORM {
				t.Errorf("type was incorrect, was: %s", subject.Type)
			}

			out, err := subject.AsForm()

			if err != nil {
				t.Errorf("there was an error: %v", err)
			}

			if !equalsMap(out, in) {
				t.Errorf("asForm was incorrect, was: %v, expected: %v", out, in)
			}
		})

		t.Run("ok (200)", func(t *testing.T) {
			subject := wendy.Ok(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 200 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":200}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("bad request (400)", func(t *testing.T) {
			subject := wendy.BadRequest(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 400 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":400}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("error (500)", func(t *testing.T) {
			subject := wendy.Error(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 500 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":500}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("forbidden (401)", func(t *testing.T) {
			subject := wendy.Forbidden(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 401 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":401}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("invalid (409)", func(t *testing.T) {
			subject := wendy.Invalid(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 409 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":409}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("not allowed (403)", func(t *testing.T) {
			subject := wendy.NotAllowed(nil)
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 403 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":403}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("not found (404)", func(t *testing.T) {
			subject := wendy.NotFound()
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 404 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":404}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})

		t.Run("redirect (303)", func(t *testing.T) {
			subject := wendy.Redirect("/success")
			data, err := serialize(subject)

			if err != nil {
				t.Errorf("serialization went wrong, %v", err)
			}

			if subject.Code != 303 {
				t.Errorf("status code was incorrect: was: %d", subject.Code)
			}

			if data != `{"status":303,"headers":{"Location":"/success"}}` {
				t.Errorf("data was not the expected, was %s", data)
			}
		})
	})
}

func serialize(res any) (string, error) {
	bs, err := json.Marshal(res)
	return string(bs), err
}

func equalsMap(first, second map[string]interface{}) bool {
	for k, v := range first {
		it, ok := second[k]

		if !ok {
			return false
		}

		if v != it {
			return false
		}
	}

	return true
}
