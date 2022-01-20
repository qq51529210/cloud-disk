package oauth2

import "encoding/json"

var (
	queryRequiredAppId        []byte
	queryRequiredRedirectUri  []byte
	queryRequiredResponseType []byte
	queryRequiredScope        []byte
	queryRequiredState        []byte
)

func init() {
	queryRequiredAppId, _ = json.Marshal(map[string]string{
		"error": "query <app_id> required",
	})
	queryRequiredRedirectUri, _ = json.Marshal(map[string]string{
		"error": "query <redirect_uri> required",
	})
	queryRequiredResponseType, _ = json.Marshal(map[string]string{
		"error": "query <response_type> required",
	})
	queryRequiredScope, _ = json.Marshal(map[string]string{
		"error": "query <scope> required",
	})
	queryRequiredState, _ = json.Marshal(map[string]string{
		"error": "query <state> required",
	})
}
