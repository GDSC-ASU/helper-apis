# GDSC Helper APIs

[![GoDoc](https://godoc.org/github.com/GDSC-ASU/helper-apis?status.png)](https://godoc.org/github.com/GDSC-ASU/helper-apis)
[![Go Report Card](https://goreportcard.com/badge/github.com/GDSC-ASU/helper-apis)](https://goreportcard.com/report/github.com/GDSC-ASU/helper-apis)
[![rex-deployment](https://github.com/GDSC-ASU/helper-apis/actions/workflows/rex-deployment.yml/badge.svg)](https://github.com/GDSC-ASU/helper-apis/actions/workflows/rex-deployment.yml)

Some APIs used for testing purposes.

TODO:

Add more readme

---

## Usage:

**GET https://apis.gdscasu.com/quote**

An endpoint that returns quote and who said it, from a pre defined list.

Sample response:

```json
{
  "text": "Great ideas often receive violent opposition from mediocre minds.",
  "author": "Albert Einstein"
}
```

**GET https://apis.gdscasu.com/image**

An endpoint that returns quote and who said it, from a pre defined list.

Query parameters:

- `count` (optional): specifies the number of returned images' paths.
- `orientation` (default is landscape): specifies the orientation of the returned image(s).

Sample response:

```json
{
  "image": ["/files/images/landscape1.jpg"]
}
```
