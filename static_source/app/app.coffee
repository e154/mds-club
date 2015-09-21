
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
  ])

angular.module('app')
  .config ['$routeProvider', '$locationProvider', '$routeSegmentProvider', '$mdThemingProvider'
  ($routeProvider, $locationProvider, $routeSegmentProvider, $mdThemingProvider) ->
    $routeSegmentProvider
      .when '/',              'base.books'
      .when '/authors',       'base.authors'
      .when '/history',       'base.history'
      .when '/about',         'base.about'
      .when '/lock',          'lock'

      .segment 'base',
        templateUrl: '/templates/base.html'
        controller: 'baseCtrl as base'

      .within()
        .segment 'books',
          default: true
          templateUrl: '/templates/books.html'
          controller: 'booksCtrl as books'

        .segment 'authors',
          templateUrl: '/templates/authors.html'
          controller: 'authorsCtrl'
          controllerAs: 'authors'

        .segment 'history',
          templateUrl: '/templates/history.html'
          controller: 'historyCtrl as history'

        .segment 'about',
          templateUrl: '/templates/about.html'
          controller: 'aboutCtrl as about'

      .up()
      .segment 'lock',
        templateUrl: '/templates/lock.html'
        controller: 'lockCtrl as lock'

    $locationProvider.html5Mode
      enabled: true
      requireBase: false

    $routeProvider.otherwise
      redirectTo: '/'

    $mdThemingProvider.theme('purple')
  ]

angular.module('app')
  .run ['$rootScope'
  ($rootScope) =>

#    gui = require('nw.gui')
#    win = gui.Window.get()
#    tray
#
#    win.on 'minimize', () ->
#      this.hide()
#      tray = new gui.Tray
#        icon: '/images/icon.png'
#      tray.on 'click', () ->
#        win.show()
#        this.remove()
#        tray = null
  ]