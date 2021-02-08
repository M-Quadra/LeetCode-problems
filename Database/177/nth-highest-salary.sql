CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      SELECT `Salary` FROM (
          SELECT 0 `Salary`
          UNION ALL (
              SELECT DISTINCT(`Salary`)
              FROM `Employee`
              ORDER BY `Salary` DESC
              LIMIT N
          )
      ) a
      LIMIT N, 1
  );
END