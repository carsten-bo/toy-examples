SELECT
    event.eventid,
    event.eventname,
    category.catname,
    category.catgroup,
    category.catdesc
FROM
    event
    LEFT JOIN category
    ON category.catid = event.catid
