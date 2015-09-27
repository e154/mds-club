
'use strict'

angular.module('appFilters', [])
angular.module('appControllers', [])
angular.module('appServices', ['ngResource'])
app = angular
  .module('app', [
    'ngRoute'
    'appControllers'
    'appFilters'
    'appServices'
    'route-segment'
    'view-segment'
    'ngSocket'
    'ngMaterial'
    'ui.pagination'
    'ngDialog'
    'ngSanitize'
    'ui.player'
  ])

angular.module('app')
  .config ['$routeProvider', '$locationProvider', '$routeSegmentProvider', '$mdThemingProvider'
  ($routeProvider, $locationProvider, $routeSegmentProvider, $mdThemingProvider) ->
    $routeSegmentProvider
      .when '/',                                                  'base.history'
      .when '/books/page~:page/limit~:limit/author~:author',      'base.books'
      .when '/about',                                             'base.about'
      .when '/lock',                                              'lock'

      .segment 'base',
        templateUrl: '/templates/base.html'
        controller: 'baseCtrl as base'

      .within()
        .segment 'books',
          templateUrl: '/templates/books.html'
          controller: 'booksCtrl'

        .segment 'history',
          default: true
          templateUrl: '/templates/history.html'
          controller: 'historyCtrl as history'

        .segment 'about',
          templateUrl: '/templates/about.html'
          controller: 'aboutCtrl'

      .up()
      .segment 'lock',
        templateUrl: '/templates/lock.html'
        controller: 'lockCtrl'

    $locationProvider.html5Mode
      enabled: true
      requireBase: false

    $routeProvider.otherwise
      redirectTo: '/'

  ]

angular.module('app')
  .run ['$rootScope'
  ($rootScope) =>

    $rootScope.cache =
      author:
        id: 0
        name: ""
      book:
        id: 0
        name: ""

    $rootScope.title = "mds club desktop client"

  ]