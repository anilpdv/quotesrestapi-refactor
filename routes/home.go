package routes

import (
	"fmt"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "quotes api")
}
