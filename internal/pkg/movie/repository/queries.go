package movrepository

const (
	queryGetMovie = `
	SELECT
		id, poster, title, titleoriginal, rating, info, description, trailer,
		releaseyear, country, genre, motto, director, budget, gross, duration
	FROM movies
	WHERE id = $1;
	`
	
	queryGetMovieCast = `
	SELECT actors.name, actors.id
	FROM movies_actors
	JOIN actors ON movies_actors.actor_id = actors.id
	WHERE movies_actors.movie_id = $1;
	`

	queryGetRelated = `
	SELECT movies.id, movies.poster, movies.title
	FROM movies_movies
	JOIN movies ON movies_movies.relation_id = movies.id
	WHERE movies_movies.movie_id = $1;
	`

	queryGetComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype, ratings.rating
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	LEFT JOIN ratings ON comments.user_id = ratings.user_id
	WHERE movies.id = $1;
	`

	queryGetUserComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype, ratings.rating
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	LEFT JOIN ratings ON comments.user_id = ratings.user_id
	WHERE movies.id = $1 AND users.id = $2;
	`

	queryGetCommentsCount = `
	SELECT COUNT(*)
	FROM comments
	WHERE movie_id = $1 AND user_id = $2;
	`

	queryGetRatingCount = `
	SELECT COUNT(*)
	FROM ratings
	WHERE user_id = $1;
	`

	queryGetUserRating = `
	SELECT rating
	FROM ratings
	WHERE user_id = $1 AND movie_id = $2;
	`

	queryPostRating = `
	INSERT INTO ratings (user_id, movie_id, rating)
	VALUES ($1, $2, $3);
	`

	queryChangeRating = `
	UPDATE ratings
	SET rating = $1
	WHERE user_id = $2;
	`

	queryGetMovieRating = `
	SELECT movies.rating
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	WHERE ratings.movie_id = $1;	
	`

	queryGetMovieVotesnum = `
	SELECT movies.votesnum
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	WHERE ratings.movie_id = $1;	
	`

	queryGetCheckRatingUser = `
	SELECT COUNT(*)
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	JOIN users on ratings.user_id = users.id
	WHERE ratings.user_id = $1 and ratings.movie_id = $2;
	`

	queryGetOldRatingUser = `
	SELECT ratings.rating
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	JOIN users ON ratings.user_id = users.id
	WHERE ratings.user_id = $1 AND ratings.movie_id = $2;
	`

	queryIncrementVotesnum = `
	UPDATE movies
	SET votesnum = votesnum + 1
	WHERE id = $1;
	`

	querySetMovieRating = `
	UPDATE movies
	SET rating = $1
	WHERE id = $2;
	`

	queryUserExist = `
	SELECT COUNT(*)
	FROM users
	WHERE id = $1;
	`
	
	queryMovieExist = `
	SELECT COUNT(*)
	FROM movies
	WHERE id = $1;
	`

	queryPostComment = `
	INSERT INTO comments (user_id, movie_id, commentdate, commenttype, content)
	VALUES ($1, $2, $3, $4, $5);
	`
)
