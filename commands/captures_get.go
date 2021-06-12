package commands

import (
	"github.com/VictorAvelar/mollie-cli/commands/displayers"
	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
	"github.com/spf13/cobra"
)

func getCapturesCmd(p *commander.Command) *commander.Command {
	gc := commander.Builder(
		p,
		commander.Config{
			Namespace: "get",
			ShortDesc: "Retrieve a single capture by its ID.",
			LongDesc: `Retrieve a single capture by its ID. Note the original payment’s ID is needed as well.
Captures are used for payments that have the authorize-then-capture flow. 
The only payment methods at the moment that have this flow are Klarna Pay 
later and Klarna Slice it.`,
			Execute: getCapturesAction,
			Example: "mollie captures get --id ct_example --payment tr_example",
		},
		getCapturesCols(),
	)

	commander.AddFlag(gc, commander.FlagConfig{
		Name:     PaymentArg,
		Usage:    "the payment id/token",
		Required: true,
	})

	commander.AddFlag(gc, commander.FlagConfig{
		Name:     IDArg,
		Usage:    "the capture id/token",
		Required: true,
	})

	return gc
}

func getCapturesAction(cmd *cobra.Command, args []string) {
	payment := ParseStringFromFlags(cmd, PaymentArg)
	id := ParseStringFromFlags(cmd, IDArg)

	if verbose {
		PrintNonemptyFlagValue(PaymentArg, payment)
		PrintNonemptyFlagValue(IDArg, id)
	}

	capture, err := API.Captures.Get(payment, id)
	if err != nil {
		logger.Fatal(err)
	}

	disp := &displayers.MollieCapture{
		Capture: capture,
	}

	err = printer.Display(
		disp,
		display.FilterColumns(parseFieldsFromFlag(cmd), getCapturesCols()),
	)

	if err != nil {
		logger.Fatal(err)
	}
}
