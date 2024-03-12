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

const INSERT_STATS = `
	insert into stats (rank, name, installs, vote, old, recent, nofiles, maintainer)
	values ($1, $2, $3, $4, $5, $6, $7, $8)
`

const UPSERT_STATS = `
	insert into stats (rank, name, installs, vote, old, recent, nofiles, maintainer)
	values ($1, $2, $3, $4, $5, $6, $7, $8)
	on conflict (name) 
	do update set
		rank = excluded.rank,
		installs = excluded.installs,
		vote = excluded.vote,
		old = excluded.old,
		recent = excluded.recent,
		nofiles = excluded.nofiles,
		maintainer = excluded.maintainer;
`

const UPSERT_SOURCES = `
	insert into sources (name, list)
	values ($1, $2)
	on conflict(name) 
	do update set
		list = excluded.list;
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
