'use strict'

angular
  .module('appControllers')
  .controller 'baseCtrl', ['$scope', '$window', '$routeSegment', '$location'
  ($scope, $window, $routeSegment, $location) ->
    vm = this

    $scope.back = ()->
      $window.history.back()

  ]