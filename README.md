gen3pre
=======
Generate aws s3 pre-signed upload url.

```bash
$ go get -u github.com/rluisr/gen3pre
```

Usage
------

### basic
```bash
$ gen3pre --bucket <bucket_name> --file <file_name>
```

### with expires time
```bash
$ gen3pre --bucket <bucket_name> --file <file_name> --exp <expire_time(duration)>
```
default is 12 hours.