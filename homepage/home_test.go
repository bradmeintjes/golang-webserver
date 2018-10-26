package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandlers(t *testing.T) {
	tests := []struct{
		name string
		in *http.Request
		out *httptest.ResponseRecorder
		expectedStatus int
		expectedBody string
	}{
		{
			name: "good",
			in: httptest.NewRequest("GET", "/home", nil),
			out: httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody: message,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := New(nil)
			h.Home(test.out, test.in)

			status := test.out.Code
			if status != test.expectedStatus {
				t.Logf("expected status: %d\ngot: %d", test.expectedStatus, status)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected body: %s\ngot: %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
