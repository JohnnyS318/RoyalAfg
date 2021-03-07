package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type TransactionDto struct {
	Amount int `json:"amount"`
}

func (h Account) Deposit(rw http.ResponseWriter, r *http.Request) {

	claims := mw.FromUserTokenContext(r.Context().Value("user"))
	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Logger.Debugf("Depositing %v €", dto.Amount)

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(claims.ID, &commands.Deposit{
		Amount:  money.New(int64(dto.Amount), currency.Code),
		GameId:  "",
		RoundId: "",
		Time:    time.Now(),
	}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h Account) Withdraw(rw http.ResponseWriter, r *http.Request) {

	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(claims.ID, &commands.Withdraw{
		Amount:  money.New(int64(dto.Amount), currency.Code),
		GameId:  "",
		RoundId: "",
		Time:    time.Now(),
	}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}