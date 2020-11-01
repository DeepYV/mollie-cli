## mollie chargebacks get

Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.

### Synopsis

Retrieve a single chargeback by its ID. Note the original payment’s ID is needed as well.
A debit to a depositor's account for an item that has been previously credited, as for a returned bad check.

```
mollie chargebacks get [flags]
```

### Options

```
      --embed string     a comma separated list of embeded resources
  -h, --help             help for get
      --id string        the chargeback id
      --payment string   original payment id/token
```

### Options inherited from parent commands

```
  -c, --config string   specifies a custom config file to be used
  -d, --debug           enables debug logging information
  -m, --mode string     indicates the api target from test/live (default "test")
      --print-json      toggle the output type to json
  -t, --token string    the type of token to use for auth (default "MOLLIE_API_TOKEN")
  -v, --verbose         print verbose logging messages (defaults to false)
```

### SEE ALSO

* [mollie chargebacks](mollie_chargebacks.md)	 - Operations with the Chargebacks API

###### Auto generated by spf13/cobra on 1-Nov-2020