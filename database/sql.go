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
