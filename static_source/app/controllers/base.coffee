'use strict'

angular
  .module('appControllers')
  .controller 'baseCtrl', ['$scope', '$window'
  ($scope, $window) ->
    vm = this

    $scope.back = ()->
      $window.history.back()
  ]