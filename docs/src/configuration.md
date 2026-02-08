# Configuration

Default path as defined in the source code :

```go{{#include ../../internal/app/config/config.go:default_config_path}}```

## Syntax

`log.level`:
  : Level used for logging.  
    Valid values: `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`

`log.format`:
  : Logging output format.  
    Valid values :
    - `json`: JSON formatted output
    - `console`: Shiny debugging colored output for console
