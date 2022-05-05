package annrepository

const (
	queryGetMovies = `
	SELECT
		id, poster, title,
		releasedate, (EXTRACT(epoch FROM (SELECT (releasedate - CURRENT_TIMESTAMP)))/86400)::int::varchar(255), info, description
	FROM announced
	ORDER BY releasedate;
	`

	queryGetAnnounced = `
	SELECT
		id, poster, title, titleoriginal, info, description, trailer,
		releasedate, country, director
	FROM announced
	WHERE id = $1;
	`
	queryGetAnnouncedCast = `
	SELECT actors.name, actors.id
	FROM announced_actors
	JOIN actors ON announced_actors.actor_id = actors.id
	WHERE announced_actors.announced_id = $1
	ORDER BY actors.id;
	`
	queryGetAnnouncedGenres = `
	SELECT genres.genre, genres.title
	FROM genres
	JOIN announced_genres ON genres.genre = announced_genres.genre
	WHERE announced_genres.announced_id = $1
	ORDER BY genres.genre;
	`
	queryGetRelated = `
	SELECT announced.id, announced.poster, announced.title
	FROM announced_announced
	JOIN announced ON announced_announced.relation_id = announced.id
	WHERE announced_announced.announced_id = $1
	ORDER BY announced_announced.relation_id;
	`
)
