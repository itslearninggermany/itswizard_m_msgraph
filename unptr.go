package itswizard_m_msgraph

func UnPtrString(input *string) string {
	if input == nil {
		return ""
	} else {
		return *input
	}
}
