# Overview

This document is a infrastructure description and deploy steps.

# Feature

1. this service is forced on how to build a high performance api service.
2. as we know, database servers are usually performance bottleneck, and may cause system delay.
3. I would like my service do as much work as possible instead of relying on database, do my best to reduce db press.
4. package router is a database router. it route every id exactly to one postgreSQL addr, database instance and table.
5. as I said, I would like do mywork instead of database. I use a idcontroller service to generate global user id.
6. based userid, service could route to lots of db. 
7. this id strategy has no limits theoretically. unlike db cluster or db master-slave, they could case db internal consumptions when datas are too match
8. as a result, I need to consumption like/unlike/match state as a resource. every resource' id is userid.
9. todo, there must be a event log protection model which collect the error and data of scene of the accident

# Tables 

User:

    CREATE TABLE t_user_01(
        id BIGINT,
        name varchar(50),
        PRIMARY KEY(id)
    );

Like:

    CREATE TABLE t_relationship_like_01 (
        from_id BIGINT,
        to_id BIGINT,
        PRIMARY KEY(from_id, to_id)
    );

Dislike:

    CREATE TABLE t_relationship_dislike_01 (
        from_id BIGINT,
        to_id BIGINT,
        PRIMARY KEY(from_id, to_id)
    );

Match:

    CREATE TABLE t_relationship_match_01 (
        from_id BIGINT,
        to_id BIGINT,
        PRIMARY KEY(from_id, to_id)
    );
    

# Example 

    databases are a array describe all database config, and database instances, and every instances tables.
    instances are a field of database, it describe this server' instances, and every instance tables.
    
    Now, I have one server.
    I will use this server to storage user/relationship_like/relationship_dislike/relationship_match resources.
    I create 2 instance for every instance which instance name like db_xxx_01 db_xxx_02
    I create 2 table for every instance which table name like t_xxx_01 t_xxx_02
    for now, in the same resource config, every instance's table should exactly same. 
    if I use like 
                    "instances": [
                                {
                                        "name": "db_user_01",
                                        "tables":["t_user_01","t_user_02"]
                                },
                                {
                                        "name": "db_user_02",
                                        "tables":["t_user_03","t_user_04"]
                                }
                      ]
    the router will not work.
   
    {
        "databases": [
                {
                    "label": "user",
                    "host": "10.1.235.98",
                    "port": 5432,
                    "credentials": {
                                    "user": "tantan007",
                                    "password": "123456"
                        },
                    "maxPoolSize": 200,
                    "instances": [
                            {
                                    "name": "db_user_01",
                                    "tables":["t_user_01","t_user_02"]
                            },
                            {
                                    "name": "db_user_02",
                                    "tables":["t_user_01","t_user_02"]
                            }
                            ]
                 },
                {
                        "label": "relationship_like",
                        "host": "10.1.235.98",
                        "port": 5432,
                        "credentials": {
                                        "user": "tantan007",
                                        "password": "123456"
                        },
                        "maxPoolSize": 300,
                        "instances": [
                                    {
                                        "name": "db_relationship_like_01",
                                        "tables":["t_relationship_like_01","t_relationship_like_02"]
                                    }
                            ]
                },
                {
                        "label": "relationship_dislike",
                        "host": "10.1.235.98",
                        "port": 5432,
                        "credentials": {
                                "user": "tantan007",
                                "password": "123456"
                        },
                        "maxPoolSize": 300,
                        "instances": [
                                {
                                        "name": "db_relationship_dislike_01",
                                        "tables":["t_relationship_dislike_01","t_relationship_dislike_02"]
                                }
                            ]
                },
                {
                    "label": "relationship_match",
                    "host": "10.1.235.98",
                    "port": 5432,
                    "credentials": {
                                "user": "tantan007",
                                "password": "123456"
                    },
                    "maxPoolSize": 300,
                    "instances": [
                            {
                                    "name": "db_relationship_match_01",
                                    "tables":["t_relationship_match_01","t_relationship_match_02"]
                            }
                    ]
                }
        ]
    }   

# Installation

# database 
1. create postgreSQL user if you don't have 
    
        adduser tantan007
        su - postgres
        psql
        createuser tantan007
        \password postgres   
        123456
        CREATE USER tantan007 WITH PASSWORD '123456';
    
2. create databases

        CREATE DATABASE db_user_01 OWNER tantan007;
        CREATE DATABASE db_relationship_like_01 OWNER tantan007;
        CREATE DATABASE db_relationship_dislike_01 OWNER tantan007;
        CREATE DATABASE db_relationship_match_01 OWNER tantan007;

3. create tables 

   3.1 create table t_user_01 in instance db_user_01   
   
        CREATE TABLE t_user_01(
            id BIGINT,
            name varchar(50),
            PRIMARY KEY(id)
        );
   
  3.2 create table t_relationship_like_01 in instance db_relationship_like_01
  
        CREATE TABLE t_relationship_like_01 (
            from_id BIGINT,
            to_id BIGINT,
            PRIMARY KEY(from_id, to_id)
        );
  

  3.3 create table t_relationship_dislike_01 in instance db_relationship_dislike_01
  
        CREATE TABLE t_relationship_dislike_01 (
           from_id BIGINT,
           to_id BIGINT,
           PRIMARY KEY(from_id, to_id)
        );
  
  3.4 create table t_relationship_match_01 in instance db_relationship_match_01 
        
        CREATE TABLE t_relationship_match_01 (
           from_id BIGINT,
           to_id BIGINT,
           PRIMARY KEY(from_id, to_id)
        );        
        
# Service
        
1. modify simple-conf.json 
     
        {   "databases": [     
                            {       
                                "label": "user",                  #label don't change        
                                "host": "10.1.235.98",            #your postgreSQL ip, need change to yours       
                                "port": 5432,                     #your postgreSQL port, need change to yours       
                                "credentials": {         
                                                    "user": "tantan007",            #your postgreSQL user, need change to yours         
                                                    "password": "123456"            #your postgreSQL password, need change to yours       
                                    },       
                                "instances": [         
                                                {              
                                                    "name": "db_user_01",         #user resource instance name           
                                                    "tables":["t_user_01"]        #user resource table name         
                                                }       
                                            ]     
                            },     
                            {       
                                "label": "relationship_like",     #label don't change        
                                "host": "10.1.235.98",            #your postgreSQL ip, need change to yours       
                                "port": 5432,                     #your postgreSQL port, need change to yours       
                                "credentials": {         
                                                    "user": "tantan007",            #your postgreSQL user, need change to yours         
                                                    "password": "123456"            #your postgreSQL password, need change to yours       
                                               },       
                                "instances": [         
                                                {           
                                                    "name": "db_relationship_like_01",     #relationship_like resource instance name           
                                                    "tables":["t_relationship_like_01"]    #relationship_like resource table name         
                                                }      
                                             ]     
                            },    
                            {       
                                "label": "relationship_dislike",          #label don't change        
                                "host": "10.1.235.98",                    #your postgreSQL ip, need change to yours      
                                "port": 5432,                             #your postgreSQL port, need change to yours       
                                "credentials": {         
                                                    "user": "tantan007",                    #your postgreSQL user, need change to yours        
                                                    "password": "123456"                    #your postgreSQL password, need change to yours       
                                               },       
                                "instances": [         
                                                {           
                                                    "name": "db_relationship_dislike_01",     #relationship_dislike resource instance name           
                                                    "tables":["t_relationship_dislike_01"]    #relationship_dislike resource table name         
                                                  }       
                                             ]     
                            },     
                            {       
                                "label": "relationship_match",            #label don't change       
                                "host": "10.1.235.98",                    #your postgreSQL ip, need change to yours       
                                "port": 5432,                             #your postgreSQL port, need change to yours       
                                "credentials": {         
                                                    "user": "tantan007",                    #your postgreSQL user, need change to yours         
                                                    "password": "123456"                    #your postgreSQL password, need change to yours       
                                                 },      
                                "instances": [                                    
                                                {           
                                                    "name": "db_relationship_match_01",   #relationship_match resource instance name           
                                                    "tables":["t_relationship_match_01"]  #relationship_dislike resource table name         
                                                }       
                                             ]     
                            }   
                        ],   
                        "idcontroller":{     
                                "host": "localhost",                        #idcontroller addr, if you setup locally， no need change    
                                 "port": 8081                                #idcontroller port, defalut 8081   
                                 },   
                        "service": {     
                                "port": 8080                                #service listen port    
                                }
         }
                       
2. start id controller
        
        before install, please make sure which go version you are using
                    if go version < 1.6 
                    
                        GO15VENDOREXPERIMENT=1 go build
                    
                    else 
                        go build
                        
                        
        cd tantan-test/idcontroller && go build && ./idcontroller
       
        curl 127.0.0.1:8081/internal/idcontroller/user/id
        {"id":"100000000002"}    #ok

3. build   

        cd tantan-test && go build
        
4. start 
        
        ./tantan-test