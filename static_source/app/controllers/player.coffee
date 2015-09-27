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

    $scope.song = {
      src: "http://mds.kallisto.ru/pionerfm/Dzherom_K._Dzherom_-_Kot_Dika_Dankermana.mp3"
      title: "Dzherom_K._Dzherom_-_Kot_Dika_Dankermana"
    }
]