'use strict'

angular
  .module('appControllers')
  .controller 'playerCtrl', ['$scope', '$sce'
  ($scope, $sce) ->

    if !$scope.book
      $scope.closeThisDialog()

    $scope.config = {
      sources: [
        {src: $sce.trustAsResourceUrl("http://static.videogular.com/assets/videos/videogular.mp4"), type: "video/mp4"}
        {src: $sce.trustAsResourceUrl("http://static.videogular.com/assets/videos/videogular.webm"), type: "video/webm"}
        {src: $sce.trustAsResourceUrl("http://static.videogular.com/assets/videos/videogular.ogg"), type: "video/ogg"}
      ]
      theme: "bower_components/videogular-themes-default/videogular.css"
      plugins: {
        poster: "http://www.videogular.com/assets/images/videogular.png"
      }
    }

  ]