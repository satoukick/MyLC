select D.Name as Department, E.Name as Employee, E.Salary
from Department D, Employee E, (select DepartmentId, max(Salary) as max from Employee group by DepartmentId) T
where E.DepartmentId = D.Id and T.max = E.Salary and T.DepartmentId = E.DepartmentId