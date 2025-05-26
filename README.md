# JSON Command Line Tool

* Installation:

```bash
go install github.com/fantasyczl/json@latest
```

* Usage:

```bash
echo '{"a":1}' | json
{
    "a": 1
}

echo '{"a":1,"map":{"c":2}}' | json
{
    "a": 1
    "map": {
        "c": 2
    }
}

echo '{"a":1,"map":[{"c":2}]}' | json                                                                                                     ✭
{
    "a": 1
    "map": [
        {
            "c": 2
        }
    ]
}
```
