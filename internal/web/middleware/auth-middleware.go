package middleware

import (
	"net/http"

	"github.com/mauFade/go-payment-gateway/internal/domain"
	"github.com/mauFade/go-payment-gateway/internal/service"
)

type AuthMiddleware struct {
	accService *service.AccountService
}

func NewAuthMiddleware(accService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accService: accService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "api key is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accService.FindByAPIKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
