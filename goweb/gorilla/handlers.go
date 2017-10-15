func wrapHandler(handler func(w http.ResposeWriter, r *http.Request)) func(w http.ResposeWriter, r *http.Request) {
	h := func(w http.ResposeWriter, r *http.Request) {
		if !userIsAuthorized(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
	return h
}

func userIsAuthorized(r *http.Request) bool{
	userID := r.Header.Get("X-HashText-User-ID")
	if userID == ""{
		return false
	}
	var found boolerr := db.QueryRow(`select 1 from "user" where user_id = $1`, userID).Scan(&found)
	switch{
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Printf("query to lookup user failed: %v")
		return false
	}
	return found
}

func userHandler(w http.ResposeWriter, r *http.Request){
	userID := r.Header.Get("X-HashText-User-ID")
	row := db.QueryRow(`select name, credit from "user" where user_id = $1`, userID)

	var name string
	var credit int
	err := row.Scan(&name, &credit)
	switch{
	case err == sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Printf("Query to look up user failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sendJSONReponse(w, userDocument{UserID: userID, Name: name, Credit: credit})
}
