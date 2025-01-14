## mollie invoices list

Retrieve all invoices on the account. Optionally filter on year or invoice number.

```
mollie invoices list [flags]
```

### Examples

```
mollie invoices list
```

### Options

```
      --from string        offset the result to the resource with the given id
  -h, --help               help for list
      --limit int          limits the number of rows to retrieve (default 250)
      --reference string   ilter for an invoice with a specific invoice number / reference
      --year string        ilter for invoices from a specific year (e.g. 2020)
```

### Options inherited from parent commands

```
  -c, --config string   specifies a custom config file to be used
      --curl            print the curl representation of a request
  -d, --debug           enables debug logging information
  -f, --fields string   select displayable fields to filter the console output, possible values are RESOURCE,ID,REFERENCE,VAT_NUMBER,STATUS,ISSUED_AT,PAID_AT,DUE_AT,NET_AMOUNT,VAT_AMOUNT,GROSS_AMOUNT
      --json            dumpts the json response instead of the column based output
  -m, --mode string     indicates the api target from test/live (default "test")
      --no-headers      Return raw data with no headers
  -t, --token string    the type of token to use for auth (default "MOLLIE_API_TOKEN")
  -v, --verbose         print verbose logging messages (defaults to false)
```

### SEE ALSO

* [mollie invoices](mollie_invoices.md)	 - Operations over Mollie's Invoices API.

###### Auto generated by spf13/cobra on 20-Jun-2022
