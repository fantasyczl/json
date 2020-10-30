# JSON Command Line Tool

* Usage:
```
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
