package server

import (
	"context"
	"fmt"
	"example/session"
	"example/auth"
	"example/handlers"
)


type MyServer struct {
	router map[string]handlers.MyHandleFunc
}

func (srv *MyServer) ListenAndServe() {
	for {
		var path, token string
		fmt.Scan(&path)
		fmt.Scan(&token)

		ctx := session.SetSessionID(context.Background())
		go srv.Request(ctx, path, token)
	}
}

func (srv *MyServer) Request(ctx context.Context, path string, token string) {
	var req handlers.MyRequest
	req.SetPath(path)

	ctx = auth.SetAuthToken(ctx, token)

	if handler, ok := srv.router[req.GetPath()]; ok {
		handler(ctx, req)
	} else {
		handlers.NotFoundHandler(ctx, req)
	}
}

