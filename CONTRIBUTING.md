### Swagger Code Generation

This repository relies on the following Dockerfile in order to run
Swagger Codegen inside a Docker container:
https://hub.docker.com/r/jimschubert/swagger-codegen-cli/.

We're currently using version 2.3.1 of Swagger Codegen.

### Generating the API Client

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
### Versioning

[Semantic Versioning](https://semver.org) is used for the generated client packages.

#### Go Export Parser

`tools/go-export-parser` can be used to validate a new version and determine the semantic version
component to increment for a new release. The tool generates a go source file that references all
exported elements of the API. This source file can be compared to the one for the previous version
to determine which semantic versioning component (major, minor, or patch) to increment.

In addition, the generated go file can be compiled to validate the correctness of the generated SDK to 
some extent.

The tool is a linux executable. (It runs in docker on other OSs).

Here is an example session:

```
$ git checkout -q 1.6.1
$ ./tools/go-export-parser github.com/TheJumpCloud/jcapi-go/v1 > ref_v1-1_6_1.go
$ go run ref_v1-1_6_1.go
Package v1
$ git checkout -q 2.0.0
$ ./tools/go-export-parser github.com/TheJumpCloud/jcapi-go/v1 > ref_v1-2_0_0.go
$ go run ref_v1-2_0_0.go
Package v1
$ diff <(sort ref_v1-1_6_1.go) <(sort ref_v1-2_0_0.go)
65a66,67
>
>
66a69
>     "time"
80a84
>     // Fields for Body1
103a108
>     // Fields for SystemSshdParams
192a198
>     var _  v1.Body1
215a222
>     var _  v1.SystemSshdParams
284d290
<     var _ []string = new(v1.System).SshdParams
313a320
>     var _ []v1.SystemSshdParams = new(v1.System).SshdParams
333a341
>     var _ bool = new(v1.Body1).Exclusion
335a344
>     var _ bool = new(v1.Fde).Active
407,408d415
<     var _ int32 = new(v1.Commandresult).RequestTime
<     var _ int32 = new(v1.Commandresult).ResponseTime
472a480,481
>     var _ string = new(v1.Commandresult).RequestTime
>     var _ string = new(v1.Commandresult).ResponseTime
517a527
>     var _ string = new(v1.System).AmazonInstanceID
531a542,543
>     var _ string = new(v1.SystemSshdParams).Name
>     var _ string = new(v1.SystemSshdParams).Value
657a670
>     var _ time.Time = new(v1.Body1).ExclusionUntil
$ 
```

In general, changed or deleted lines imply a MAJOR component increment, added lines imply a MINOR component
increment, no change implies a PATCH component increment. The above example indicates a new MAJOR version
because the type of `SshdParams`, `RequestTime`, and `ResponseTime` changed.

##### Notes

The go-export-parser tool requires all of jcapi-go's dependent packages to be installed.

If you get failures like:

```
$ tools/go-export-parser github.com/TheJumpCloud/jcapi-go/v1 > ref_v1_2_0_1.go
/go/src/github.com/TheJumpCloud/jcapi-go/v1/api_client.go:20:5: could not import golang.org/x/oauth2 (cannot find package "golang.org/x/oauth2" in any of:
        /usr/local/go/src/golang.org/x/oauth2 (from $GOROOT)
        /go/src/golang.org/x/oauth2 (from $GOPATH))
/go/src/github.com/TheJumpCloud/jcapi-go/v1/api_client.go:21:5: could not import golang.org/x/net/context (cannot find package "golang.org/x/net/context" in any of:
        /usr/local/go/src/golang.org/x/net/context (from $GOROOT)
        /go/src/golang.org/x/net/context (from $GOPATH))
2019/01/24 20:57:21 Couldn't load package path %!d(string=github.com/TheJumpCloud/jcapi-go/v1): err: couldn't load packages due to errors: github.com/TheJumpCloud/jcapi-go/v1
```

Install the dependencies with `go get`:
```
$ go get golang.org/x/oauth2
$ tools/go-export-parser github.com/TheJumpCloud/jcapi-go/v1 > ref_v1_2_0_1.go
$
```
