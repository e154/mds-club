'use strict'

angular
  .module('appControllers')
  .controller 'playerCtrl', ['$scope', '$sce'
  ($scope, $sce) ->

    if !$scope.book
      $scope.closeThisDialog()

    $scope.options = {
      audioWidth: "100%"
      audioHeight: 40
      startVolume: 0.5
      features: ['playpause','current','progress','duration','tracks','volume','fullscreen']
    }

]