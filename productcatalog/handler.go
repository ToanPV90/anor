package productcatalog

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/aliml92/anor/search"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"

	"github.com/aliml92/anor"
	"github.com/aliml92/anor/html"
	"github.com/aliml92/anor/pkg/httperrors"
)

var ErrInvalidSlug = errors.New("invalid slug value")

type BindValidator interface {
	Binder
	Validator
}

type Binder interface {
	Bind(r *http.Request) error
}

type Validator interface {
	Validate() error
}

func bindValid[T BindValidator](r *http.Request, v T) error {
	if err := v.Bind(r); err != nil {
		return fmt.Errorf("bind request: %w", err)
	}
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}

type Handler struct {
	userSvc     anor.UserService
	productSvc  anor.ProductService
	categorySvc anor.CategoryService
	searcher    search.Searcher
	render      *html.Render
	logger      *slog.Logger
	session     *scs.SessionManager
}

func NewHandler(
	userSvc anor.UserService,
	productSvc anor.ProductService,
	categorySvc anor.CategoryService,
	searcher search.Searcher,
	render *html.Render,
	logger *slog.Logger,
	session *scs.SessionManager,
) *Handler {
	return &Handler{
		userSvc:     userSvc,
		productSvc:  productSvc,
		categorySvc: categorySvc,
		searcher:    searcher,
		render:      render,
		logger:      logger,
		session:     session,
	}
}

func (h *Handler) clientError(w http.ResponseWriter, err error, statusCode int) {
	httperrors.ClientError(h.logger, w, err, statusCode)
}

func (h *Handler) serverInternalError(w http.ResponseWriter, err error) {
	httperrors.ServerInternalError(h.logger, w, err)
}

func (h *Handler) logClientError(err error) {
	httperrors.LogClientError(h.logger, err)
}

func calcTotalPages(total int64, perPage int) int {
	a := total / int64(perPage)
	b := total % int64(perPage)
	if b != 0 {
		a++
	}
	return int(a)
}
