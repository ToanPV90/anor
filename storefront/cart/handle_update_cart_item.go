package cart

import (
	"errors"
	"github.com/aliml92/anor"
	"github.com/aliml92/anor/html/templates/pages/cart/components"
	"github.com/aliml92/anor/session"
	"github.com/invopop/validation"
	"net/http"
	"strconv"
)

type UpdateCartItemRequest struct {
	CartItemID int64
	Qty        int
}

func (req *UpdateCartItemRequest) Bind(r *http.Request) error {
	cartItemIDStr := r.PathValue("id")
	cartItemID, err := strconv.ParseInt(cartItemIDStr, 10, 64)
	if err != nil {
		return err
	}
	req.CartItemID = cartItemID

	qty := r.FormValue("qty")
	req.Qty, err = strconv.Atoi(qty)
	if err != nil {
		return err
	}

	return nil
}

func (req *UpdateCartItemRequest) Validate() error {
	err := validation.Errors{
		"id":  validation.Validate(req.CartItemID, validation.Required, validation.Min(1)),
		"qty": validation.Validate(req.Qty, validation.Required, validation.Min(1)),
	}.Filter()

	return err
}

func (h *Handler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	var req UpdateCartItemRequest
	err := anor.BindValid(r, &req)
	if err != nil {
		if errors.Is(err, errInternal) {
			h.serverInternalError(w, err)
			return
		}
		h.clientError(w, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.cartService.UpdateItem(ctx, req.CartItemID, anor.CartItemUpdateParams{Qty: req.Qty})
	if err != nil {
		h.serverInternalError(w, err)
		return
	}

	u := session.UserFromContext(ctx)
	c, err := h.cartService.Get(ctx, u.CartID, true)
	if err != nil && !errors.Is(err, anor.ErrCartNotFound) {
		h.serverInternalError(w, err)
		return
	}

	w.Header().Add("HX-Trigger-After-Settle", `{"anor:showToast":"item qty updated successfully"}`)
	cs := components.CartSummary{
		HxSwapOOB:      true,
		CartItemsCount: len(c.CartItems),
		TotalAmount:    calculateTotalAmount(c.CartItems),
		Currency:       anor.DefaultCurrency,
	}

	h.Render(w, r, "pages/cart/components/cart_summary.gohtml", cs)
}
