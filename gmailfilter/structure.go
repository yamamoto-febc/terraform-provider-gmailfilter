package gmailfilter

const (
	gmailUser = "me"
)

func expandStringSlice(v interface{}) []string {
	if v == nil {
		return nil
	}
	l := v.([]interface{})
	transformed := make([]string, 0, len(l))
	for _, raw := range l {
		transformed = append(transformed, raw.(string))
	}
	return transformed
}
