package util

// WriteHeaders return a map of default headers for response objects
// for API Gateway procxy responses. The defaults can be extended by passing
// a map of strings as the input
func WriteHeaders(ext ...map[string]string) (defaultHeaders map[string]string) {
	defaultHeaders = map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Request-Method":    "GET, POST, DELETE, OPTIONS",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Headers":     "Content-Type",
	}
	if ext != nil {
		for k := range ext[0] {
			defaultHeaders[k] = ext[0][k]
		}
	}
	return defaultHeaders
}
