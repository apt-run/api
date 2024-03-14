package database

const CREATE_SOURCE_TABLE = `
	create table if not exists sources (
		sourceid 	serial,
		name 		text not null unique,
		list		jsonb not null,
		primary 	key(sourceid)
	);
`
const CREATE_PACKAGE_TABLE = `
	create table if not exists packages (
		packageid 		serial,
		name 			text not null unique,
		description 	text not null,
		installs 	    text not null,
		versions 	    text not null,
		url 			text not null unique,
		primary 		key(packageid)
	);
`
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
		primary 		key(packageid)
	);
`
const DROP_TABLE_SOURCES = `
	drop table sources;
`
const DROP_TABLE_PACKAGE = `
	drop table package;
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
const UPSERT_SOURCES = `
	insert into sources (name, list)
	values ($1, $2)
	on conflict(name) 
	do update set
		list = excluded.list;
`

const SEARCH_PACKAGES = `
	select jsonb_build_object('packages', jsonb_agg(value))
	from (
		select value 
		from sources, json_array_elements(sources.list::json -> 'packages') as package(value)
		where package.value->>'name' like ($1 || '%')
		limit $2
	) subquery;
`

const SEARCH_PACKAGES_PAGINATE = `
	select jsonb_build_object('packages', jsonb_agg(value))
	from (
		select value 
		from sources, json_array_elements(list::json -> 'packages') as package(value)
		where package.value->>'name' like ($1 || '%')
		limit $1 offset $1	
	) subquery;
`

const PAGINATE_PACKAGES_START = `
	select jsonb_build_object('packages', jsonb_agg(value))
	from (
		select value
		from sources, json_array_elements(sources.list::json -> 'packages')
		limit 20
	) subquery;
`

const PAGINATE_PACKAGES = `
	select jsonb_build_object('packages', jsonb_agg(value))
	from (
		select value
		from sources, json_array_elements(sources.list::json -> 'packages')
		limit $1 offset $1
	) subquery;
`
