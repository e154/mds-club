'use strict'

angular
  .module('appControllers')
  .controller 'authorsCtrl', ['$scope', 'AuthorsResource'
  ($scope, AuthorsResource) =>
    vm = this

    $scope.search_name = ""
    $scope.authors_list = []
    $scope.current_page = 1
    $scope.items_per_page = 24
    $scope.total_items = 1

    updateAuthors = (val)=>
      AuthorsResource.get
        'a1': 'page~'+$scope.current_page
        'a2': 'limit~'+$scope.items_per_page
        'a3': 'search='+val
      ,
        (data)=>
          $scope.authors_list = data.authors
          $scope.total_items = data.total_items
        ,
        (response)=>
          console.log 'error:#{response}'

    $scope.$watch 'search_name',
      (val, old_val)=>
        if val == old_val
          return

        updateAuthors(val)

    updateAuthors("")
  ]