create database if not exists  game;
use game;
create table if not exists user(
 id int(11) not null auto_increment,
 name varchar(255) not null default '',
 gold int(11) not null default 0,
 exp int(11) not null default 0,
 diamond int(11) not null default 0,
 vip_level  int(11) not null default 0,
 player_level int(11) not null default 1,
 avatar varchar(255) not null default '',
 mobile varchar(20) not null default '',
 passwd varchar(255) not null default '',
 fbtoken varchar(255) not null default '',
 total_betting int(11) not null default 0,
 total_game int(11) not null default 0,
 last_recharge_time
 created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 primary key (id), unique key(`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table if not exists game_settings(
 id int(11) not null auto_increment,
 `key` varchar(255) not null default '',
 description varchar(255) not null default '',
 created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 primary key (id) ,unique key(`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table if not exists vip_increase_table(
 id int(11) not null auto_increment,
 vip_level varchar(255) not null default '',
 description varchar(255) not null default '',
 need_exp      int(11) not null default 999999,
 vip_func   varchar(255) not null default '',
 created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 primary key (id),unique key(`vip_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table if not exists shop_items(
 id int(11) not null auto_increment,
 item_name varchar(255) not null default '',
 description varchar(255) not null default '',
 money      int(11) not null default 999999,
 discount   varchar(255) not null default '',
 item_func  varchar(255) not null default '',
 created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 primary key (id),unique key(`item_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table if not exists mail(
 id int(11) not null auto_increment,
 mail_type int(11) not null default 0,
 mail_param varchar(255) not null default '',
 title varchar(255) not null default '',
 content text not null,
 attach_gold  int(11) not null default 0,
 attach_exp  int(11) not null default 0,
 attach_diamond  int(11) not null default 0,
 attach_item  varchar(255) not null default 0,
 created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table if not exists exp(
id int(11) not null auto_increment,
need_exp int(11) not null default 0,
present varchar(255) not null default '',
created_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;