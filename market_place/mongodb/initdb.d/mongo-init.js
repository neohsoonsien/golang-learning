db = db.getSiblingDB("marketplace");
db.createUser({
  user: "username",
  pwd: "password",
  roles: [
    {
      role: "readWrite",
      db: "marketplace"
    }
  ]
});

db.createCollection("invoices");
db.invoices.insertMany([
  {
    principalId: "user_1",
    number: "20250620_0001",
    vendor: "Vendor A",
    dateTime: 1753632400,
    details: [
      {
        name: "Item 1",
        category: "Sauce",
        unitPrice: 15.00,
        weightUnit: "g",
        count: 1,
        countUnit: "g",        
      },
      {
        name: "Item 2",
        category: "Sauce",
        unitPrice: 8.00,
        weightUnit: "",
        count: 10,
        countUnit: "",
      }
    ]
  },
  {
    principalId: "user_2",
    number: "20250620_0002",
    vendor: "Vendor B",
    dateTime: 1753637000,
    details: [
      {
        name: "Item 3",
        category: "Veges",
        unitPrice: 9.00,
        weightUnit: "kg",
        count: 3,
        countUnit: "kg",
      },
      {
        name: "Item 4",
        category: "Veges",
        unitPrice: 7.00,
        weightUnit: "",
        count: 6,
        countUnit: "",        
      }
    ]
  },
  {
    principalId: "user_3",
    number: "20250620_0003",
    vendor: "Vendor C",
    dateTime: 1753637130,
    details: [
      {
        name: "Item 5",
        category: "Veges",
        unitPrice: 9.00,
        weightUnit: "kg",
        count: 3,
        countUnit: "kg",
      },
      {
        name: "Item 6",
        category: "Veges",
        unitPrice: 7.00,
        weightUnit: "",
        count: 6,
        countUnit: "",        
      }
    ]
  }
]);
