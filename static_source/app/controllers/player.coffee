'use strict'

angular
  .module('appControllers')
  .controller 'playerCtrl', ['$scope', 'FileResource', 'HistoryResource', '$timeout'
  ($scope, FileResource, HistoryResource, $timeout) ->

    if !$scope.book
      $scope.closeThisDialog()

    $scope.options = {
      audioWidth: "100%"
      audioHeight: 40
      startVolume: 0.5
      loop: false
      features: ['playpause','current','progress','duration','tracks','volume','fullscreen', 'playlistfeature', 'playlist', 'prevtrack', 'nexttrack']
      playlist: true
      playlistposition: 'bottom'
    }

    addToHistory = (id)=>
      HistoryResource.post
        'a1': 'book~' + id
      ,
        {}
      ,
        (data)=>
          console.log(data)
      ,
        (response)=>
          console.log 'error:#{response}'

#   player init
    $timeout ()->
      $scope.audio = document.getElementsByTagName('audio')[0]
      $scope.audio.addEventListener 'play', (event)=>
        console.log("play")
        addToHistory($scope.book.Id)

    , 1000

    $scope.files = []
    FileResource.get
      'a1': 'list'
      'a2': 'book~' + $scope.book.Id
    ,
      (data)=>
        if !data.files
          return

        $scope.files = []
        angular.forEach data.files, (file, key)->
          if file.Url && file.Url.indexOf("http://") > -1
            $scope.files.push(file)
      ,
      (response)=>
        console.log 'error:#{response}'
  ]