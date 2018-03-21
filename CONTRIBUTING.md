### Swagger code generation

This repository relies on the following docker file in order to run swagger-codegen inside a docker container:
https://hub.docker.com/r/jimschubert/swagger-codegen-cli/

We're currently using the version 2.3.1 of swagger-codegen.

### Generating the API Client

Copy the API yaml files to the local `/input` directory.

The API v2 json/yaml spec can be found using the StopLight export link on our documentation page: `https://docs.jumpcloud.com/2.0`.

To generate the API v2 client, run the command below (assuming your API v2 yaml file is `input/index2.yaml`):  

```
docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/index2.yaml -l go -c /config/config_v2.json -o /swagger-api/out/v2
```
This will generate the API v2 client files under `output/v2`

Once you are satisfied with the generated API client, you can replace the existing files under the `v2` folder with your generated files:
```
rm -rf v2
mv output/v2 .
```

There currently is a bug with swagger-codegen where it generates redundant enums for certain structures (namely GraphType and GroupType).
Make sure to run the following commands in order to remove the enums declarations:
```
sed '/type GraphType string/q' v2/graph_type.go > tmp && mv tmp v2/graph_type.go
sed '/type GroupType string/q' v2/group_type.go > tmp && mv tmp v2/group_type.go
```
