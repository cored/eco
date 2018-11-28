package server_test

import (
	"context"
	"testing"

	pb "github.com/cored/eco/protos"
	"github.com/cored/eco/server"
	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	s := server.Server{}
	testcases := []struct {
		description string
		text        string
		expected    string
		err         error
	}{
		{
			description: "When an empty string",
			text:        "",
			expected:    "",
			err:         nil,
		},
		{
			description: "When any other string",
			text:        "Hello World!",
			expected:    "Hello World!",
			err:         nil,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			text := &pb.Text{Text: testcase.text}
			resp, err := s.Echo(context.Background(), text)

			assert.Equal(t, testcase.err, err)
			assert.Equal(t, testcase.expected, resp.Text)
		})
	}
}
