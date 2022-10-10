## Generate CVE list table for Yocto-based OS

### Usage

1. Add [CVE-CHECK] flag to Yocto
2. Rebuild system image
3. Build & run CVE list generator

```
$ go build
$ ./yocto-report path/to/cve-json-generated-by-yocto.json
```
