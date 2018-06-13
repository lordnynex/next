// Init a Mongo database.

//
// Create indexes.
//

db.users.createIndex({ email: 1 }, { unique: true });

//
// Seed a database.
//
