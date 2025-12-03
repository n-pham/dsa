
/* from https://mystery.knightlab.com
SELECT name, sql FROM sqlite_master
select description
from crime_scene_report
where date = 20180115
and city = 'SQL City'
and type = 'murder'
-- Security footage shows that there were 2 witnesses. The first witness lives at the last house on "Northwestern Dr". The second witness, named Annabel, lives somewhere on "Franklin Ave".
select *
from person
where address_street_name = 'Northwestern Dr' -- or name like 'Annabel%'
order by address_number desc
limit 1
-- person_id          name  license_id
-- 16371	Annabel Miller	490173
-- 14887	Morty Schapiro	118009
select * from interview where person_id in (16371, 14887)
I heard a gunshot and then saw a man run out. He had a "Get Fit Now Gym" bag. The membership number on the bag started with "48Z". Only gold members have those bags. The man got into a car with a plate that included "H42W".
I saw the murder happen, and I recognized the killer from my gym when I was working out last week on January the 9th.
select * from drivers_license where plate_number like '%H42W%'
select m.*
from get_fit_now_check_in ci 
inner join get_fit_now_member m on ci.membership_id = m.id
where ci.check_in_date = 20180109
select m.*
from get_fit_now_check_in ci 
inner join get_fit_now_member m on ci.membership_id = m.id
inner join person p on m.person_id = p.id
inner join drivers_license d on p.license_id = d.id
where ci.check_in_date = 20180109
and d.plate_number like '%H42W%'
-- 
select * from interview where person_id in (67318)
I was hired by a woman with a lot of money. I don't know her name but I know she's around 5'5" (65") or 5'7" (67"). She has red hair and she drives a Tesla Model S. I know that she attended the SQL Symphony Concert 3 times in December 2017.
*/
select p.name
from person p
inner join drivers_license d on p.license_id = d.id
inner join facebook_event_checkin e on e.person_id = p.id
where 1=1
-- and d.hair_color = 'red'
-- and d.car_model like '%Tesla%'
and e.event_name = 'SQL Symphony Concert'
and e.date between 20171201 and 20171231
group by p.name
having count(*) = 3
-- Miranda Priestly
