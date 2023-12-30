# GDSC Helper APIs

<div align="center">
[![GoDoc](https://godoc.org/github.com/GDSC-ASU/helper-apis?status.png)](https://godoc.org/github.com/GDSC-ASU/helper-apis)
[![Go Report Card](https://goreportcard.com/badge/github.com/GDSC-ASU/helper-apis)](https://goreportcard.com/report/github.com/GDSC-ASU/helper-apis)
[![rex-deployment](https://github.com/GDSC-ASU/helper-apis/actions/workflows/rex-deployment.yml/badge.svg)](https://github.com/GDSC-ASU/helper-apis/actions/workflows/rex-deployment.yml)
</div>

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
	text:   "Great ideas often receive violent opposition from mediocre minds.",
	author: "Albert Einstein",
},
```
