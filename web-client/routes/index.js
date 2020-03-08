var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { 
    title: 'Hash-n-go',
    time: {
      selected: 60,
      options: [60, 50, 30, 20, 10, 5, 1, 0.5, 0.1]
    }
  });
});

module.exports = router;
