# gourl
gourl, looks like curl

## Features
1. output json-formatted string
2. simplify request args
3. similar arguments as `curl`

## How to use
gourl [COMMAND]
COMMAND list:
        --method [post|get], -m:                request method, eg: POST, GET
        --body [string], -b, -d:                post request body
        --header [string], -h, -H               request header
        --json, -j              equals to -H 'Content-Type: application/json'
        --uri, optional         request uri

Examples:

```
gourl http://example.com -d '{"content": "hhh"}' -j
```

output

```json
{
  "attitudeStatus": {
    "upvotes": 1152,
    "downvotes": 171,
    "myvote": "NOT_VOTED"
  }
}
```



