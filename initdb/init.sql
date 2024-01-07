-- users初期データ
insert into
    users(id, last_name, first_name, dept_id, team_id)
select
    i,
    format('ふぁーすとねーむ%s', i),
    format('らすとねーむ%s', i),
    round((random() * (1 - 5)) :: numeric, 0) + 5,
    round((random() * (1 - 2)) :: numeric, 0) + 2
from
    generate_series(1, 1000) as i;

-- depts初期データ
insert into
    depts(dept_id, dept_name, leader_id)
select
    i,
    format('GZ%s', i),
    i
from
    generate_series(1, 3) as i;

-- teams初期データ
insert into
    teams(team_id, team_name, leader_id)
select
    i,
    format('GF%s', i),
    i
from
    generate_series(1, 10) as i;

-- fees初期データ
insert into
    fees(user_id, seq, date, departure, arraival, price)
select
    i,
    n,
    CAST(
        date '2023-11-10' + '1 day' :: INTERVAL * d AS date
    ),
    round((random() * (1 - 100)) :: numeric, 0) + 100,
    round((random() * (1 - 100)) :: numeric, 0) + 100,
    round((random() * (140 - 2000)) :: numeric, 0) + 2000
from
    generate_series(1, 1000) as i,
    generate_series(1, 2) as n,
    generate_series(1, 200) as d;

-- reports初期データ
insert into
    reports(
        user_id,
        date,
        seq,
        location,
        start_time,
        end_time,
        rest_flg,
        proj_id
    )
select
    i,
    CAST(
        date '2023-11-10' + '1 day' :: INTERVAL * d AS date
    ),
    n,
    round((random() * (1 - 3)) :: numeric, 0) + 3,
    CAST(
        date '2023-11-10' + '1 day' :: INTERVAL * d AS date
    ),
    CAST(
        date '2023-11-10' + '1 day' :: INTERVAL * d AS date
    ),
    round((random() * (0 - 1)) :: numeric, 0) + 1,
    round((random() * (1000 - 1100)) :: numeric, 0) + 1100
from
    generate_series(1, 1000) as i,
    generate_series(1, 8) as n,
    generate_series(1, 200) as d;