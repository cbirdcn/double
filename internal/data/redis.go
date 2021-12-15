package data

import "strconv"

func getUserKey(id int64) string{
	return "user:" + strconv.FormatInt(id, 10)
}

func getUserNameSetKey() string{
	return "uq_user_name"
}

func getClubNameSetKey() string{
	return "uq_club_name"
}

func getAutoIncrementId() int64{
	// todo
	return 1
}

func getIdByNameHashKey() string{
	return "name_to_id"
}