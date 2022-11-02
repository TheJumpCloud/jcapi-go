
# Setup "hello-world" project

1. Created directory `hello-world`
2. `cd hello-world`
3. Created file `go.work` to point module to local version of SDK
4. Created file `main.go` and populated with example from main [README.md](https://github.com/TheJumpCloud/jcapi-go/tree/latest)
5. Created "go.mod" file by running `go mod init github.com/TheJumpCloud/jcapi-go/hello-world`
6. Ran `go mod tidy` as prompt suggests
7. Run `go run main.go` to execute program

# SDK fixes

1. `cd jcapiv2`
2. Run `go mod init github.com/TheJumpCloud/jcapi-go/jcapiv2`
3. Run `go mod tidy`
4. Run `cd ../hello-world; go run main.go;` but found errors
   1. Error: `undefined: os` to fix I opened each file, saved it, and let linting do its magic
      1. `../jcapi-go/jcapiv2/api_logos.go`
      2. `../jcapi-go/jcapiv2/api_managed_service_provider.go`
      3. `../jcapi-go/jcapiv2/api_providers.go`
5. Made some fixes in the example `../hello-world/main.go`
6. It worked! `Details for User group 635ac16209b40b00016188f1: {Attributes:0xc000189ef0 Description: Email: Id:635ac16209b40b00016188f1 MemberQuery:<nil> MemberQueryExemptions:[] MemberSuggestionsNotify:false MembershipAutomated:false Name:PesterTest_UserGroup SuggestionCounts:<nil> Type_:user_group}`

# Misc.

The VSCode linter still is reporting a lot of errors that probably need to be fixed.
