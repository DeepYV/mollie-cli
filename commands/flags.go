package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

// AddResourceFlag attaches the --resource flag to the given command.
func AddResourceFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  ResourceArg,
		Usage: "filter for methods that can be used in combination with the provided resource (orders/payments)",
	})
}

// AddLocaleFlag attaches the --locale flag to the given command.
func AddLocaleFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  LocaleArg,
		Usage: "get the payment method name in the corresponding language",
	})
}

// AddSequenceTypeFlag attaches the --sequence-type flag to
// the given command.
func AddSequenceTypeFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  SequenceTypeArg,
		Usage: "filter methods by sequence type (oneoff, first, recurring)",
	})
}

// AddCurrencyFlags attaches the --amount-currency and the
// --amount-value flags, this allows to specify the value and
// currency code to any given struct.
//
// Both flags are required together to compose mollie amount objects.
func AddCurrencyFlags(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  AmountCurrencyArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-value)",
	})
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  AmountValueArg,
		Usage: "get only payment methods that support the amount and currency (linked to amount-currency)",
	})
}

// AddCurrencyCodeFlag attaches the --currency flag to the given command.
func AddCurrencyCodeFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  CurrencyArg,
		Usage: "the currency to receiving the minimumAmount and maximumAmount in",
	})
}

// AddBillingCountryFlag attaches the --billing-country flag
// to the given command.
func AddBillingCountryFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  BillingCountryArg,
		Usage: "filter for methods supporting the ISO-3166 alpha-2 customer billing country",
	})
}

// AddWalletFlag attaches the --wallet flag to the given command.
func AddWalletFlag(cmd *commander.Command) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:  WalletsArg,
		Usage: "a comma-separated list of the wallets you support in your checkout (applepay)",
	})
}

// AddIDFlag attaches the --id flag to the given command.
// It accepts a boolean to indicate if the flag is a required
// field for the command.
func AddIDFlag(cmd *commander.Command, req bool) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:     IDArg,
		Usage:    "the payment method id",
		Required: req,
	})
}

// AddIncludeFlag attaches the --include flag to the given command.
// It accepts a boolean to indicate if this flag should be a persistent
// field for the command.
func AddIncludeFlag(cmd *commander.Command, p bool) {
	commander.AddFlag(cmd, commander.FlagConfig{
		Name:       IncludeArg,
		Shorthand:  "i",
		Usage:      "this resource allows to enrich the request by including other objects",
		Persistent: p,
	})
}

// AddPrompterFlag attaches the --prompt flag to the given command.
// It accepts a boolean to indicate if this flag should be a persistent
// field for the command.
func AddPrompterFlag(cmd *commander.Command, p bool) {
	commander.AddFlag(cmd, commander.FlagConfig{
		FlagType:   commander.BoolFlag,
		Name:       "prompt",
		Shorthand:  "p",
		Usage:      "prompts for values instead of parsing them from flags (not required only)",
		Persistent: p,
		Default:    false,
	})
}
