module dc.local.decryptorservice/pkg/server

go 1.16

replace dc.local/zlogger => ./../../modules/zlogger

replace dc.local/decryptor/base64OeapSha1 => ./../../modules/decryptor/base64OeapSha1

replace dc.local.decryptorservice/pkg/cryptkey => ../../pkg/cryptkey

replace dc.local.decryptorservice/pkg/config => ../../pkg/config

require (
	dc.local.decryptorservice/pkg/config v0.0.0-00010101000000-000000000000
	dc.local.decryptorservice/pkg/cryptkey v0.0.0-00010101000000-000000000000
	dc.local/decryptor/base64OeapSha1 v0.0.0-00010101000000-000000000000
	dc.local/zlogger v0.0.0-00010101000000-000000000000
)
