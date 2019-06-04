select distinct A.Email
from Person A, Person B
where A.Email = B.Email and A.Id != b.Id