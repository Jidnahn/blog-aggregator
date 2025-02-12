-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
    ) RETURNING *
) SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT f.name AS feed_name, u.name AS user_name 
FROM feed_follows AS fllw
INNER JOIN feeds AS f ON fllw.feed_id = f.id
INNER JOIN users AS u ON fllw.user_id = u.id
WHERE fllw.user_id = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2 RETURNING *;