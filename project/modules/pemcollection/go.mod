module pemcollection

go 1.16

replace dc.local/zlogger => ../zlogger

require (
	dc.local/zlogger v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.29.0 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a
)
