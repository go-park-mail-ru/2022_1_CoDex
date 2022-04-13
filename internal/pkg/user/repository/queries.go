package usrrepository

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc, password
	FROM users
	WHERE email = $1;
	`

	queryGetById = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE id = $1;
	`

	queryAddUser = `
	INSERT INTO
		users (username, email, password)
	VALUES
		($1, $2, $3)
	RETURNING id;
	`

	queryUpdateUser = `
	UPDATE users
	SET username = $1
	WHERE id = $2;
	`

	queryUpdAvatarByUsID = `
	UPDATE users 
	SET picture_url = $2
	WHERE user_id = $1
	`
)
