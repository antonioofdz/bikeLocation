package handlers

import (
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/antonioofdz/personalProjectDra/app/pkg/database"
)

func CheckToken(next http.Handler) http.Handler {
	checkValidGuidFn := func(uuid string) bool {
		r := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}")
		return r.MatchString(uuid)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		existTokenBD, err := checkTokenBD(token)
		if err != nil {
			http.Error(w, "INVALID TOKEN", http.StatusUnauthorized)
			return
		}

		if !checkValidGuidFn(token) || !existTokenBD {
			http.Error(w, "INVALID TOKEN", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkTokenBD(token string) (bool, error) {
	sqlCheckTokenBD := "SELECT count(*) from users WHERE token=?"
	db, err := database.Open()
	if err != nil {
		return false, err
	}
	defer db.Close()

	count := 0
	err = db.QueryRow(sqlCheckTokenBD, token).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
