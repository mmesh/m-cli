module mmesh.dev/m-cli

go 1.17

replace mmesh.dev/m-api-go => ../m-api/_dist/mmesh.dev/m-api-go

replace mmesh.dev/m-lib => ../m-lib

require (
	github.com/AlecAivazis/survey/v2 v2.3.2
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f
	github.com/biter777/countries v1.3.4
	github.com/briandowns/spinner v1.13.0
	github.com/gosuri/uitable v0.0.4
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/pquerna/otp v1.3.0
	github.com/rivo/tview v0.0.0-20211001102648-5508f4b00266
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	google.golang.org/grpc v1.38.0
	mmesh.dev/m-api-go v0.0.0-00010101000000-000000000000
	mmesh.dev/m-lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/creack/pty v1.1.17 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dchest/blake2b v1.0.0 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.4.1-0.20210905002822-f057f0a857a1 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/johnsiilver/getcert v0.0.0-20201005022436-57fe9706afac // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/slack-go/slack v0.10.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023 // indirect
	golang.org/x/oauth2 v0.0.0-20210615190721-d04028783cf1 // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/api v0.22.2 // indirect
	k8s.io/apimachinery v0.22.2 // indirect
	k8s.io/client-go v0.22.2 // indirect
	k8s.io/klog/v2 v2.9.0 // indirect
	k8s.io/utils v0.0.0-20210819203725-bdf08cb9a70a // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
