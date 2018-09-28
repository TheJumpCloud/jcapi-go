### Swagger code generation

This repository relies on the following Dockerfile in order to run
Swagger Codegen inside a Docker container:
https://hub.docker.com/r/jimschubert/swagger-codegen-cli/.

We're currently using version 2.3.1 of Swagger Codegen.

### Generating the API client

Copy the Swagger specification YAML files to the local `./input` directory.

The API v1 and v2 files can be found using the OAS export link on our
documentation pages: https://docs.jumpcloud.com/1.0 and
https://docs.jumpcloud.com/2.0.

Update the version number for each package in `config_v1.json` or
`config_v2.json`.

To generate the API v1 or v2 client, run the commands below (assuming your
API v1 and v2 specification files are `./input/index1.yaml` and
`./input/index2.yaml`):

```
docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/index1.yaml -l go -c /config/config_v1.json -o /swagger-api/out/v1
docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/index2.yaml -l go -c /config/config_v2.json -o /swagger-api/out/v2
```

This will generate the API v1 and v2 client files under the local
`./output/v1` and `./output/v2` directories.

Once you are satisfied with the generated API client, you can replace the
existing files under the `v1` or `v2` directory with your generated files:

```
rm -rf v1
mv output/v1 .

rm -rf v2
mv output/v2 .
```

There currently is a bug with Swagger Codegen where it generates redundant
enums for certain structures (namely `GraphType` and `GroupType`).
Make sure to run the following commands in order to remove the enums
declarations:

```
sed '/type GraphType string/q' v2/graph_type.go > tmp && mv tmp v2/graph_type.go
sed '/type GroupType string/q' v2/group_type.go > tmp && mv tmp v2/group_type.go
```
