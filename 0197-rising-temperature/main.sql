# Write a SQL query to find all duplicate emails in a table named Person.
# 
# +----+---------+
# | Id | Email   |
# +----+---------+
# | 1  | a@b.com |
# | 2  | c@d.com |
# | 3  | a@b.com |
# +----+---------+
# For example, your query should return the following for the above table:
# 
# +---------+
# | Email   |
# +---------+
# | a@b.com |
# +---------+
# Note: All emails are in lowercase.

# leetcode 196.
# Write your MySQL query statement below

# method 1
SELECT
  w2.Id
FROM
  Weather AS w1, Weather AS w2
WHERE
  DATE_ADD(w1.RecordDate, INTERVAL 1 DAY) = w2.RecordDate
  AND w1.Temperature < w2.Temperature;

# method 2
SELECT
    weather.id AS 'Id'
FROM
    weather
        JOIN
    weather w ON DATEDIFF(weather.date, w.date) = 1
        AND weather.Temperature > w.Temperature
;
