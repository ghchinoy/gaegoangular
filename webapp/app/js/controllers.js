'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
  .controller('WordCtrl', ['$scope', '$http', function($scope, $http) {

  	var words = {};

  	var getWords = function() {
  		$http.get("/getWords")
  		.success(function(data) {
  			words = data;
  			console.log("Data ", data);
  		});
  	};

  	getWords();

  	var pos = 0;
  	$scope.next = function() {
  		if (pos === words.length-1 ) {
  			pos = 0;
  		} else {
  			pos++;
  		}
  		$scope.word = words[pos];
  	};

  	var next = function() {
  		
  	};

  }]);
