CREATE KEYSPACE IF NOT EXISTS main
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  };
