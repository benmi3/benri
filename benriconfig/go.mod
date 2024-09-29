module gitlab.com/benmi/benri/benriconfig

go 1.22.5

require (
	gitlab.com/benmi/benri/modules/ddns v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require golang.org/x/net v0.29.0 // indirect

replace gitlab.com/benmi/benri/modules/ddns => ../modules/ddns
