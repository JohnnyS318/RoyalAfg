package serviceconfig

import (
	"time"

	"github.com/spf13/viper"
)

const Port = "port"
const PlayersRequiredForStart = "player_required_for_start"
const PlayerActionTimeout = "player_action_timeout"
const BuyInOptions = "buy_in_options"
const RabbitUrl = "rabbit_url"
const RabbitExchange = "rabbit_exchange"
const RabbitBankQueue = "rabbit_bank_queue"
const MatchMakerJWTKey = "matchmaker_signing_key"

const NeedEnterToStart = "gamestart_needs_enter"
const GameRequiresUniquePlayers = "game_require_unique"

const StepSleepDuration = "step_sleep_duration"

const GameStartTimeout = "game_start_timeout"

//SetDefaults configures all defaults
func SetDefaults() {
	viper.SetDefault(Port, 5000)
	//viper.SetDefault(BuyInOptions, []int{500, 1000})
	viper.SetDefault(BuyInOptions, [][]int{{500, 1500, 25}, {1500, 5000, 100}, {5000, 15000, 250}, {15000, 50000, 1000}, {50000, 200000, 2500}})
	viper.SetDefault(PlayersRequiredForStart, 3)
	viper.SetDefault(PlayerActionTimeout, 1*time.Minute)

	viper.SetDefault(RabbitBankQueue, "ryl_bank")
	viper.SetDefault(RabbitExchange, "ryl")
	viper.SetDefault(MatchMakerJWTKey, "matchmakerkey")

	viper.SetDefault(NeedEnterToStart, false)
	viper.SetDefault(GameRequiresUniquePlayers, false)

	viper.SetDefault(StepSleepDuration, time.Second)
	viper.SetDefault(GameStartTimeout, 15)

}
