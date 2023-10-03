module dc.local/decryptor/base64OeapSha1

go 1.16

replace dc.local/decryptor/common => ../common

replace dc.local/zlogger => ../../zlogger

require (
	dc.local/decryptor/common v0.0.0-00010101000000-000000000000
	dc.local/zlogger v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)
