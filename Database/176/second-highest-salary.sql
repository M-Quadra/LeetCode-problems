SELECT 
IFNULL((SELECT Salary
        FROM Employee
        GROUP BY Salary DESC
        LIMIT 1, 1),
       NULL) SecondHighestSalary