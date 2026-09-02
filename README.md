## parser-go-yaml

Parser for Yaml in Go.

The Parser implements an outdated interface and needs to be updated to the new interface, see https://git.katydid.org.za/parser-go

## Known Issues

The YAML parser implementation is in progress and has the following knwown limitations
- Array elements with more than 1 key are parsed incorrectly
- Comments are not supported
- The file delimeter "---" is not supported
- JSON syntax is not yet supported
- Quotes are currently included as part of strings
- Mixing tabs and spaces in indendation is not supported
