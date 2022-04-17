package query

const (
	GetApprovedScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 2 and current_date = application_start;`
	UpdateScholarshipStatus = `
	UPDATE scholarship SET status = $1 WHERE id = $2;`
)
