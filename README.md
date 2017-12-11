# example-docker-app
An example docker app that prints its version environment variable and logs to stdout.  Just set the environment variable `VERSION` and run.  

Build the container:

`make build`

Run the container:

`make run`


## Example Dockerfile

```
FROM integrii/example-docker-app
ENV VERSION="My Version 1.0"
EXPOSE 80
ENTRYPOINT /app/app
```
