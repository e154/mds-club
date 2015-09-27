'use strict'

angular
  .module('appControllers')
  .controller 'booksCtrl', ['$scope', '$rootScope', '$routeSegment', '$location', 'BooksResource', 'ngDialog'
  ($scope, $rootScope, $routeSegment, $location, BooksResource, ngDialog) ->

    $scope.book_list = []
    $scope.total_items = 1
    $scope.max_size = 5

    $scope.search_book_name = $rootScope.cache.book.name
    if $routeSegment.$routeParams.author == "all"
      $scope.search_author_name = ""
    else
      $scope.search_author_name = $routeSegment.$routeParams.author

    $scope.current_page = $routeSegment.$routeParams.page || 1
    $scope.items_per_page = ""

    varsUpdate = ()=>

      $scope.items_per_page = $routeSegment.$routeParams.limit || 24

    $scope.updateBooks = updateBooks = ()=>

      varsUpdate()

      if !$scope.search_book_name?
        $scope.search_book_name = ""

      $rootScope.cache.book.name = $scope.search_book_name

      BooksResource.get
        'a1': 'page~'+$scope.current_page
        'a2': 'limit~'+$scope.items_per_page
        'a3': 'author~'+ ($scope.search_author_name || "all")
        'a4': 'search='+ ($scope.search_book_name || "all")
      ,
        (data)=>
          $scope.book_list = angular.copy(data.books)
          $scope.total_items = data.total_items

          $location.path($routeSegment.getSegmentUrl('base.books', page: $scope.current_page, limit: $scope.items_per_page))

          $scope.openPlayer($scope.book_list[0])
      ,
        (response)=>
          console.log 'error:#{response}'

    firstTime = true
    $scope.$watch 'current_page',
      (val, old_val)=>

        if !val || val == old_val
          return

        if firstTime
          firstTime = false
          if val == 1
            $scope.current_page = old_val

        updateBooks()

    updateBooks()


#    modal player
    $scope.player = null
    $scope.openPlayer = (book)=>

      $scope.book = book
      if !$scope.player
        $scope.player = ngDialog.open(
          template: '/templates/playerModal.html'
          controller: 'playerCtrl'
          className: 'ngdialog-theme-player'
          scope: $scope,
          plain: false,
          overlay: true,
          showClose: false
        )

        $scope.player.closePromise.then (date)=>
          $scope.player = null

#    $scope.$on '$locationChangeSuccess', ()=>
#      updateBooks()
  ]