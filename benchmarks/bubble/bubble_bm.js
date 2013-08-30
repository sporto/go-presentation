var Benchmark = require('benchmark');
var bubble = require('./bubble');

var suite = new Benchmark.Suite;

suite
	.add('bubble', function() {
		var a = [2, 10, 1, 9, 5, 6, 8, 3, 7, 4];
		bubble(a);
	})
	// add listeners
	.on('cycle', function(event) {
		console.log(String(event.target));
	})
	// .on('complete', function() {
	// 	console.log('Fastest is ' + this.filter('fastest').pluck('name'));
	// })
	.run({'async': false});