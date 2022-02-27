db.createUser({
    'user': "riverRat",
    'pwd': "Password123",
    'roles': [
        {
            'role': 'readWrite',
            'db': 'river-right'
        },
    ],
});
