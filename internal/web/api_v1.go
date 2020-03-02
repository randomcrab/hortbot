package web

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/volatiletech/sqlboiler/queries"
	"go.uber.org/zap"
)

func (a *App) routeAPIv1(r chi.Router) {
	r.Get("/vars/get/{varName}/{channel}", a.apiV1VarsGet)
}

func (a *App) apiV1VarsGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	varName := chi.URLParam(r, "varName")
	channelName := strings.ToLower(chi.URLParam(r, "channel"))

	variable := &models.Variable{}

	err := queries.Raw("SELECT variables.* FROM variables JOIN channels ON variables.channel_id = channels.id WHERE variables.name = $1 AND channels.name = $2", varName, channelName).Bind(ctx, a.DB, variable)
	switch err {
	case nil:
		// Do nothing.
	case sql.ErrNoRows:
		a.httpError(w, r, http.StatusNotFound)
		return
	default:
		ctxlog.Error(ctx, "error querying for variable", zap.Error(err))
		a.httpError(w, r, http.StatusInternalServerError)
		return
	}

	v := struct {
		Channel      string    `json:"channel"`
		Var          string    `json:"var"`
		Value        string    `json:"value"`
		LastModified time.Time `json:"lastModified"`
	}{
		Channel:      channelName,
		Var:          variable.Name,
		Value:        variable.Value,
		LastModified: variable.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&v); err != nil {
		ctxlog.Error(ctx, "failed to write response", zap.Error(err))
		a.httpError(w, r, http.StatusInternalServerError)
		return
	}
}
