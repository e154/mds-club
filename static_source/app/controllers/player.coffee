'use strict'

angular
  .module('appControllers')
  .controller 'playerCtrl', ['$scope'
  ($scope) ->

    if !$scope.book
      $scope.closeThisDialog()



  ]