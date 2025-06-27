const data = [
  {
    "id": "Alice",
    "followers": ["Bob", "Charlie", "Dave"]
  },
  {
    "id": "Bob",
    "followers": ["Charlie"]
  },
  {
    "id": "Charlie",
    "followers": []
  },
  {
    "id": "Dave",
    "followers": ["Alice", "Bob"]
  },
  {
    "id": "Eva",
    "followers": ["Bob"]
  },
  {
    "id": "Fiona",
    "followers": ["Eva"]
  }
]

const getNthDegreeFollowers = (id, level) => {

}

// console.log(getNthDegreeFollowers("Bob", 1)) => ["Charlie"]
// console.log(getNthDegreeFollowers("Alice", 2)) => ["Bob", "Charlie", "Dave"]
// console.log(getNthDegreeFollowers("Fiona", 3)) => ["Eva", "Bob", "Charlie"]