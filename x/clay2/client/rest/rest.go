package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers clay2-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/clay2/post", listPostHandler(cliCtx, "clay2")).Methods("GET")
	r.HandleFunc("/clay2/post", createPostHandler(cliCtx)).Methods("POST") # 1
}
