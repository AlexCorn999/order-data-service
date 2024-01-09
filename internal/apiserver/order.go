package apiserver

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/AlexCorn999/order-data-service/internal/domain"
)

// checkOrder ...
func (s *APIServer) checkOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./web/order_page.html")
	} else if r.Method == "POST" {

		if err := r.ParseForm(); err != nil {
			s.logger.Error("checkOrder", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		orderID := r.FormValue("orderID")

		res, err := s.cacheStorage.GetOrderByID(orderID)
		if err != nil {
			if errors.Is(err, domain.ErrIncorrectOrder) {
				s.logger.Error("checkOrder ", err)
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			} else {
				s.logger.Error("checkOrder ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// Создание шаблона HTML-страницы
		tmpl, err := template.ParseFiles("./web/order_details.html")
		if err != nil {
			s.logger.Error("checkOrder ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Передача данных объекта res в шаблон
		err = tmpl.Execute(w, res)
		if err != nil {
			s.logger.Error("checkOrder ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
