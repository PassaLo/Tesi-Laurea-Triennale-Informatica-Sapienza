func (rt *_router) updateBluetoothMode(w http.ResponseWriter, r *http.Request, 
ps httprouter.Params, ctx reqcontext.RequestContext) {
	filename, err := strconv.Atoi(ps.ByName("tripid"))
	if err != nil {
		var errorMessage types.ErrorMessage
		errorMessage.ErrorMessage = "tripid not well formed"
		ctx.Logger.WithError(err).Error(errorMessage)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errorMessage)
		return
	}

	newMode, err := io.ReadAll(r.Body)
	if err != nil {
		var errorMessage types.ErrorMessage
		errorMessage.ErrorMessage = "error reading request body"
		ctx.Logger.WithError(err).Error(errorMessage)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errorMessage)
		return
	}

	correctUser, err := rt.db.CheckTripUser(ctx.UserID, filename)
	if err != nil {
		var errorMessage types.ErrorMessage
		errorMessage.ErrorMessage = "Can't check if the user made that trip"
		ctx.Logger.WithError(err).Error(errorMessage)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(errorMessage)
		return
	} else if !correctUser {
		var errorMessage types.ErrorMessage
		errorMessage.ErrorMessage = "This trip was not made by that user"
		ctx.Logger.WithError(err).Error(errorMessage)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(errorMessage)
		return
	}

	message, err := rt.fs.UpdateBluetoothMode(context.TODO(), strconv.Itoa(filename), 
  string(newMode))
	if err != nil {
		var errorMessage types.ErrorMessage

		if message == types.ErrorFileNotFound {
			errorMessage.ErrorMessage = "file not found"
			w.WriteHeader(http.StatusNotFound)
		} else {
			errorMessage.ErrorMessage = "can't update bluetooth connections file"
			w.WriteHeader(http.StatusInternalServerError)
		}
		ctx.Logger.WithError(err).Error(errorMessage)
		w.Header().Set("content-type", "application/json")
		_ = json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusOK)
}
