package testdata

const JSONFields = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name": {
            "type": "string"
        },
        "timestamp": {
            "type": "string"
        },
        "identifier": {
            "type": "integer"
        },
        "someThing": {
            "type": "number"
        },
        "complete": {
            "type": "boolean"
        },
        "snakeNumb": {
            "oneOf": [
                {
                    "type": "integer"
                },
                {
                    "type": "string"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}
`
