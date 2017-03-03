# wait-for-service

A Docker container that exits as soon as a connection with a service can be established.

```shell
$ docker run --rm aduermael/wait-for-service addr:port [retries [delay]]
```

(retries 10 times by default with 1 second delays)

The container exits with code 1 if service can't be reached.