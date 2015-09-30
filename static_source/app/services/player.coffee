'use strict'

PlayerService = (ngDialog, $rootScope) ->
  (book)=>

    console.log(book)

    player = null

    scope = $rootScope.$new()
    scope.book = book

    if !player
      player = ngDialog.open(
        template: '/templates/playerModal.html'
        controller: 'playerCtrl'
        className: 'ngdialog-theme-player'
        scope: scope,
        plain: false,
        overlay: true,
        showClose: false
      )

      player.closePromise.then (date)=>
        player = null

angular
  .module('appServices')
  .service 'PlayerService', ['ngDialog', '$rootScope', PlayerService]