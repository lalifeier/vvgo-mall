module github.com/lalifeier/vvgo-mall

go 1.19

require (
	entgo.io/contrib v0.3.5
	entgo.io/ent v0.11.10
	github.com/Shopify/sarama v1.38.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.263
	github.com/aliyun/aliyun-oss-go-sdk v2.2.7+incompatible
	github.com/aws/aws-sdk-go-v2 v1.17.7
	github.com/aws/aws-sdk-go-v2/config v1.18.19
	github.com/aws/aws-sdk-go-v2/credentials v1.13.18
	github.com/aws/aws-sdk-go-v2/service/s3 v1.31.0
	github.com/casbin/casbin/v2 v2.66.0
	github.com/casbin/ent-adapter v0.3.0
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/gin-gonic/gin v1.9.0
	github.com/go-chassis/sc-client v0.7.0
	github.com/go-kratos/consul v0.1.5
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/config/kubernetes/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/config/polaris/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/log/aliyun/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/polaris/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/servicecomb/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20230326145430-f03f5f89881f
	github.com/go-kratos/kratos/v2 v2.6.1
	github.com/go-kratos/swagger-api v1.0.1
	github.com/go-oauth2/oauth2/v4 v4.5.2
	github.com/go-playground/assert/v2 v2.2.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-session/session v3.1.2+incompatible
	github.com/go-sql-driver/mysql v1.7.0
	github.com/go-zookeeper/zk v1.0.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/gnostic v0.6.9
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/hashicorp/consul/api v1.20.0
	github.com/jinzhu/copier v0.3.5
	github.com/magiconair/properties v1.8.7
	github.com/minio/minio-go/v7 v7.0.52
	github.com/mojocn/base64Captcha v1.3.5
	github.com/nacos-group/nacos-sdk-go v1.1.4
	github.com/pkg/errors v0.9.1
	github.com/polarismesh/polaris-go v1.3.0
	github.com/prometheus/client_golang v1.14.0
	github.com/qiniu/go-sdk/v7 v7.14.0
	github.com/sirupsen/logrus v1.9.0
	github.com/sony/sonyflake v1.1.0
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.2
	github.com/tencentyun/cos-go-sdk-v5 v0.7.41
	github.com/tx7do/kratos-transport v1.0.5
	github.com/tx7do/kratos-transport/transport/kafka v1.0.2
	go.etcd.io/etcd/client/v3 v3.5.7
	go.mongodb.org/mongo-driver v1.11.3
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/exporters/zipkin v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.7.0
	google.golang.org/genproto v0.0.0-20230327215041-6ac7f18bb9d5
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	k8s.io/client-go v0.26.3
)

require (
	github.com/aliyun/aliyun-log-go-sdk v0.1.43 // indirect
	github.com/apolloconfig/agollo/v4 v4.3.0 // indirect
	github.com/bufbuild/protocompile v0.5.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/sonic v1.8.6 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.8.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.10.2 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/go-chassis/cari v0.9.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jhump/protoreflect v1.15.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.15.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.4 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/tx7do/kratos-transport/broker/kafka v1.0.2 // indirect
	github.com/xdg/scram v1.0.5 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/term v0.6.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.26.3 // indirect
	k8s.io/apimachinery v0.26.3 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230327201221-f5883ff37f0c // indirect
	k8s.io/utils v0.0.0-20230313181309-38a27ef9d749 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require (
	ariga.io/atlas v0.10.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.14.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.7 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230111030713-bf00bc1b83b6 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-kratos/aegis v0.1.4 // indirect
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.12.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/glog v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.16.2 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/lib/pq v1.10.7
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/mozillazg/go-httpheader v0.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/shirou/gopsutil/v3 v3.23.2 // indirect
	github.com/tidwall/btree v1.6.0 // indirect
	github.com/tidwall/buntdb v1.2.10 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	github.com/zclconf/go-cty v1.13.1 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/automaxprocs v1.5.2
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/image v0.6.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
