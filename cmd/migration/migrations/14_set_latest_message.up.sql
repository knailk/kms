WITH LatestMessages AS (
    SELECT
        cm.chat_session_id,
        cm.id AS latest_message_id,
        cm.created_at AS latest_message_created_at
    FROM
        chat_messages cm
    WHERE
        cm.is_deleted = 0 
        AND cm.created_at = (
            SELECT
                MAX(created_at)
            FROM
                chat_messages
            WHERE
                chat_session_id = cm.chat_session_id
                AND is_deleted = 0
        )
)
UPDATE
    chat_sessions cs
SET
    latest_message_id = lm.latest_message_id
FROM
    LatestMessages lm
WHERE
    cs.id = lm.chat_session_id;
