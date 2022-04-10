package util

func GetNameStatus(status int) string {
	switch status {
	case 1:
		return "Waiting for approve"
	case 2:
		return "Approved"
	case 3:
		return "Registration"
	case 4:
		return "Review"
	case 5:
		return "Announcement"
	case 6:
		return "Funding"
	case 7:
		return "Finish"
	case 8:
		return "Reject"
	default:
		return ""
	}
}
