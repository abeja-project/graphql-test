package abeja_test

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/abeja-project/graphql-test/internal/abeja"
	"github.com/abeja-project/graphql-test/internal/database"
)

func TestHanlder(t *testing.T) {
	db := database.New()
	srv := httptest.NewServer(abeja.New(db))
	defer srv.Close()

	for _, tt := range []struct {
		name   string
		query  string
		expect string
	}{
		{
			name:   "fetch empty todos",
			query:  `{"query": "{ todos{id,text,done,user{name}} }"}`,
			expect: `{"data":{"todos":[]}}`,
		},

		{
			name:   "create first todo",
			query:  `{"query": "mutation{ createTodo(input:{text:\"Code fixen\", userId: \"1\"}){id,text} }"}`,
			expect: `{"data":{"createTodo":{"id":"1","text":"Code fixen"}}}`,
		},

		{
			name:   "fetch created todo",
			query:  `{"query": "{ todos{id,text,done,user{name}} }"}`,
			expect: `{"data":{"todos":[{"id":"1","text":"Code fixen","done":false,"user":{"name":"Anja"}}]}}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := srv.Client().Post(srv.URL, "", strings.NewReader(tt.query))
			if err != nil {
				t.Fatalf("Got unexpected err: %v", err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Can not read response body: %v", err)
			}

			if resp.StatusCode != 200 {
				t.Fatalf("Got status %s, expected 200\n%s", resp.Status, body)
			}

			if string(body) != tt.expect {
				t.Errorf("Got:\n\n%s\n\nExpected:\n\n%s", body, tt.expect)
			}
		})
	}
}
