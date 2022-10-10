## Generate CVE list table for Yocto-based OS

### Usage

1. Add [CVE-CHECK](https://hub.mender.io/t/how-to-run-cve-checks-using-the-yocto-project/1142)
flag to Yocto
2. Rebuild system image
3. Build & run CVE list generator

```
$ go build
$ ./yocto-report path/to/cve-json-generated-by-yocto.json
```
