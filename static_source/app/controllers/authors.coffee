'use strict'

angular
  .module('appControllers')
  .controller 'authorsCtrl', ['$scope', 'AuthorsResource', '$rootScope', '$routeSegment', '$location'
  ($scope, AuthorsResource, $rootScope, $routeSegment, $location) =>

    vm = this

    console.log("authorsCtrl")

    $scope.search_name = $rootScope.cache.author.name
    $scope.authors_list = []
    $scope.total_items = 1
    $scope.max_size = 5

    $scope.current_page = $routeSegment.$routeParams.page || 1
    $scope.items_per_page = $routeSegment.$routeParams.limit || 24

    updateAuthors = ()=>
      if !$scope.search_name?
        $scope.search_name = ""

      $rootScope.cache.author.name = $scope.search_name

      AuthorsResource.get
        'a1': 'page~'+$scope.current_page
        'a2': 'limit~'+$scope.items_per_page
        'a3': 'search='+$scope.search_name
      ,
        (data)=>
          $scope.authors_list = data.authors
          $scope.total_items = data.total_items

          $location.path($routeSegment.getSegmentUrl('base.authors', page: $scope.current_page, limit: $scope.items_per_page))
        ,
        (response)=>
          console.log 'error:#{response}'

    $scope.$watch 'search_name',
      (val, old_val)=>
        if val == old_val
          return

        $scope.current_page = 1
        updateAuthors()

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

        updateAuthors()

    updateAuthors()

  ]