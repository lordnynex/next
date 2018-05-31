// Create an application user.

db.users.update(
  { email: "admin@example.com" },
  {
    "$set": {
      updated_at: new Date()
    },
    "$setOnInsert": {
      created_at: new Date(),
      name: "",
      role: "admin"
    }
  },
  { upsert: true }
)
