package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEndpoint(t *testing.T) {
	testCases := []apiTestCase{
		{
			users:  []*User{{Id: "6a43df", FullName: "Tom Jefferson", Email: "jefferson999@mirro.com"}},
			input:  UpdateUserRequest{Id: "6a43df", Email: pointer("t.jefferson@mirro.com")},
			output: &User{Id: "6a43df", FullName: "Tom Jefferson", Email: "t.jefferson@mirro.com"},
		},
		{
			users: []*User{{Id: "56781a", FullName: "Eric Nilsson", Email: "eric_fantastic@offtop.com"}},
			input: UpdateUserRequest{Id: "56781c", FullName: pointer("Eric Fantastic")},
			err:   UserNotFound,
		},
		{
			users:  []*User{{Id: "556f36", FullName: "Antony Downtown", Email: "antony.downtown@gmail.com"}},
			input:  UpdateUserRequest{Id: "556f37", FullName: pointer("Antony Uptown")},
			output: &User{Id: "556f36", FullName: "Antony Uptown", Email: "antony.downtown@gmail.com"},
		},
		{
			users:  []*User{{Id: "34d35", FullName: "Mickle Now", Email: "m.n@story.com"}},
			input:  UpdateUserRequest{Id: "34d35"},
			output: &User{Id: "34d35", FullName: "Mickle Now", Email: "m.n@story.com"},
		},
		{
			users: []*User{},
			input: UpdateUserRequest{Id: "34d35", FullName: pointer("Nina Mitk"), Email: pointer("m.n@story.com")},
			err:   UserNotFound,
		},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			api := &UserApi{storage: test.users}
			res, err := api.Update(test.input)
			if !cmp.Equal(res, test.output) {
				t.Log("actual result is not as expected", cmp.Diff(res, test.output))
				t.Fail()
			}
			if !cmp.Equal(err, test.err) {
				t.Log("received error is not as expected", cmp.Diff(err, test.err))
				t.Fail()
			}
		})
	}
}
