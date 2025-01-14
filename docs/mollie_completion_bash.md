## mollie completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(mollie completion bash)

To load completions for every new session, execute once:

#### Linux:

	mollie completion bash > /etc/bash_completion.d/mollie

#### macOS:

	mollie completion bash > /usr/local/etc/bash_completion.d/mollie

You will need to start a new shell for this setup to take effect.


```
mollie completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -c, --config string   specifies a custom config file to be used
      --curl            print the curl representation of a request
  -d, --debug           enables debug logging information
      --json            dumpts the json response instead of the column based output
  -m, --mode string     indicates the api target from test/live (default "test")
  -t, --token string    the type of token to use for auth (default "MOLLIE_API_TOKEN")
  -v, --verbose         print verbose logging messages (defaults to false)
```

### SEE ALSO

* [mollie completion](mollie_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 20-Jun-2022
