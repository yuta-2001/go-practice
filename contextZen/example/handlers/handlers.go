package handlers

import (
	"context"
	"fmt"
	"time"
	"errors"
	"example/auth"
)

type MyHandleFunc func(context.Context, MyRequest)

var GetGreeting MyHandleFunc = func(ctx context.Context, req MyRequest) {
	var res MyResponse

	userID, err := auth.VerifyToken(ctx)
	if err != nil {
		req = MyResponse{code: 403, Err: err}
		fmt.Println(res)
		return
	}

	dbReqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)

	rcvChan := db.DefaultDB.Search(dbReqCtx, userID)
	data, ok := <-rcvChan
	cancel()

	if !ok {
		res = MyResponse{Code: 408, Err: errors.New("DB request timeout")}
		fmt.Println(res)
		return
	}

	res = MyResponse{
		Code: 200,
		Body: fmt.Sprintf("From path %s, Hello! your ID is %d\ndata -> %s", req.path, userID, data),
	}

	fmt.Println(res)
}
