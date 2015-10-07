'use strict'

angular
  .module('appControllers')
  .controller 'baseCtrl', ['$scope', '$window', '$routeSegment', '$location'
  ($scope, $window, $routeSegment, $location) ->
    vm = this

    $scope.back = ()->
      $window.history.back()

    $scope.selectedIndex = null
    $scope.$watch 'selectedIndex', (val, old_val) =>

      if val == old_val
        return

      url = ""
      switch $scope.selectedIndex
        when 0
           url = $routeSegment.getSegmentUrl('base.history')
        when 1
          url = $routeSegment.getSegmentUrl('base.books', page: 1, limit: 24, author: 'all')
        when 2
          url = $routeSegment.getSegmentUrl('base.about')

      $location.path url
  ]