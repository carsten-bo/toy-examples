{{ config(
    materialized = 'view'
) }}

WITH sales AS (

    SELECT
        *
    FROM
        {{ ref('stg_sales') }}
),
users AS (
    SELECT
        *
    FROM
        {{ ref('stg_users') }}
),
sales_events AS (
    SELECT
        sales.*,
        users.firstname AS seller_firstname,
        users.lastname AS seller_lastname
    FROM
        sales
        LEFT JOIN users
        ON users.userId = sales.sellerId
)
SELECT
    *
FROM
    sales_events
LIMIT
    10
