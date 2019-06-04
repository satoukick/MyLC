create function getNthHighestNumber(N int) returns int
begin
declare M int;
set M = N - 1;
return (
    select distinct Salary from Employee order by Salary limit M,1;
);

end
