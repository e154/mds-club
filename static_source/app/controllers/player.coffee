'use strict'

angular
  .module('appControllers')
  .controller 'playerCtrl', ['$scope', 'FileResource'
  ($scope, FileResource) ->

    if !$scope.book
      $scope.closeThisDialog()

    $scope.options = {
      audioWidth: "100%"
      audioHeight: 40
      startVolume: 0.5
      features: ['playpause','current','progress','duration','tracks','volume','fullscreen']
    }

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
        console.log($scope.files)
    ,
      (response)=>
        console.log 'error:#{response}'
  ]