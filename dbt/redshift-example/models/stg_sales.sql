WITH events AS (
    SELECT
        *
    FROM
        {{ ref("stg_events") }}
)
SELECT
    sales.salesid,
    sales.sellerid,
    sales.buyerid,
    sales.qtysold,
    sales.pricepaid,
    sales.commission,
    sales.saletime,
    events.*
FROM
    sales
    LEFT JOIN events
    ON sales.eventid = events.eventid
