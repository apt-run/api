package database

const CREATE_STATS_TABLE = `
	create table if not exists stats (
		packageid 		serial,
		rank 			integer not null,
		name 			text not null,
		installs 		integer not null,
		vote 			integer not null,
		old 			integer not null,
		recent 			integer not null,
		nofiles 		integer not null,
		maintainer 		text not null,
		url				text,
		versions		jsonb,
		description		text,
		primary 		key(packageid)
	);
`
const DROP_TABLE_STATS = `
	drop table stats;
`
const INSERT_STATS = `
	insert into stats (rank, name, installs, vote, old, recent, nofiles, maintainer)
	values ($1, $2, $3, $4, $5, $6, $7, $8)
`
const SELECT_NAMES = `
	select json_build_object('results', 
			jsonb_agg(json_build_object(
				'name', subquery.name,
				'installs', subquery.installs,
				'maintainer', subquery.maintainer))) as results
	from (
		select 
			to_jsonb(stats.name) as name,
			to_jsonb(stats.installs) as installs,
			to_jsonb(stats.maintainer) as maintainer
		from stats
		where name like ($1 || '%')
		order by name 
		limit $2
	) as subquery;
`
const SELECT_INSTALLS = `
	select json_build_object('results', 
			jsonb_agg(json_build_object(
				'name', subquery.name,
				'installs', subquery.installs,
				'maintainer', subquery.maintainer))) as results
	from (
		select 
			to_jsonb(stats.name) as name,
			to_jsonb(stats.installs) as installs,
			to_jsonb(stats.maintainer) as maintainer
		from stats
		order by installs desc
		limit $1
	) as subquery;
`
const SELECT_MAINTAINERS = `
	select json_build_object('results', 
			jsonb_agg(json_build_object(
				'name', subquery.name,
				'installs', subquery.installs,
				'maintainer', subquery.maintainer))) as results
	from (
		select 
			to_jsonb(stats.name) as name,
			to_jsonb(stats.installs) as installs,
			to_jsonb(stats.maintainer) as maintainer
		from stats
		where maintainers like ($1 || '%')
		order by maintainers 
		limit $2
	) as subquery;
`
