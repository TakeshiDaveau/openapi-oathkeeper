## openapi-oathkeeper generate

Generate Ory Oathkeeper rules from an OpenAPI 3 to file or Std output

```
openapi-oathkeeper generate [flags]
```

### Options

```
      --allowed-audiences stringToString   Allowed Audiences (default [])
      --allowed-issuers stringToString     Allowed Issuers (default [])
  -f, --file string                        OpenAPI File Path
  -h, --help                               help for generate
      --jwks-uris stringToString           JWKS Uris (default [])
  -o, --output string                      Oathkeeper Rules output path
  -p, --prefix string                      OpenAPI Prefix Id
      --server-url stringArray             API Server Urls
  -u, --url string                         OpenAPI URL
```

### SEE ALSO

* [openapi-oathkeeper](openapi-oathkeeper.md)	 - Generate Ory Oathkeeper Rules from OpenAPI 3.0 files

###### Auto generated by spf13/cobra on 12-Apr-2023
