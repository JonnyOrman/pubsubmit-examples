# pubsubmit examples

Examples demonstrating how to use [pubsubmit](https://github.com/JonnyOrman/pubsubmit).

It contains the following examples:
- `example1` - publishes any data submitted
- `example2` - uses a struct to define the data to be published
- `example3` - config retrieved from the environment
- `example4` - config retrieved from both the config file and the environment
- `example5` - uses the Docker image

# How to run

With Docker installed, clone the repo and run
```
docker-compose up --build
```

To use each of the examples, `POST` to the following:
- `example1` - `localhost:3001`
- `example2` - `localhost:3002`
- `example3` - `localhost:3003`
- `example4` - `localhost:3004`
- `example5` - `localhost:3005`

Include the `Content-Type` header set to `application/json` and a body that contains the data to be published, such as:
```
{
    "prop1": "abc",
    "prop2": 123,
}
```