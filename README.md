# convert

This is a tool for translation and file format converting.

For file format, we currently support `json` and `yaml`.

For translation client, we currently use Aliyun machine translation.

## Installation

```
go get -u github.com/hyperjiang/convert
```

## File converter

```
convert file -i "input-file" -o "output-file" -s json -t yaml
```

## Translator

For translation, we only support simple single-layer key-value format, e.g.

```
{
    "key1": "value1",
    "key2": "value2",
}
```

or

```
key1: value1
key2: value2
```

### Use Aliyun machine translation

Setup environment variables:

```
export ALI_REGION_ID="cn-hangzhou"
export ALI_ACCESS_KEY_ID="your-key"
export ALI_ACCESS_SECRET="your-secret"
```

```
convert aliyun -i "input-file" -o "output-file" -s source-language -t target-language
```
