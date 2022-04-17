package query

const (
	GetApprovedScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 2 AND current_date = application_start;`
	UpdateScholarshipStatus = `
	UPDATE scholarship SET status = $1 WHERE id = $2;`
	GetRegistrationScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 3 AND (
        current_date BETWEEN application_start AND application_end
    ) AND current_date = announcement_date - interval '3' day;`
	GetReviewScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 4 AND current_date = announcement_date;`
	GetEmailBlazingScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 6 AND current_date BETWEEN funding_start AND funding_end;`
	GetFundingScholarship = `
	SET TIMEZONE='Asia/Jakarta';
	SELECT id
	FROM scholarship
	WHERE status = 8 AND current_date > funding_end;`
)
