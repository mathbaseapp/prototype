db.evaluator.index.aggregate([
    {
        $match: {
            "key": {$in: ["<mi>t</mi>", "<mi>y</mi>"]}
        }
    },
    {
        $group: {
            _id: "$document.url",
            title: { $first: '$document.title' },
            location: {$push: "$location" },
            count: { $sum: 1 }
        }
    },
    {
        $sort: {count: -1}
    },
    { 
        $limit : 30
    }
])
.forEach(result => print(tojson(result)))


db.evaluator.index.find({"key": {$in: ["<mi>t</mi>"]}})

db.evaluator.index.aggregate([
    {
        $match: {
            "key": {$in: ["<mi>t</mi>", "<mi>y</mi>"]}
        }
    },
    {
        $group: {
            _id: "$document.url",
            title: { $first: '$document.title' },
            location: {$push: "$location" },
            count: { $sum: 1 }
        }
    },
    {
        $sort: {count: -1}
    },
    { 
        $limit : 30
    }
])