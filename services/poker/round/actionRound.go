package round

import (
	"fmt"
	"time"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

type ActionRoundOptions struct {
	Success            bool
	Payload            *money.Money
	SuccessfulAction *events.Action
	PlayerId string
	PreFlop bool
	CanCheck bool
	CheckCount byte
	Current int
	BlockingIndex int
	BlockingList *BlockingList
}

//RecursiveAction acquires actions from every player so that everybody folds, bets the same amount, or go all in
func (r *Round) RecursiveAction(options *ActionRoundOptions){

	//_____Preceding checks_____

	//Checks in Poker are difficult to keep track. We count all the checks and decide whether more checks are legal.
	log.Logger.Debugf("Checkcount: %v; InCount: %v", options.CheckCount, r.InCount)
	if options.CanCheck && options.CheckCount >= r.InCount {
		log.Logger.Debugf("number of checks exceed lmits")
		options.CanCheck = false
	}

	//Check if only one player is in the game or the blocking list is empty. (One player always wins) (Anchor)
	if r.InCount < 2 || options.BlockingList.CheckIfEmpty() {
		return
	}

	options.Current = options.BlockingList.Get(options.BlockingIndex)
	options.PlayerId = r.Players[options.Current].ID
	options.Success = false

	log.Logger.Debugf("Action round setup")

	// Check for any abnormalities. Player must be active and has money to bet.
	if options.Current < 0 || !r.Players[options.Current].Active || r.Bank.IsAllIn(options.PlayerId) {

		log.Logger.Warnf("Player is not active or already all in. Continuing with next in list")
		// remove from blocking list
		options.BlockingList.RemoveBlocking(options.BlockingIndex)
		options.BlockingIndex = options.BlockingIndex % options.BlockingList.Length()
		r.RecursiveAction(options)
		return
	}

	//____Acquire actions_____

	err := r.ActionTries(options)

	//_____Subsequent checks_____
	_, ok := err.(errors.InvalidActionError)
	if err == nil || ok {
		if err != nil && ok {
			log.Logger.Warnf("Actions were not successful the player is counted as folded")
			options.SuccessfulAction = &events.Action{
				Action:  events.FOLD,
				Payload: moneyUtils.Zero(),
			}
			r.Action(options)
		}
		//Sending results to clients
		utils.SendToAll(r.Players, events.NewActionProcessedEvent(
			options.SuccessfulAction.Action,
			options.Current,
			options.Payload.Display() ,
			r.Bank.GetPlayerBet(options.PlayerId),
			r.Bank.GetPlayerWallet(options.PlayerId),
			r.Bank.GetPot(),
		))
		log.Logger.Debugf("Send results to clients")
		time.Sleep(1 * time.Second)
	}else {
		log.Logger.Infof("Player folded during call")
	}


	if !options.BlockingList.CheckIfEmpty() {
		//Get next in blocking list
		next := options.BlockingList.GetNext(options.SuccessfulAction.Action != events.CHECK, options.BlockingIndex)
		options.BlockingIndex = next

		log.Logger.Debugf("Blocking list is not empty continue with %v", next)

		r.RecursiveAction(options)
		return
	}

	//No one is blocking to continue, so we return
	return
}



func (r *Round) ActionTries(options *ActionRoundOptions) error {
	for i := 3; i > 0 ; i-- {
		if !r.Players[options.Current].Active {
			return errors.PlayerFoldedError{}
		}
		action, err := r.waitForAction(options.Current, options.PreFlop, options.CanCheck)
		if err != nil {
			log.Logger.Warnf("Timeout exeeded or error while waiting for action")
			if !r.Players[options.Current].Active {
				log.Logger.Warnf("Player Inactive")
				return errors.PlayerFoldedError{}
			}
			r.playerError(options.BlockingIndex, fmt.Sprintf("The action was not valid. %v more tries", i))

		}
		if action == nil {
			if !r.Players[options.Current].Active {
				log.Logger.Warnf("Player Inactive")
				return errors.PlayerFoldedError{}
			}
			r.playerError(options.BlockingIndex, fmt.Sprintf("The action was not valid. %v more tries", i))
			continue
		}
		options.SuccessfulAction = action
		options.Payload = action.Payload

		r.Action(options)

		if options.Success {
			break
		}
	}
	log.Logger.Debugf("Action tries concluded")
	if !options.Success {
		return errors.InvalidActionError{}
	}
	return nil
}

func (r *Round) Action(options *ActionRoundOptions) {
	defer log.Logger.Debugf("Action taken successfully: %v", options.Success)
	switch options.SuccessfulAction.Action {
	case events.FOLD:
		err := r.Fold(options.PlayerId)
		if err != nil {
			log.Logger.Errorw("Error during folding", "error", err)
		}
		options.BlockingList.RemoveBlocking(options.BlockingIndex)
		options.Success = true

		return

	case events.CHECK:
		if !options.PreFlop {
			options.CheckCount++
			options.BlockingList.RemoveBlocking(options.BlockingIndex)
			err := options.BlockingList.AddBlocking(options.Current)
			if err == nil {
				options.Success = true
			}
			return
		}

	case events.ALL_IN:
		raise, err := r.Bank.PerformAllIn(options.PlayerId)
		if err == nil {
			options.Success = true
			if raise {
				options.BlockingList.AddAllButThisBlocking(r.Players, options.Current, r.Bank)
			} else {
				options.BlockingList.RemoveBlocking(options.BlockingIndex)
			}
		}
		return

	case events.RAISE:
		err := r.Bank.PerformRaise(options.PlayerId, options.Payload)
		if err == nil {
			options.Success = true
			options.BlockingList.AddAllButThisBlocking(r.Players, options.Current, r.Bank)
			return
		}
		r.playerError(options.BlockingIndex, fmt.Sprintf("Raise must be higher than the highest bet."))
		return

	case events.BET:
		err := r.Bank.PerformBet(options.PlayerId)
		if err == nil {
			options.Success = true
			options.BlockingList.RemoveBlocking(options.BlockingIndex)
			return
		}
		r.playerError(options.BlockingIndex, fmt.Sprintf("Bet must be equal to the current highest bet." ))
	}
}