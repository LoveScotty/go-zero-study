package handler

import (
	"net/http"

	"scotty/go-zero-study/greet/internal/logic"
	"scotty/go-zero-study/greet/internal/svc"
	"scotty/go-zero-study/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
