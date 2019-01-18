
# Go API Tools

## ref_all

The `ref_all` tools creates a go source file that references all elements of an SDK interface. It can
be used to compare versions of the SDK. More specifically, newly-generated SDKs can be compared to the existing
master version to determine which [Semantic Versioning](https://semver.org) component to increment.

In addition, the generated go file can be compiled to validate the correctness of the gernerated SDK.

Here is an example session:

```
MB-Bawn:jcapi-go bobbawn$ git checkout publish-1.35.0
Switched to branch 'publish-1.35.0'
Your branch is up to date with 'origin/publish-1.35.0'.
(reverse-i-search)`go run': go run api_v2_example.go 
MB-Bawn:jcapi-go bobbawn$ go run tools/ref_all.go -imports "fmt,net/http,time" -pkgBase github.com/TheJumpCloud/jcapi-go v1 > ref_v1_1_35_0.go
2019/01/18 08:46:04 Skipping value ContextAPIKey with no Type
2019/01/18 08:46:04 Skipping value ContextAccessToken with no Type
2019/01/18 08:46:04 Skipping value ContextBasicAuth with no Type
2019/01/18 08:46:04 Skipping value ContextOAuth2 with no Type
MB-Bawn:jcapi-go bobbawn$ git branch
  hy-276-ref-tool
  master
* publish-1.35.0
  ref-tool
  test-swagger-codegen-2.4.0
MB-Bawn:jcapi-go bobbawn$ git checkout master
Switched to branch 'master'
Your branch is up to date with 'origin/master'.
MB-Bawn:jcapi-go bobbawn$ go run tools/ref_all.go -imports "fmt,net/http,time" -pkgBase github.com/TheJumpCloud/jcapi-go v1 > ref_v1_master.go
2019/01/18 08:46:34 Skipping value ContextAPIKey with no Type
2019/01/18 08:46:34 Skipping value ContextAccessToken with no Type
2019/01/18 08:46:34 Skipping value ContextBasicAuth with no Type
2019/01/18 08:46:34 Skipping value ContextOAuth2 with no Type
MB-Bawn:jcapi-go bobbawn$ diff <(sort ref_v1_master.go) <(sort ref_v1_1_35_0.go)
65a66,67
> 
> 
82a85
>     // Fields for Body1
105a109
>     // Fields for SystemSshdParams
193a198
>     var _  v1.Body1
216a222
>     var _  v1.SystemSshdParams
285d290
<     var _ []string = new(v1.System).SshdParams
314a320
>     var _ []v1.SystemSshdParams = new(v1.System).SshdParams
334a341
>     var _ bool = new(v1.Body1).Exclusion
336a344
>     var _ bool = new(v1.Fde).Active
408,409d415
<     var _ int32 = new(v1.Commandresult).RequestTime
<     var _ int32 = new(v1.Commandresult).ResponseTime
473a480,481
>     var _ string = new(v1.Commandresult).RequestTime
>     var _ string = new(v1.Commandresult).ResponseTime
518a527
>     var _ string = new(v1.System).AmazonInstanceID
532a542,543
>     var _ string = new(v1.SystemSshdParams).Name
>     var _ string = new(v1.SystemSshdParams).Value
658a670
>     var _ time.Time = new(v1.Body1).ExclusionUntil
```

In general, changed or deleted lines imply a MAJOR component increment, added lines imply a MINOR component
increment, no change implies a PATCH component increment.
