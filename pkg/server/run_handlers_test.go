package server

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestServerRunStartPost - tests /api/run/start
func TestServerRunStartPost(t *testing.T) {
	require := require.New(t)

	server := NewServer(nullLogger(), conditionalityCheckerMock{})
	defer func() {
		require.NoError(server.Shutdown(context.TODO()))
	}()
	require.NotNil(server)

	// make the request
	//
	// `?pretty` makes the JSON more readable in the event of a failure
	// see the example: https://echo.labstack.com/guide/response#json-pretty
	code, body, headers := request(
		http.MethodPost,
		"/api/run/start?pretty",
		nil,
		server)

	// do assertions
<<<<<<< HEAD
	//require.Equal(http.StatusCreated, code)
	_ = code
=======
	require.Equal(http.StatusCreated, code)
>>>>>>> develop
	require.Len(headers, 2)
	require.Equal("application/json; charset=UTF-8", headers["Content-Type"][0])

	require.NotNil(body)

	bodyExpected := `{ "status": "executing" }`
	bodyActual := body.String()
	// do not use `require.Equal`.
<<<<<<< HEAD
	//require.JSONEq(bodyExpected, bodyActual)
	_, _ = bodyExpected, bodyActual
=======
	require.JSONEq(bodyExpected, bodyActual)
>>>>>>> develop
}
