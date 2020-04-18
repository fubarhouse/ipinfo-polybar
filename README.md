# IPInfo Polybar module

This tool will perform a `GET` request to `ipinfo.io` and output the response as a formatted string.

## Usage

```
ipinfo-polybar -f "%city%, %country% (%ip%)"
```

## Example

```
[module/ipinfo]
type = custom/script
label = %output%
exec = /path/to/ipinfo-polybar -f "%city%, %country% (%ip%)"
interval = 60
```