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
    d,
    n,
    round((random() * (1 - 3)) :: numeric, 0) + 3,
    d + cast('9 hour' AS interval) + cast(n || 'hour' AS interval),
    d + cast('10 hour' AS interval) + cast(n || 'hour' AS interval),
    round((random() * (0 - 1)) :: numeric, 0) + 1,
    round((random() * (1000 - 1100)) :: numeric, 0) + 1100
from
    generate_series(1, 400) as i,
    generate_series('2023-10-1', '2024-2-3', '1 day' :: interval) as d,
    generate_series(1, 8) as n;

-- projects初期データ
insert into
    projects(
        proj_id,
        proj_name,
        customer_id,
        leader_id
    )
select
    i,
    format('プロジェクト%s', i),
    round((random() * (1 - 10)) :: numeric, 0) + 10,
    round((random() * (1 - 400)) :: numeric, 0) + 400
from
    generate_series(1000, 1100) as i;