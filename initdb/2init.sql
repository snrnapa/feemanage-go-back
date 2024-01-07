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

-- teams初期データ
insert into
    teams(team_id, team_name, reader_id)
select
    i,
    format('チームネーム%s', i),
    i
from
    generate_series(1, 1000) as i;

-- depts初期データ
insert into
    depts(dept_id, dept_name, reader_id)
select
    i,
    format('チームネーム%s', i),
    i
from
    generate_series(1, 1000) as i;

-- -- gormで定義する可能性が高いためコメントアウト
-- -- m_dept初期データ
-- insert into
--     m_dept (id, dept_name)
-- Values
--     (1, 'PPX');
-- insert into
--     m_dept (id, dept_name)
-- Values
--     (2, 'CAX');
-- insert into
--     m_dept (id, dept_name)
-- Values
--     (3, 'VSX');
-- insert into
--     m_dept (id, dept_name)
-- Values
--     (4, 'TSX');
-- insert into
--     m_dept (id, dept_name)
-- Values
--     (5, 'RSX');
-- -- m_team初期データ
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (1, 1, 'HUW');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (1, 2, 'PPW');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (2, 1, 'PPB');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (2, 2, 'LPB');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (3, 1, 'ACV');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (3, 2, 'ZCV');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (4, 1, 'XBV');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (4, 2, 'XMN');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (5, 1, 'ZXZ');
-- insert into
--     m_team (dept_id, seq, team_name)
-- Values
--     (5, 2, 'YZM');
-- -- t_fee初期データ
-- insert into
--     t_fee(id, fee_seq, round_trip, fee, use_date)
-- select
--     i,
--     n,
--     round((random() * (1 - 2)) :: numeric, 0) + 2,
--     round((random() * (140 - 2000)) :: numeric, 0) + 2000,
--     CAST(
--         date '2023-11-10' + '1 day' :: INTERVAL * d AS date
--     )
-- from
--     generate_series(1, 1000) as i,
--     generate_series(1, 2) as n,
--     generate_series(1, 200) as d;