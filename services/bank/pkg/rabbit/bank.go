package rabbit

import (
	"encoding/json"

	ycq "github.com/jetbasrawi/go.cqrs"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type BankCommandHandler struct {
	logger     *zap.SugaredLogger
	bus        ycq.EventBus
	dispatcher ycq.Dispatcher
}

func NewBankCommandHandler(logger *zap.SugaredLogger, bus ycq.EventBus, dispatcher ycq.Dispatcher) *BankCommandHandler {
	return &BankCommandHandler{
		logger:     logger,
		bus:        bus,
		dispatcher: dispatcher,
	}
}

func (h *BankCommandHandler) Handle(d *amqp.Delivery) {
	h.logger.Infof("Received bank")
	cmd, err := readBankCommand(d.Body)
	if err == nil {
		h.logger.Infof("Deserialized: %v", cmd)
		switch cmd.CommandType {
		case bank.Withdraw:
			_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Withdraw{
				Amount:  dtos.FromDTO(cmd.Amount),
				GameId:  cmd.Game,
				RoundId: cmd.Lobby,
				Time:    cmd.Time,
			}))
		case bank.Deposit:
			_ = h.dispatcher.Dispatch(ycq.NewCommandMessage(cmd.UserId, &commands.Deposit{
				Amount:  dtos.FromDTO(cmd.Amount),
				GameId:  cmd.Game,
				RoundId: cmd.Lobby,
				Time:    cmd.Time,
			}))
		}
	}else {
		h.logger.Errorw("Message deserialization error", "error", err)
	}
}

func readBankCommand(raw []byte) (*bank.Command, error) {
	cmd := bank.Command{}
	err := json.Unmarshal(raw, &cmd)
	return &cmd, err
}
