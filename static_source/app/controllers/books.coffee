'use strict'

angular
  .module('appControllers')
  .controller 'booksCtrl', ['$scope', '$rootScope', '$routeSegment', '$location', 'BooksResource'
  ($scope, $rootScope, $routeSegment, $location, BooksResource) ->

    $scope.search_name = $rootScope.cache.book.name
    $scope.book_list = []
    $scope.total_items = 1
    $scope.max_size = 5
    $scope.author = "all"

    $scope.current_page = $routeSegment.$routeParams.page || 1
    $scope.items_per_page = $routeSegment.$routeParams.limit || 24

    updateBooks = ()=>
      if !$scope.search_name?
        $scope.search_name = ""

      $rootScope.cache.book.name = $scope.search_name

      BooksResource.get
        'a1': 'page~'+$scope.current_page
        'a2': 'limit~'+$scope.items_per_page
        'a3': 'author~'+$scope.author
        'a4': 'search='+$scope.search_name
      ,
        (data)=>
          $scope.book_list = angular.copy(data.books)
          $scope.total_items = data.total_items

          $location.path($routeSegment.getSegmentUrl('base.books', page: $scope.current_page, limit: $scope.items_per_page))
      ,
        (response)=>
          console.log 'error:#{response}'

    $scope.$watch 'search_name',
      (val, old_val)=>
        if val == old_val
          return

        $scope.current_page = 1
        updateBooks()

    firstTime = true
    $scope.$watch 'current_page',
      (val, old_val)=>

        if !val || val == old_val
          return

        if firstTime
          console.log(firstTime)
          firstTime = false
          if val == 1
            $scope.current_page = old_val

        updateBooks()

    updateBooks()


  ]