'use strict'

angular
  .module('appControllers')
  .controller 'historyCtrl', ['$scope', 'HistoryResource', '$routeSegment', 'PlayerService', 'BooksResource'
  ($scope, HistoryResource, $routeSegment, PlayerService, BooksResource) ->

    $scope.current_page = $routeSegment.$routeParams.page || 1
    $scope.items_per_page = $routeSegment.$routeParams.limit || 24
    $scope.total_items = 0
    $scope.history = []

    updateHistory = ()->
      HistoryResource.get
        'a1': 'page~'+$scope.current_page
        'a2': 'limit~'+$scope.items_per_page
      ,
        (data)=>
          $scope.total_items = data.total_items
          $scope.history = []
          angular.forEach data.history, (story, key)=>
            $scope.history.push(story)
      ,
        (response)=>
          console.log 'error:#{response}'

    updateHistory()

    firstTime = true
    $scope.$watch 'current_page',
      (val, old_val)=>

        if !val || val == old_val
          return

        if firstTime
          firstTime = false
          if val == 1
            $scope.current_page = old_val

        updateHistory()

    $scope.openPlayer = (book_id)=>
      BooksResource.get
        'a1': 'id~' + book_id
      ,
        (data)=>
          PlayerService(data.book)
      ,
        (response)=>
          console.log 'error:#{response}'
  ]