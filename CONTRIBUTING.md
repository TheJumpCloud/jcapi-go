### Swagger code generation

This repository relies on the following docker file in order to run swagger-codegen inside a docker container:
https://hub.docker.com/r/jimschubert/swagger-codegen-cli/

We're currently using the version 2.2.2 of swagger-codegen.

Note that there is now an official swagger Docker file for the swagger-codegen-cli but it seems to only be supporting the latest version of swagger-codegen (2.3.0, which generates a completely different API interface from 2.2.2).
This docker file can be found here: https://hub.docker.com/r/swaggerapi/swagger-codegen-cli/
We might want to consider using this Docker file once it supports different versions of swagger-codegen.

### Generating the API Client

Copy the API yaml files to the local `/input` directory.

The API v2 yaml file can be found here: `https://github.com/TheJumpCloud/SI/blob/master/routes/webui/api/v2/index.yaml`

To generate the API v2 client, run the command below (assuming your API v2 yaml file is `input/index2.yaml`):  

```
$ docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/index2.yaml -l go -c /config/config_v2.json -o /swagger-api/out/v2
```
This will generate the API v2 client files under `output/v2`

Once you are satisfied with the generated API client, you can replace the existing files under the `v2` folder with your generated files.

There currently seems to be a bug with swagger-codegen where it fails to generate certain enum structs correctly (namely GraphType and GroupType).
Make sure to run the following commands in order to replace these empty structs with just strings:
```
sed -e ':a' -e 'N' -e '$!ba' -e 's/struct {\n}/string/g' v2/group_type.go > tmp && mv tmp v2/group_type.go
sed -e ':a' -e 'N' -e '$!ba' -e 's/struct {\n}/string/g' v2/graph_type.go > tmp && mv tmp v2/graph_type.go
```
