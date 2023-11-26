package util

func IsSupportedStatusID(statusID int) bool {
	switch statusID {
	case 0, 1:
		return true
	}

	return false
}
