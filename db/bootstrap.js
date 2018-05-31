// Bootstrap the Mongo database.

//
// Create indexes.
//

db.users.createIndex({ email: 1 }, { unique: true });

//
// Seed the db.
//
