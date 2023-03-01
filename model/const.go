package model

// Request Header keys
const (
	HeaderKeyUserAgent = "User-Agent"
	HeaderKeyUserIP    = "X-Forwarded-For"
	HeaderKeyRequestID = "Request_ID"
)

// Logger additional fields key
const (
	LoggerKeyRequestID = "REQUEST_ID"
	LoggerKeyOperation = "OPERATION"
	LoggerKeyUserIP    = "USER_IP"
	LoggerKeyUserAgent = "USER_AGENT"
	ContextLogger      = "contextLogger"
	ContextHeader      = "contextHeader"
)

const (
	SuccessfulResultMessage = "✅ Təşəkkürlər! Hesabatı əldə etmək üçün e-poçt ünvanınızı yoxlaya bilərsiniz."
	InvalidRequestMessage   = "⚠️ Hazırda hesabatlar pulsuz verildiyi üçün bir istifadəçi maksimum 3 hesabat sorğulaya " +
		"bilər."
	NoResultMessage = "❌Təəssüf ki, sizin axtardığınız VİN koda uyğun avtomobil məlumatları tapılmadı. VIN kodu yeni" +
		"dən nəzərdən keçirə bilərsiniz."
)
