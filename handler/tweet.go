package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/model"
	"github.com/anujc4/tweeter_api/request"
	"github.com/anujc4/tweeter_api/response"
	// "github.com/gorilla/schema"
	// "github.com/gorilla/mux"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
// var decoder = schema.NewDecoder()

func (env *HttpApp) CreateTweet(w http.ResponseWriter, req *http.Request) {
	var request request.CreateTweetRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	if err := request.ValidateCreateTweetRequest(); err != nil {
		app.RenderErrorJSON(w, app.NewError(err))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweet, err := appModel.CreateTweet(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	app.RenderJSONwithStatus(w, http.StatusCreated, response.TransformTweetResponse(*tweet))
    // app.RenderJSON(w, "Not yet implemented!")
}

func (env *HttpApp) GetTweets(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		app.RenderErrorJSON(w, app.NewParseFormError(err))
		return
	}

	var request request.GetTweetsRequest
	if err := decoder.Decode(&request, req.Form); err != nil {
		app.RenderErrorJSON(w, app.NewError(err).SetCode(http.StatusBadRequest))
		return
	}

	appModel := model.NewAppModel(req.Context(), env.DB)
	tweets, err := appModel.GetTweets(&request)
	if err != nil {
		app.RenderErrorJSON(w, err)
		return
	}
	resp := response.MapTweetsResponse(*tweets, response.TransformTweetResponse)
	app.RenderJSON(w, resp)
}

func (env *HttpApp) GetTweetByID(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	//mux docs
	// vars := mux.Vars(req)
	// userID := vars["user_id"]
	// appModel := model.NewAppModel(req.Context(), env.DB)
	// users, err := appModel.GetUserByID(userID)
	// if err != nil {
	// 	app.RenderErrorJSON(w, err)
	// 	return
	// }
	// resp := response.MapUsersResponse(*users, response.TransformUserResponse)
	// app.RenderJSON(w, resp)
    app.RenderJSON(w, "Not yet implemented!")
}

func (env *HttpApp) UpdateTweet(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	app.RenderJSON(w, "Not yet implemented!")
}

func (env *HttpApp) DeleteTweet(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement this
	// vars := mux.Vars(req)
	// userID := vars["user_id"]

	// appModel := model.NewAppModel(req.Context(), env.DB)
	// err := appModel.DeleteUser(userID)
	// if err != nil {
	// 	app.RenderErrorJSON(w, err)
	// 	return
	// }
	// app.RenderJSONwithStatus(w, 1175, "Deleted !")
    app.RenderJSON(w, "Not yet implemented!")
}

