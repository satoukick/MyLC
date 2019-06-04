-- select Score, (select count(distinct Score) from Scores where Score >= s.Score) Rank
-- from Score s
-- order by Score desc


select a.Score, (select count(b.Score) + 1 from (select Score from Scores group by Score) b where a.Score > b.Score) as Rank
from Score as a
order by a.Score desc














select a.Score, (select count(b.Score) + 1 from (select Score from Scores group by Score) b where a.Score > b.Score) Rank
from Score a
order by a.Score desc